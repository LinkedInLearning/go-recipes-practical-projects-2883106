package main

import (
	"fmt"
	"log"
)

type ClickEvent struct {
	// ...
}

type HoverEvent struct {
	// ...
}

var eventCounts = make(map[string]int) // type -> count

func recordEvent(evt interface{}) {
	switch evt.(type) {
	case *ClickEvent:
		eventCounts["click"]++
	case *HoverEvent:
		eventCounts["hover"]++
	default:
		log.Printf("warning: unknown event: %#v of type %T\n", evt, evt)
	}
}

func main() {
	recordEvent(&ClickEvent{})
	recordEvent(&HoverEvent{})
	recordEvent(&ClickEvent{})
	recordEvent(3)
	// 2021/04/30 15:07:17 warning: unknown event: 3 of type int

	fmt.Println("event counts:", eventCounts)
	// event counts: map[click:2 hover:1]
}
