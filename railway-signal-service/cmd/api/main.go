package main

import (
	"context"
	"fmt"
	"log"

	"crosstech-hw/railway-signal-service/internal/config"

	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return
	}

	opts := &pg.Options{
		Addr:     cfg.PG_ADDR,
		User:     cfg.PG_USER,
		Password: cfg.PG_PASSWORD,
		Database: cfg.PG_APP_NAME,
	}

	db := pg.Connect(opts)
	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

func loadConfig() (*config.Config, error) {
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
