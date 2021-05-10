package main

import (
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID   string
	Time time.Time
}

type DoorEvent struct {
	Event
	Action string // open, close
}

type TemperatureEvent struct {
	Event
	Value float64
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}

	evt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}
	return &evt, nil
}

func main() {
	evt, err := NewDoorEvent("front door", time.Now(), "open")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", evt)
	// &{Event:{ID:front door Time:2021-04-30 14:47:40.31038 +0300 IDT m=+0.000170354} Action:open}
}
