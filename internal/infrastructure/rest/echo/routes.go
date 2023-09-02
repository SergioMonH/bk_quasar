package rest

import "github.com/labstack/echo"

// ConsumptionRoutes is the struct that contains the routes for the consumption resource
type SateliteRoutes struct {
	e *echo.Echo
	h *sateliteHandler
}

// NewConsumptionRoutes creates a new consumption routes
func NewSatelliteRoutes(e *echo.Echo, h *sateliteHandler) *SateliteRoutes {
	return &SateliteRoutes{
		e: e,
		h: h,
	}
}

// SateliteMessageAndLocation register the route for the topsecret satelite resource
func (r *SateliteRoutes) SateliteMessageAndLocation() {
	r.e.POST("/topsecret/", r.h.ShowMessageAndLocation)
}

// GetSatelite register the route for the topsecret_split/:satelite_name resource
func (r *SateliteRoutes) GetSatelite() {
	r.e.GET("/topsecret_split/:satellite_name", r.h.GetSateliteSplit)
}

// SaveSatelite register the route for the topsecret_split/:satelite_name resource
func (r *SateliteRoutes) SaveSatelite() {
	r.e.POST("/topsecret_split/:satellite_name", r.h.PostSatelite)
}
