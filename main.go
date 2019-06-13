package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"people-api/db"
	"people-api/routes"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		)

	router.Route("/app", func(r chi.Router){
		r.Mount("/people", routes.Routes() )
	})

	return router
}


func main() {
	fmt.Printf("hello, world\n")

	router := Routes()

	db.Create()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler ) error {
		log.Printf("%s%s\n", method, route )
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("logging err %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8081", router))
}