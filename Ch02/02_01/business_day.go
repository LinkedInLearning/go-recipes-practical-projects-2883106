package main

import (
	"fmt"
	"time"
)

func isBusinessDay(date time.Time) bool {
	wday := date.Weekday()
	if wday == time.Saturday || wday == time.Sunday {
		return false
	}

	return true
}

func nextBusinessDay(date time.Time) time.Time {
	const day = 24 * time.Hour
	for {
		date = date.Add(day)
		if isBusinessDay(date) {
			break
		}
	}

	return date
}

func main() {
	date := time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println(date, date.Weekday()) // 2021-12-31 00:00:00 +0000 UTC Friday
	nbd := nextBusinessDay(date)
	fmt.Println(nbd, nbd.Weekday()) // 2022-01-03 00:00:00 +0000 UTC Monday

	date = time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC)
	fmt.Println(date, date.Weekday()) // 2022-01-04 00:00:00 +0000 UTC Tuesday
	nbd = nextBusinessDay(date)
	fmt.Println(nbd, nbd.Weekday()) // 2022-01-05 00:00:00 +0000 UTC Wednesday
}
