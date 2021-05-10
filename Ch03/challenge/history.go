package main

import (
	"fmt"
	"log"
)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func cmdFreq(fileName string) (map[string]int, error) {
	// FIXME
	return nil, nil
}

func main() {
	freqs, err := cmdFreq("zsh_history")
	if err != nil {
		log.Fatal(err)
	}

	for cmd, count := range freqs {
		fmt.Printf("%s -> %d\n", cmd, count)
	}
}
