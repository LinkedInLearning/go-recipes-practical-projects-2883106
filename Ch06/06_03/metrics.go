package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	db *DB
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	Host   string    `json:"host"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var m Metric
	const maxSize = 1 << 20 // MB
	dec := json.NewDecoder(io.LimitReader(r.Body, maxSize))
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := db.Add(m)
	log.Printf("metric: %+v (id=%s)", m, id)

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"id": id,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error reply: %s", err)
	}
}

func main() {
	http.HandleFunc("/metric", handleMetric)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
