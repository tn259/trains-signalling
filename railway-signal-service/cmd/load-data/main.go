package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// This script loads the json data provided into the database

type Signal struct {
	SignalID   int      `json:"signal_id"`
	SignalName string   `json:"signal_name"`
	ELR        string   `json:"elr"`
	Mileage    *float64 `json:"mileage"` // Some values are NaN
}

type Track struct {
	TrackID   int       `json:"track_id"`
	Source    string    `json:"source"`
	Target    string    `json:"target"`
	SignalIDs []*Signal `json:"signal_ids"`
}

var jsonFile string

func parseCLIOptions() {
	flag.StringVar(&jsonFile, "jsonFile", "data_pretty.json", "JSON file")
	flag.Parse()
}

func main() {
	parseCLIOptions()

	// read the json file
	file, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// decode the json file
	decoder := json.NewDecoder(file)
	var tracks []*Track
	err = decoder.Decode(&tracks)
	if err != nil {
		panic(err)
	}

	fmt.Print(*tracks[0])
	// TODO load the data to the database or via the REST API?
}
