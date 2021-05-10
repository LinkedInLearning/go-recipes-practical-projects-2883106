package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func dayDistance(r io.Reader) (float64, error) {
	rdr := csv.NewReader(r)
	total, lNum := 0.0, 0
	for {
		//2021-01-02T23:58:36,2021-01-02T23:58:40,3.41,1
		fields, err := rdr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		lNum++
		if lNum == 1 {
			continue // skip header
		}

		dist, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return 0, err
		}

		total += dist
	}

	return total, nil
}

type result struct {
	date time.Time
	dist float64
	err  error
}

func dateWorker(date time.Time, ch chan<- result) {
	res := result{date: date}
	defer func() {
		ch <- res
	}()

	url := fmt.Sprintf("http://localhost:8080/%s", date.Format("2006-01-02"))
	resp, err := http.Get(url)
	if err != nil {
		res.err = err
		return
	}

	if resp.StatusCode != http.StatusOK {
		res.err = fmt.Errorf("bad status: %d %s", resp.Request.Response.StatusCode, resp.Status)
		return
	}

	defer resp.Body.Close()
	res.dist, res.err = dayDistance(resp.Body)
}

func monthDistance(month time.Time) (float64, error) {
	numWorkers, ch := 0, make(chan result)
	date := month
	for date.Month() == month.Month() {
		go dateWorker(date, ch)
		numWorkers++
		date = date.Add(24 * time.Hour)
	}

	totalDistance := 0.0
	for i := 0; i < numWorkers; i++ {
		res := <-ch
		if res.err != nil {
			return 0, res.err
		}
		totalDistance += res.dist
	}

	return totalDistance, nil
}

func main() {
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)

	start := time.Now()
	dist, err := monthDistance(month)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)
	fmt.Printf("distance=%.2f, duration=%v\n", dist, duration)
}
