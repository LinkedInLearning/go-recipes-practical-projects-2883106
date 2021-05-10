package main

import (
	"fmt"
	"log"
	"time"
)

func timeit(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		log.Printf("%s took %s", name, duration)
	}
}

func dot(v1, v2 []float64) (float64, error) {
	defer timeit("dot")()

	if len(v1) != len(v2) {
		return 0, fmt.Errorf("dot of different size vectors")
	}

	d := 0.0
	for i, val1 := range v1 {
		val2 := v2[i]
		d += val1 * val2
	}

	return d, nil
}

func main() {
	v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(dot(v, v))
}
