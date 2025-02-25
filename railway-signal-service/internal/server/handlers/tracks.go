package handlers

import (
	"crosstech-hw/railway-signal-service/internal/database"
	"fmt"
	"net/http"
	"strconv"

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
	ts, err := t.d.Tracks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to get tracks: %v", err)))
	}
	return c.JSON(http.StatusOK, ts)
}

func (t *Tracks) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	tByID, err := t.d.TrackByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to get track with id %d: %v", id, err)))
	}
	if tByID == nil {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("track with id %d not found", id)))
	}
	return c.JSON(http.StatusOK, tByID)
}

func (t *Tracks) Create(c echo.Context) error {
	var track *database.Track
	if err := c.Bind(&track); err != nil || track == nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to bind track: %v", err)))
	}
	if err := t.d.CreateTrack(track); err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to create track: %v", err)))
	}
	return c.JSON(http.StatusCreated, track)
}

func (t *Tracks) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	var track *database.Track
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to bind track: %v", err)))
	}
	track.ID = id
	updated, err := t.d.UpdateTrack(track)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to update track: %v", err)))
	}
	if !updated {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("track with id %d not found", id)))
	}
	return c.JSON(http.StatusOK, track)
}

func (t *Tracks) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse(fmt.Errorf("failed to parse id: %v", err)))
	}
	deleted, err := t.d.DeleteTrack(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse(fmt.Errorf("failed to delete track: %v", err)))
	}
	if !deleted {
		return c.JSON(http.StatusNotFound, newErrorResponse(fmt.Errorf("track with id %d not found", id)))
	}
	return c.NoContent(http.StatusNoContent)
}
