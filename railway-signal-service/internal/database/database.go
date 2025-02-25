package database

import (
	"crosstech-hw/railway-signal-service/internal/config"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DB struct {
	*pg.DB
}

func Connect(cfg config.Config) (*DB, error) {
	opts := &pg.Options{
		Addr:     cfg.PG_ADDR,
		User:     cfg.PG_USER,
		Password: cfg.PG_PASSWORD,
		Database: cfg.PG_APP_NAME,
	}

	log.Println(opts)

	db := pg.Connect(opts)

	ctx := db.Context()
	for {
		if err := db.Ping(ctx); err == nil {
			break
		}
		// TODO add a timeout
	}

	log.Println("Connected to database successfully")

	if err := createSchema(db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}

// createSchema creates the database schema based on the defined models
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*ELR)(nil),
		(*Signal)(nil),
		(*Track)(nil),
		(*TrackSignal)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return fmt.Errorf("failed to create table %v: %w", model, err)
		}
	}

	log.Print("Database schema created successfully")
	return nil
}
