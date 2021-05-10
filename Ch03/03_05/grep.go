package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// grep returns lines in r that contain term
func grep(r io.Reader, term string) ([]string, error) {
	var matches []string
	s := bufio.NewScanner(r)
	for s.Scan() {
		if strings.Contains(s.Text(), term) {
			matches = append(matches, s.Text())
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func main() {
	file, err := os.Open("journal.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matches, err := grep(file, "System is rebooting")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d reboots\n", len(matches))
}
