package main

import (
	"crosstech-hw/railway-signal-service/internal/config"
	"crosstech-hw/railway-signal-service/internal/database"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
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
	fmt.Println("JSON deserialize successful")

	// Connect to the database and create the schema if necessary
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db, err := database.Connect(*cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.CreateSchema()
	if err != nil {
		panic(err)
	}

	// insert the data into the database
	insertData(db, tracks)
}

func insertData(db *database.DB, t []*Track) {
	fmt.Println("Inserting data into the database")
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				panic(err)
			}
		}
	}()

	for _, track := range t {
		// Insert track
		_, err := tx.Exec("INSERT INTO tracks (id, source, target) VALUES (?, ?, ?)", track.TrackID, track.Source, track.Target)
		if err != nil {
			panic(err)
		}

		for _, signal := range track.SignalIDs {
			// Upsert Signal
			_, err := tx.Exec("INSERT INTO signals (id, name) VALUES (?, ?) ON CONFLICT (id) DO NOTHING", signal.SignalID, signal.SignalName)
			if err != nil {
				panic(err)
			}

			// Upsert ELR
			var elrID int
			// Use QueryOne to get the id of the elr
			// Need to user UPDATE SET instead of DO NOTHING in order to get the id returned
			// This avoids doing a SELECT below to get the ELR ID
			_, err = tx.QueryOne(pg.Scan(&elrID), "INSERT INTO elrs (name) VALUES (?) ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id", signal.ELR)
			if err != nil {
				panic(err)
			}

			// Insert TrackSignal with FK constraints
			_, err = tx.Exec("INSERT INTO track_signals (mileage, elr_id, signal_id, track_id) VALUES (?, ?, ?, ?)", signal.Mileage, elrID, signal.SignalID, track.TrackID)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Data inserted successfully")
}
