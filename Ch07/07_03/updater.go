package main

import (
	"fmt"
	"log"
	"sync"
)

func updateAll(version string, hosts <-chan string) {
	var wg sync.WaitGroup
	for host := range hosts {
		wg.Add(1)
		go func(host, version string) {
			defer wg.Done()
			update(host, version)
		}(host, version)
	}

	wg.Wait()
}

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			host := fmt.Sprintf("srv%d", i+1)
			ch <- host
		}
		close(ch)
	}()

	version := "1.0.2"
	updateAll(version, ch)
	log.Printf("all servers updated")
}
