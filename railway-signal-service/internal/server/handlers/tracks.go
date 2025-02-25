package handlers

import (
	"crosstech-hw/railway-signal-service/internal/database"

	"github.com/labstack/echo/v4"
)

type Tracks struct {
	d database.Dao
}

func NewTracks(d database.Dao) *Tracks {
	return &Tracks{
		d: d,
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
