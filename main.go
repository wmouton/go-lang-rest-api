package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Method("GET")
	router.HandleFunc("/api/books/{id}", getBook).Method("GET")
	router.HandleFunc("/api/books", createBook).Method("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Method("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Method("DELETE")

}
