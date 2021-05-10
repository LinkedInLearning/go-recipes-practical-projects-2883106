package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	db     = make(map[string][]byte)
	dbLock sync.RWMutex
)

func handleSet(w http.ResponseWriter, r *http.Request) {
	// FIXME
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// FIXME
}

func handleList(w http.ResponseWriter, r *http.Request) {
	// FIXME
}

func main() {
	r := mux.NewRouter()
	// FIXME: routing
	http.Handle("/", r)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
