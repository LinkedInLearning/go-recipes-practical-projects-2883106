package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book is information about book
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func handleGetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	book, err := getBook(isbn)
	if err != nil {
		log.Printf("error - get: unknown ISBN - %q", isbn)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Printf("error - json: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{isbn}", handleGetBook).Methods("GET")

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
