package server

import (
	"crosstech-hw/railway-signal-service/internal/database"
	"crosstech-hw/railway-signal-service/internal/server/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// New creates a new http server with endpoints for signals and tracks
func New(d database.Dao) {
	r := echo.New()

	sh := handlers.NewSignals(d)
	th := handlers.NewTracks(d)

	// Endpoints - signals
	r.GET("/signals", sh.Get)
	r.GET("/signals/:id", sh.GetOne)
	r.POST("/signals", sh.Create)
	r.PUT("/signals/:id", sh.Update)
	r.DELETE("/signals/:id", sh.Delete)

	// Endpoints - tracks
	r.GET("/tracks", th.Get)
	r.GET("/tracks/:id", th.GetOne)
	r.POST("/tracks", th.Create)
	r.PUT("/tracks/:id", th.Update)
	r.DELETE("/tracks/:id", th.Delete)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	log.Println("Server started successfully")
}
