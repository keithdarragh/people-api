package models

type Address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City string `json:"city"`
	Postcode string `json:"postcode"`
	Country string `json:"country"`
}