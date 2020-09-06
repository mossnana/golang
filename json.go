package main

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// Person is person
type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
	Movies    []Movie `json:"movies"`
}

// Address is address
type Address struct {
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}

// Movie is movie
type Movie struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

func main() {
	data := []byte(`
		{
			"first_name": "bob",
			"last_name": "malshal",
			"address": {
				"postcode": "50300",
				"country": "thailand"
			},
			"movies": [
				{
					"name": "หอแต๋วแตก",
					"year": 2000
				},
				{
					"name": "หอแต๋วไม่แตก",
					"year": 2001
				}
			]
		}
	`)
	var person Person
	var personWithJsoniter Person
	// Use standard
	err1 := json.Unmarshal(data, &person)
	// Use jsoniter for better performance
	var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary
	err2 := jsonTool.Unmarshal(data, &personWithJsoniter)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
	}
	fmt.Println(person)
}
