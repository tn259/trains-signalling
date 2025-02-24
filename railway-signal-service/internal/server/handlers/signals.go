package handlers

import (
	"crosstech-hw/railway-signal-service/internal/database"

	"github.com/labstack/echo/v4"
)

type Signals struct {
	db *database.DB
}

func NewSignals(db *database.DB) *Signals {
	return &Signals{
		db: db,
	}
}

func (s *Signals) Get(c echo.Context) error {
	return nil
}

func (s *Signals) GetOne(c echo.Context) error {
	return nil
}

func (s *Signals) Create(c echo.Context) error {
	return nil
}

func (s *Signals) Update(c echo.Context) error {
	return nil
}

func (s *Signals) Delete(c echo.Context) error {
	return nil
}
