package main

import "fmt"

func MostCommon[T comparable](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("MostCommon on empty slice")
	}

	freq := make(map[T]int) // value -> count
	for _, v := range values {
		freq[v]++
	}

	var mode T
	max := 0
	for v, count := range freq {
		if count > max {
			mode, max = v, count
		}
	}

	return mode, nil
}

func main() {
	fmt.Println(MostCommon([]int{3, 1, 4, 1, 5, 9, 2, 6}))
	fmt.Println(MostCommon([]rune{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println(MostCommon[int](nil))
}
