package main

import (
	"log"

	"crosstech-hw/railway-signal-service/internal/config"
	"crosstech-hw/railway-signal-service/internal/database"
	"crosstech-hw/railway-signal-service/internal/server"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return
	}

	db, err := database.Connect(*cfg)
	if err != nil {
		log.Fatal("db.Connect():", err)
		return
	}
	defer db.Close()

	server.New(db)
}

func loadConfig() (*config.Config, error) {
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
