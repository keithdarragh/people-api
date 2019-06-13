package models

type Person struct {
	ID 		string `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Balance string `json:"balance"`
	Email   string `json:"email"`
	Address Address `json:"address"`
}
