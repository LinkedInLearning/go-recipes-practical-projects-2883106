package main

import (
	"fmt"
)

// filter returns a slice with only the values that pred(val) returned true
func filter(pred func(int) bool, values []int) []int {
	var out []int
	for _, val := range values {
		if pred(val) {
			out = append(out, val)
		}
	}

	return out
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(filter(isOdd, values)) // [1 3 5 7]
}
