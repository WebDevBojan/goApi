package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// init books var as a slice book struct
var books []Book 

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBoook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "123123", Title: "Book Title", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "433123", Title: "Book Two", Author: &Author{Firstname: "Ben", Lastname: "Banana"}})

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBoook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}