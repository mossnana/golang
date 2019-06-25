package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init Books Data
// []Book is mean you want books to collection
var books []Book

// Book Struct ( === Class) (Model in MVC)
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

// Author Struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

// Router Functions
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("params", params)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// Run function main ตั้งแต่เริ่มโดย ไม่ต้องไปเรียก main() ภายหลัง
func main() {
	// Init Mux Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "4883021", Title: "Book Book", Author: &Author{Firstname: "Permpoon", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "4883022", Title: "Book x", Author: &Author{Firstname: "Permpoon", Lastname: "er"}})
	books = append(books, Book{ID: "3", Isbn: "4883023", Title: "x Book", Author: &Author{Firstname: "Permpoon", Lastname: "Dasdoe"}})
	books = append(books, Book{ID: "4", Isbn: "4883024", Title: "x x", Author: &Author{Firstname: "Permpoon", Lastname: "sda"}})

	// Route Handlers or Endpoints
	// e.g.
	// router.HandleFunc("{route}", {handler function}).Methods("GET")
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// http package
	// log.Fatal() = throw error
	log.Fatal(http.ListenAndServe(":8000", router))
}
