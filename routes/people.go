package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"people-api/db"
	"people-api/utils"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{personId}", GetPerson)
	router.Delete("/{personId}", DeleteAPerson)
	router.Post("/", CreatePerson)
	router.Get("/", GetAllPeople)
	router.Options("/", ReturnOK)


	return router
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	personId := chi.URLParam(r, "personId")
	returnedPerson, err := db.GetById(personId)

	if err != nil {
		response := make(map[string]string)
		response["message"] = "error getting person"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		render.JSON(w, r, response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, returnedPerson)
	}


}

func ReturnOK(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "success"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

func DeleteAPerson(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	personId := chi.URLParam(r, "personId")
	db.DeleteById(personId)

	response := make(map[string]string)
	response["message"] = "successfully deleted"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}


func CreatePerson(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	print("here we go")
	person, err := utils.ParsePersonJson(r)

	if err != nil {
		panic(err)
	}

	createdPerson := db.Save(person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	render.JSON(w, r, createdPerson)
}



func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	people, err := db.GetAll()

	if err != nil {
		panic(err)
	}

	render.JSON(w, r, people)
}