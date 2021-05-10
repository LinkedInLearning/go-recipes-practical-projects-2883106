package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Record is a weather record
type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"` // celsius
	Rain    float64 // millimeter
}

func readRecord(r io.Reader) (Record, error) {
	var rec Record
	dec := json.NewDecoder(r)
	if err := dec.Decode(&rec); err != nil {
		return Record{}, err
	}

	return rec, nil
}

func main() {
	file, err := os.Open("record.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rec, err := readRecord(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rec)
	// {Time:2020-03-06 00:00:00 +0000 UTC Station:DS9 Temp:21.6 Rain:0}
}
