package main

import (
	"fmt"
)

func ExampleCommaOK() {
	prices := map[string]int{
		"Banana": 0, // Banana's are free!
	}

	price, ok := prices["Banana"]
	if ok {
		fmt.Printf("The price of Banana is $%d\n", price)
	} else {
		fmt.Printf("We don't have Bananas")
	}

	price, ok = prices["Apple"]
	if ok {
		fmt.Printf("The price of Apple is $%d\n", price)
	} else {
		fmt.Printf("We don't have Apples")
	}

	// Output:
	// The price of Banana is $0
	// We don't have Apples
}
