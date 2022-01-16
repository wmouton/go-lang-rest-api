package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string `json:"lastname"`
}

// Init books variable as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatalln("Error: ", err)
	}
}
// Get A Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				log.Fatalln("Error: ", err)
			}
			return
		}
	}
	if err := json.NewEncoder(w).Encode(&Book{}); err != nil {
		log.Fatalln(err)
	}
}

// Create A Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe (not for production)
	books = append(books, book)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Fatalln("Error: ", err)
	}
}

// Update A Book
func updateBook(_ http.ResponseWriter, _ *http.Request) {}

// Delete A Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatalln("Error: ", err)
	}
}

func main() {
	// Server Message On Startup
	log.Println("sever started on port 8080...")
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "681565", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
