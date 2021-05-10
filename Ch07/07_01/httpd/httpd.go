package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const timeFmt = "2006-01-02T15:04:05"

func handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:] // trim leading /

	day, err := time.Parse("2006-01-02", s)
	if err != nil {
		msg := fmt.Sprintf("bad date: %q", s)
		log.Printf("error: %s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "start,end,distance,passengers\n")

	start := day
	nextDay := day.Add(24 * time.Hour)
	for start.Before(nextDay) {
		start = start.Add(3 * time.Second)
		end := start.Add(17 * time.Second)
		distance := 3.14
		passengers := 1
		fmt.Fprintf(w, "%s,%s,%.2f,%d\n", start.Format(timeFmt), end.Format(timeFmt), distance, passengers)
	}
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
