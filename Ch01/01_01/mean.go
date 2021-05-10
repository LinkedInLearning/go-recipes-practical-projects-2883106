package main

import (
	"fmt"
)

func main() {
	fmt.Println(mean([]int{1, 2, 3}))    // 2
	fmt.Println(mean([]int{1, 2, 3, 4})) // 2.5
}

func mean(nums []int) float64 {
	s := sum(nums)
	return float64(s) / float64(len(nums))
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	return total
}
