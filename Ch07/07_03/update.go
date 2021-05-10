package main

import (
	"log"
	"math/rand"
	"time"
)

func update(host, version string) {
	// TODO
	n := rand.Intn(100) + 50
	time.Sleep(time.Duration(n) * time.Millisecond)
	log.Printf("%s updated to %s", host, version)
}
