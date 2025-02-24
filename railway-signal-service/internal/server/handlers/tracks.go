package handlers

import (
	"crosstech-hw/railway-signal-service/internal/database"

	"github.com/labstack/echo/v4"
)

type Tracks struct {
	db *database.DB
}

func NewTracks(db *database.DB) *Tracks {
	return &Tracks{
		db: db,
	}
}

func (t *Tracks) Get(c echo.Context) error {
	return nil
}

func (t *Tracks) GetOne(c echo.Context) error {
	return nil
}

func (t *Tracks) Create(c echo.Context) error {
	return nil
}

func (t *Tracks) Update(c echo.Context) error {
	return nil
}

func (t *Tracks) Delete(c echo.Context) error {
	return nil
}
