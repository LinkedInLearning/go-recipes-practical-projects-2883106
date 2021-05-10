package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func postMetric(m Metric) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const url = "https://httpbin.org/post"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	const maxSize = 1 << 20 // 1MB
	r := io.LimitReader(resp.Body, maxSize)
	var reply struct {
		JSON Metric
	}
	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return err
	}
	log.Printf("GOT: %+v\n", reply.JSON)
	return nil
}

func main() {
	m := Metric{
		Time:   time.Now(),
		CPU:    0.23,
		Memory: 87.32,
	}
	if err := postMetric(m); err != nil {
		log.Fatal(err)
	}
	// 2021/04/30 17:53:15 GOT: {Time:2021-04-30 17:53:14.437272 +0300 IDT CPU:0.23 Memory:87.32}
}
