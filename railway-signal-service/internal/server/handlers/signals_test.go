package handlers

import (
	"bytes"
	"crosstech-hw/railway-signal-service/internal/database"
	"crosstech-hw/railway-signal-service/internal/server/handlers/handlerstest"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

var signal1 = database.Signal{
	ID:   1,
	Name: "signal1",
}

func TestSignals_Get(t *testing.T) {
	d := handlerstest.NewFakeDao()
	s := NewSignals(d)

	setupContext := func() echo.Context {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/signals", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return c
	}

	t.Run("Success", func(t *testing.T) {
		// GIVEN
		c := setupContext()

		// WHEN
		_ = s.Get(c)

		// THEN
		if c.Response().Status != http.StatusOK {
			t.Fatalf("expected status 200 OK, got %v", c.Response().Status)
		}
	})
	t.Run("Error", func(t *testing.T) {
		// GIVEN
		d.SignalsError = fmt.Errorf("error")
		c := setupContext()

		// WHEN
		_ = s.Get(c)

		// THEN
		if c.Response().Status != http.StatusInternalServerError {
			t.Fatalf("expected status 500 Internal Server Error, got %v", c.Response().Status)
		}
	})
}

func TestSignals_GetOne(t *testing.T) {
	d := handlerstest.NewFakeDao()
	s := NewSignals(d)

	setupContext := func(id string) echo.Context {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/signals/"+id, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id)
		return c
	}

	t.Run("Success", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = &signal1
		c := setupContext("1")

		// WHEN
		_ = s.GetOne(c)

		// THEN
		if c.Response().Status != http.StatusOK {
			t.Fatalf("expected status 200 OK, got %v", c.Response().Status)
		}
	})
	t.Run("ID parse error", func(t *testing.T) {
		// GIVEN
		c := setupContext("invalid")

		// WHEN
		_ = s.GetOne(c)

		// THEN
		if c.Response().Status != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %v", c.Response().Status)
		}
	})
	t.Run("Internal error", func(t *testing.T) {
		// GIVEN
		d.SignalErrorsByID[1] = fmt.Errorf("error")
		c := setupContext("1")

		// WHEN
		_ = s.GetOne(c)

		// THEN
		if c.Response().Status != http.StatusInternalServerError {
			t.Fatalf("expected status 500 Internal Server Error, got %v", c.Response().Status)
		}
	})
	t.Run("Not found error", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = nil
		d.SignalErrorsByID[1] = nil
		c := setupContext("1")

		// WHEN
		_ = s.GetOne(c)

		// THEN
		if c.Response().Status != http.StatusNotFound {
			t.Fatalf("expected status 404 Not Found, got %v", c.Response().Status)
		}
	})
}

func TestSignals_Create(t *testing.T) {
	d := handlerstest.NewFakeDao()
	s := NewSignals(d)

	setupContext := func(s *database.Signal) echo.Context {
		e := echo.New()
		body := io.Reader(nil)
		if s != nil {
			j, err := json.Marshal(s)
			if err != nil {
				t.Fatalf("failed to marshal signal: %v", err)
			}
			body = bytes.NewReader(j)
		}
		req := httptest.NewRequest(http.MethodPost, "/signals", body)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c
	}
	t.Run("Success", func(t *testing.T) {
		// GIVEN
		c := setupContext(&signal1)

		// WHEN
		_ = s.Create(c)

		// THEN
		if c.Response().Status != http.StatusCreated {
			t.Fatalf("expected status 201 Created, got %v", c.Response().Status)
		}
	})
	t.Run("Bind error", func(t *testing.T) {
		// GIVEN
		c := setupContext(nil)

		// WHEN
		_ = s.Create(c)

		// THEN
		if c.Response().Status != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %v", c.Response().Status)
		}
	})
	t.Run("Internal error", func(t *testing.T) {
		// GIVEN
		d.CreateSignalError = fmt.Errorf("error")
		c := setupContext(&signal1)

		// WHEN
		_ = s.Create(c)

		// THEN
		if c.Response().Status != http.StatusInternalServerError {
			t.Fatalf("expected status 500 Internal Server Error, got %v", c.Response().Status)
		}
	})
}

func TestSignals_Update(t *testing.T) {
	d := handlerstest.NewFakeDao()
	s := NewSignals(d)

	setupContext := func(id string, s *database.Signal) echo.Context {
		e := echo.New()
		body := io.Reader(nil)
		if s != nil {
			jsonBody, err := json.Marshal(s)
			if err != nil {
				t.Fatalf("failed to marshal signal: %v", err)
			}
			body = bytes.NewReader(jsonBody)
		}
		req := httptest.NewRequest(http.MethodPut, "/signals/"+id, body)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id)
		c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c
	}

	t.Run("Success", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = &signal1
		c := setupContext("1", &signal1)

		// WHEN
		_ = s.Update(c)

		// THEN
		if c.Response().Status != http.StatusOK {
			t.Fatalf("expected status 200 OK, got %v", c.Response().Status)
		}
	})
	t.Run("ID parse error", func(t *testing.T) {
		// GIVEN
		c := setupContext("invalid", &signal1)

		// WHEN
		_ = s.Update(c)

		// THEN
		if c.Response().Status != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %v", c.Response().Status)
		}
	})
	t.Run("Bind error", func(t *testing.T) {
		// GIVEN
		c := setupContext("1", nil)

		// WHEN
		_ = s.Update(c)

		// THEN
		if c.Response().Status != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %v", c.Response().Status)
		}
	})
	t.Run("Internal error", func(t *testing.T) {
		// GIVEN
		d.SignalErrorsByID[1] = fmt.Errorf("error")
		c := setupContext("1", &signal1)

		// WHEN
		_ = s.Update(c)

		// THEN
		if c.Response().Status != http.StatusInternalServerError {
			t.Fatalf("expected status 500 Internal Server Error, got %v", c.Response().Status)
		}
	})
	t.Run("Not found error", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = nil
		d.SignalErrorsByID[1] = nil
		c := setupContext("1", &signal1)

		// WHEN
		_ = s.Update(c)

		// THEN
		if c.Response().Status != http.StatusNotFound {
			t.Fatalf("expected status 404 Not Found, got %v", c.Response().Status)
		}
	})
}

func TestSignals_Delete(t *testing.T) {
	d := handlerstest.NewFakeDao()
	s := NewSignals(d)

	setupContext := func(id string) echo.Context {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/signals/"+id, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id)
		return c
	}

	t.Run("Success", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = &signal1
		c := setupContext("1")

		// WHEN
		_ = s.Delete(c)

		// THEN
		if c.Response().Status != http.StatusNoContent {
			t.Fatalf("expected status 200 OK, got %v", c.Response().Status)
		}
	})
	t.Run("ID parse error", func(t *testing.T) {
		// GIVEN
		c := setupContext("invalid")

		// WHEN
		_ = s.Delete(c)

		// THEN
		if c.Response().Status != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %v", c.Response().Status)
		}
	})
	t.Run("Internal error", func(t *testing.T) {
		// GIVEN
		d.SignalErrorsByID[1] = fmt.Errorf("error")
		c := setupContext("1")

		// WHEN
		_ = s.Delete(c)

		// THEN
		if c.Response().Status != http.StatusInternalServerError {
			t.Fatalf("expected status 500 Internal Server Error, got %v", c.Response().Status)
		}
	})
	t.Run("Not found error", func(t *testing.T) {
		// GIVEN
		d.SignalsByID[1] = nil
		d.SignalErrorsByID[1] = nil
		c := setupContext("1")

		// WHEN
		_ = s.Delete(c)

		// THEN
		if c.Response().Status != http.StatusNotFound {
			t.Fatalf("expected status 404 Not Found, got %v", c.Response().Status)
		}
	})
}
