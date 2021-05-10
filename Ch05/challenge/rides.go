// What is the maximal ride speed in rides.json?
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func maxRideSpeed(r io.Reader) (float64, error) {
	// FIXME
	return 0, nil
}

func main() {
	file, err := os.Open("rides.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed) // 40.5
}
