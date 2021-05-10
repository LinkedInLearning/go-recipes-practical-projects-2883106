package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// laggingStations return stations that are lagging in their check time
func laggingStations(r io.Reader, timeout time.Duration) ([]string, error) {
	var reply struct {
		LastCheckTime string
		Stations      []struct {
			Name      string
			Status    string
			LastCheck struct {
				Time string
			}
		}
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}

	checkTime, err := parseTime(reply.LastCheckTime)
	if err != nil {
		return nil, err
	}

	var lagging []string
	for _, station := range reply.Stations {
		if station.Status != "Active" {
			continue
		}
		lastCheck, err := parseTime(station.LastCheck.Time)
		if err != nil {
			return nil, err
		}
		if checkTime.Sub(lastCheck) > timeout {
			lagging = append(lagging, station.Name)
		}
	}

	return lagging, nil
}

func parseTime(ts string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 PM", ts)
}

func main() {
	file, err := os.Open("stations.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lagging, err := laggingStations(file, time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range lagging {
		fmt.Println(name)
	}
	// station 3
}
