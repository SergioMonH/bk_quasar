package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"prueba_meli/internal/application"
	repositories "prueba_meli/internal/infrastructure/repositories/json"
	rest "prueba_meli/internal/infrastructure/rest/echo"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	jsonSateliteRepository := repositories.NewJsonSateliteRepository("satelites.json")

	sateliteService := application.NewSateliteService(jsonSateliteRepository)

	sateliteHandler := rest.NewSateliteHandler(sateliteService)

	sateliteRoutes := rest.NewSatelliteRoutes(e, sateliteHandler)
	sateliteRoutes.GetSatelite()
	sateliteRoutes.SateliteMessageAndLocation()
	sateliteRoutes.SaveSatelite()

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful shutdown
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
