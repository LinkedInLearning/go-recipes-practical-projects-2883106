package main

import (
	"fmt"
)

// A Thermostat measures and controls the temperature
type Thermostat struct {
	ID string

	value float64
}

// Value return the current temperature in Celsius
func (t *Thermostat) Value() float64 {
	return t.value
}

// Set tells the thermostat to set the temperature
func (t *Thermostat) Set(value float64) {
	t.value = value
}

// Kind returns the device kind
func (*Thermostat) Kind() string {
	return "thermostat"
}

func main() {
	t := Thermostat{"Living Room", 16.2}
	fmt.Printf("%s before: %.2f\n", t.ID, t.Value())
	// Living Room before: 16.20
	t.Set(18)
	fmt.Printf("%s after:  %.2f\n", t.ID, t.Value())
	// Living Room after:  18.00
}
