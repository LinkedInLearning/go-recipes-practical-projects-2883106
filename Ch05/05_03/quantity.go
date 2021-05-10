package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Quantity is combination of value and unit (e.g. 2.7cm)
type Quantity struct {
	Value float64
	Unit  string
}

// MarshalJSON implements the json.Marshaler interface
// Example encoding: "42.195km"
func (q *Quantity) MarshalJSON() ([]byte, error) {
	if q.Unit == "" {
		return nil, fmt.Errorf("empty  unit")
	}
	text := fmt.Sprintf("%f%s", q.Value, q.Unit)
	return json.Marshal(text)
}

func main() {
	q := Quantity{1.78, "meter"}
	json.NewEncoder(os.Stdout).Encode(&q) // "1.780000meter"
}
