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

func monthDistance(month time.Time) (float64, error) {
	totalDistance := 0.0
	date := month
	for date.Month() == month.Month() {
		url := fmt.Sprintf("http://localhost:8080/%s", date.Format("2006-01-02"))
		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}

		if resp.StatusCode != http.StatusOK {
			return 0, fmt.Errorf("bad status: %d %s", resp.Request.Response.StatusCode, resp.Status)
		}

		defer resp.Body.Close()
		dist, err := dayDistance(resp.Body)
		if err != nil {
			return 0, err
		}
		totalDistance += dist
		date = date.Add(24 * time.Hour)
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
