package handlers

import (
	"crosstech-hw/railway-signal-service/internal/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Signals struct {
	d database.Dao
}

func NewSignals(d database.Dao) *Signals {
	return &Signals{
		d: d,
	}
}

func (s *Signals) Get(c echo.Context) error {
	ss, err := s.d.Signals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to get signals: %v", err)))
	}
	return c.JSON(http.StatusOK, ss)
}

func (s *Signals) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	sByID, err := s.d.SignalByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to get signal with id %d: %v", id, err)))
	}
	if sByID == nil {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("signal with id %d not found", id)))
	}
	return c.JSON(http.StatusOK, sByID)
}

func (s *Signals) Create(c echo.Context) error {
	var signal *database.Signal
	if err := c.Bind(&signal); err != nil || signal == nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to bind signal: %v", err)))
	}
	if err := s.d.CreateSignal(signal); err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to create signal: %v", err)))
	}
	return c.JSON(http.StatusCreated, signal)
}

func (s *Signals) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	var signal *database.Signal
	if err := c.Bind(&signal); err != nil || signal == nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to bind signal: %v", err)))
	}
	signal.ID = id
	updated, err := s.d.UpdateSignal(signal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to update signal with id %d: %v", id, err)))
	}
	if !updated {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("signal with id %d not found", id)))
	}
	return c.JSON(http.StatusOK, signal)
}

func (s *Signals) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	deleted, err := s.d.DeleteSignal(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to delete signal with id %d: %v", id, err)))
	}
	if !deleted {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("signal with id %d not found", id)))
	}
	return c.NoContent(http.StatusNoContent)
}
