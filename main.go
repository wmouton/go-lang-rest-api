package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Iibn	string	`json:"isbn"`
	Iitle	string	`json:"title"`
	Author	*Author	`json:"id"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"'`
	Lastname	string `json:"lastname"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {}
// Get A Book
func getBook(w http.ResponseWriter, r *http.Request) {}
// Create A Book
func createBook(w http.ResponseWriter, r *http.Request) {}
// Update A Book
func updateBook(w http.ResponseWriter, r *http.Request) {}
// Delete A Book
func deleteBook(w http.ResponseWriter, r *http.Request) {}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Method("GET")
	router.HandleFunc("/api/books/{id}", getBook).Method("GET")
	router.HandleFunc("/api/books", createBook).Method("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Method("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Method("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
