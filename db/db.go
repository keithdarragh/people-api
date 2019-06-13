package db

import (
	"encoding/json"
	"fmt"
	r "gopkg.in/dancannon/gorethink.v5"
	"gopkg.in/rethinkdb/rethinkdb-go.v5"
	"log"
	"people-api/models"
)

var session *r.Session

func Create() {
	fmt.Println("Connecting to RethinkDB")

	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "people",
	})

	if err != nil {
		log.Fatal("Could not connect")
	}

	err = r.DB("people").TableDrop("people").Exec(session)
	err = r.DB("people").TableCreate("people").Exec(session)
	if err != nil {
		log.Fatal("Could not create table")
	}

	err = r.DB("people").Table("people").IndexCreate("ID").Exec(session)
	if err != nil {
		log.Fatal("Could not create index")
	}
}

func Save(person models.Person) models.Person {
	cursor, err := r.Table("people").Insert(person).RunWrite(session)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cursor.GeneratedKeys)
	person.ID = cursor.GeneratedKeys[0]
	return person
}

func GetAll() ([]models.Person, error) {
	var returnedPerson []models.Person

	cursor, err := r.Table("people").Run(session)

	if err != nil {
		fmt.Println(err)
		return returnedPerson, rethinkdb.ErrEmptyResult
	}

	if cursor.IsNil() {
		fmt.Println(err)
		return []models.Person{}, nil
	}

	cursor.All(&returnedPerson)
	cursor.Close()

	return returnedPerson, nil

}

func GetById(id string) (models.Person, error) {
	var returnedPerson models.Person

	cursor, err := r.Table("people").Get(id).Run(session)
	if err != nil {
		fmt.Println(err)
		return returnedPerson, rethinkdb.ErrEmptyResult
	}
	if cursor.IsNil() {
		return returnedPerson, rethinkdb.ErrEmptyResult
	}

	cursor.One(&returnedPerson)
	cursor.Close()

	return returnedPerson, nil

}

func DeleteById(id string) bool {
	_, err := r.Table("people").Get(id).Delete().Run(session)
	if err != nil {
		fmt.Println(err)
	}

	return true
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}