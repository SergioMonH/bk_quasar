package rest_test

import (
	rest "prueba_meli/internal/infrastructure/rest/echo"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSatelliteRoutes_Register(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	h := rest.NewSateliteHandler(nil)

	satelliteRoutes := rest.NewSatelliteRoutes(e, h)

	satelliteRoutes.GetSatelite()

	assert.Len(e.Routes(), 1)
}

func TestSatelliteRoutes_SaveSatelite(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	h := rest.NewSateliteHandler(nil)

	satelliteRoutes := rest.NewSatelliteRoutes(e, h)

	satelliteRoutes.SaveSatelite()

	assert.Len(e.Routes(), 1)
}

func TestSatelliteRoutes_SateliteMessageAndLocation(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	h := rest.NewSateliteHandler(nil)

	satelliteRoutes := rest.NewSatelliteRoutes(e, h)

	satelliteRoutes.SateliteMessageAndLocation()

	assert.Len(e.Routes(), 1)
}
