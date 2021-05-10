package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/* Example history file
: 1542784278:0;git push
: 1542784308:0;ls
: 1542784310:0;go test
: 1542784314:0;go test -v
: 1542784386:0;which gometalinter
: 1542784314:0;go test -v
*/
var cmdRe = regexp.MustCompile(`;go ([a-z]+)`)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func cmdFreq(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	freqs := make(map[string]int)
	s := bufio.NewScanner(file)
	for s.Scan() {
		matches := cmdRe.FindStringSubmatch(s.Text())
		if len(matches) == 0 {
			continue
		}
		cmd := matches[1]
		freqs[cmd]++
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
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
