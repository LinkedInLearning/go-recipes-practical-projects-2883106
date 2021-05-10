package main

import (
	"fmt"
	"time"
)

func main() {
	lennon := time.Date(1940, time.October, 9, 18, 30, 0, 0, time.UTC)
	fmt.Println(lennon) // 1940-10-09 18:30:00 +0000 UTC

	fmt.Println(lennon.Format("2006-01-02"))  // 1940-10-09
	fmt.Println(lennon.Format("Mon, Jan 02")) // Wed, Oct 09

	fmt.Println(lennon.Format(time.RFC3339Nano)) // 1940-10-09T18:30:00Z

	d := 3500 * time.Millisecond
	fmt.Println(d) // 3.5s
}
