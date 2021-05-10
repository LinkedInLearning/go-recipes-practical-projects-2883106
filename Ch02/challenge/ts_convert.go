package main

import (
	"fmt"
)

// tsConvert convert time stamp in "YYYY-MM-DDTHH:MM" format from one time zone to another
func tsConvert(ts, from, to string) (string, error) {
	return "", nil // FIXME
}

func main() {
	ts := "2021-03-08T19:12"
	out, err := tsConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Println(out) // 2021-03-09T05:12
}
