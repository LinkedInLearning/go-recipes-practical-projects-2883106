package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

const (
	maxSize = 10 * (1 << 20) // 10MB
)

var (
	db     = make(map[string][]byte)
	dbLock sync.RWMutex
)

func handleSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, maxSize)
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Printf("read error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	db[key] = data

	resp := map[string]interface{}{
		"key":  key,
		"size": len(data),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	dbLock.RLock()
	defer dbLock.RUnlock()

	data, ok := db[key]
	if !ok {
		log.Printf("error get - unknown key: %q", key)
		http.Error(w, fmt.Sprintf("%q not found", key), http.StatusNotFound)
		return
	}

	if _, err := w.Write(data); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {
	dbLock.RLock()
	defer dbLock.RUnlock()

	keys := make([]string, 0, len(db))
	for key := range db {
		keys = append(keys, key)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(keys); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kv/{key}", handleSet).Methods("POST")
	r.HandleFunc("/kv/{key}", handleGet).Methods("GET")
	r.HandleFunc("/kv", handleList).Methods("GET")
	http.Handle("/", r)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
