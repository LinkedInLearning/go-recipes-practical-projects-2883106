package main

import (
	"fmt"
	"strings"
)

// Letter in Greek
type Letter struct {
	Symbol  string
	English string
}

var letters = []Letter{
	{"Σ", "Sigma"},
	// TODO
}

// englishFor return the English name for a greek letter
func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}

	return "", fmt.Errorf("unknown greek letter: %#v", greek)
}

func main() {
	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
}
