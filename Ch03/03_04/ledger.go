package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

/*
12 shares of MSFT for $234.57
10 shares of TSLA for $692.4
*/
var transRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)

// Transaction is a b
type Transaction struct {
	Symbol string
	Volume int
	Price  float64
}

func parseLine(line string) (Transaction, error) {
	matches := transRe.FindStringSubmatch(line)
	if matches == nil {
		return Transaction{}, fmt.Errorf("bad line: %q", line)
	}
	var t Transaction
	t.Symbol = matches[2]
	t.Volume, _ = strconv.Atoi(matches[1])
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}

func main() {
	line := "12 shares of MSFT for $234.57"
	t, err := parseLine(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t) // {Symbol:MSFT Volume:12 Price:234.57}
}
