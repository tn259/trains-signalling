package main

import (
	"log"

	"crosstech-hw/railway-signal-service/internal/config"
	"crosstech-hw/railway-signal-service/internal/database"
	"crosstech-hw/railway-signal-service/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return
	}

	// Connect to the database
	db, err := database.Connect(*cfg)
	if err != nil {
		log.Fatal("db.Connect():", err)
		return
	}
	defer db.Close()

	err = db.CreateSchema()
	if err != nil {
		log.Fatal("db.CreateSchema():", err)
		return
	}

	// Start the server
	dao := database.NewPGDao(db)
	server.New(dao)
}
