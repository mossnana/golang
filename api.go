package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// New Person Data Type
type Person struct {
	ID string `json:"id, omitemply"`
	Firstname string `json:"firstname, omitemply"`
	Lastname string `json:"lastname, omitemply"`
	Address *Address `json:"address, omitemply"`
}

// New Address Sub-Data Type
type Address struct {
	City string `json:"city, omitemply"`
	State string `json:"state, omitemply"`
}

// Main Array to Storage Data
var people []Person

// Get All User Function
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		} else {
			io.WriteString(w, "No User Found!!!")
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Get One User Function
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

// Put New User
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	if person.ID != params["id"]{
		person.ID = params["id"]
		people = append(people, person)
		json.NewEncoder(w).Encode(people)
		return
	}else{
		io.WriteString(w, "Duplicated User!!!")
		return
	}
}

// Delete a User
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		} else {
			io.WriteString(w, "No User Found!!!")
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

// Main Function
func main() {
	// Router import
	router := mux.NewRouter()
	// Save Data to Array for First
	people = append(people, Person{ID: "1", Firstname: "Permpoon", Lastname: "Chao", Address: &Address{City: "Thailand", State: "ChiangMai"}})
	people = append(people, Person{ID: "2", Firstname: "Wiraphorn", Lastname: "Wanna", Address: &Address{City: "Thailand", State: "ChiangRai"}})
	// Routing Path Config
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	// Start Server
	log.Fatal(http.ListenAndServe(":12345", router))
}
