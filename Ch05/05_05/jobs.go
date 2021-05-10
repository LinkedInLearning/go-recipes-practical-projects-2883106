package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// StartJob is a request to start a job
type StartJob struct {
	Type  string
	User  string
	Count int
}

// JobStatus is a request for job status
type JobStatus struct {
	Type string
	ID   string
}

func handleStart(req StartJob) error {
	fmt.Printf("start: %#v\n", req)
	return nil // FIXME
}

func handleStatus(req JobStatus) error {
	fmt.Printf("status: %#v\n", req)
	return nil // FIXME
}

func handleRequest(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	val, ok := m["type"]
	if !ok {
		return fmt.Errorf("'type' missing from JSON")
	}

	typ, ok := val.(string)
	if !ok {
		return fmt.Errorf("'type' is not a string")
	}

	switch typ {
	case "start":
		var sj StartJob
		if err := mapstructure.Decode(m, &sj); err != nil {
			return fmt.Errorf("bad 'start' request: %w", err)
		}
		return handleStart(sj)
	case "status":
		var js JobStatus
		if err := mapstructure.Decode(m, &js); err != nil {
			return fmt.Errorf("bad 'status' request: %w", err)
		}
		return handleStatus(js)
	}

	return fmt.Errorf("unknown request type: %q", typ)
}

func main() {
	data := []byte(`{"type": "start", "user": "joe", "count": 7}`)
	if err := handleRequest(data); err != nil {
		fmt.Println("ERROR:", err)
	}
	// start: main.StartJob{Type:"start", User:"joe", Count:7}

	data = []byte(`{"type": "status", "id": "seven"}`)
	if err := handleRequest(data); err != nil {
		fmt.Println("ERROR:", err)
	}
	// status: main.JobStatus{Type:"status", ID:"seven"}
}
