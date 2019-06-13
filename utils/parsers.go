package utils

import (
	"encoding/json"
	"net/http"
	"people-api/models"
)

func ParsePersonJson(r *http.Request) (models.Person, error) {
	decoder := json.NewDecoder(r.Body)
	var person models.Person
	err := decoder.Decode(&person)
	return person, err
}