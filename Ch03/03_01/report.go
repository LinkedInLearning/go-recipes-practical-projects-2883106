package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Trade represents a trade
type Trade struct {
	Symbol string
	Volume int
	Price  float64
}

// genReport generates a fixed with report in the format
// Symbol: 10 chars, left padded
// Volume: 4 digits, 0 padded
// Price: 2 digits after the decimal
func genReport(w io.Writer, trades []Trade) {
	for i, t := range trades {
		log.Printf("%d: %#v", i, t)
		// ... 2: main.Trade{Symbol:"BRK-A", Volume:1, Price:399100}
		fmt.Fprintf(w, "%-10s %04d %.2f\n", t.Symbol, t.Volume, t.Price)
		// MSFT       0231 234.57
	}
}

func main() {
	log.SetPrefix("LOG: ")

	trades := []Trade{
		{"MSFT", 231, 234.57},
		{"TSLA", 123, 686.75},
		{"BRK-A", 1, 399100},
	}
	genReport(os.Stdout, trades)
}
