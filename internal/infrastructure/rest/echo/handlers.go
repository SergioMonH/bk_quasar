package rest

import (
	"net/http"
	"prueba_meli/internal/domain/models"
	"prueba_meli/internal/domain/services"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type sateliteHandler struct {
	sateliteService services.SateliteService
}

// NewSateliteHandler creates a new consumption handler
func NewSateliteHandler(sateliteService services.SateliteService) *sateliteHandler {
	return &sateliteHandler{
		sateliteService: sateliteService,
	}
}

// ShowMessageAndLocation is the handler for the topsecret satelite resource
func (h *sateliteHandler) ShowMessageAndLocation(c echo.Context) error {

	satelitesRequest := models.SatellitesGroup{}

	validate := validator.New()

	if err := c.Bind(&satelitesRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := validate.Struct(satelitesRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sort.Slice(satelitesRequest.SatellitesList, func(i, j int) bool {
		return satelitesRequest.SatellitesList[i].Distance < satelitesRequest.SatellitesList[j].Distance
	})

	locationAndMessage, err := h.sateliteService.ShowMessagesAndLocation(satelitesRequest.SatellitesList)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, locationAndMessage)
}

// GetSatelite is the handler for the topsecret_split/:satelite_name resource
func (h *sateliteHandler) GetSatelite(c echo.Context) error {

	sateliteName := c.Param("satellite_name")

	satelite, err := h.sateliteService.GetSateliteByName(sateliteName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sateliteResponse := GetSateliteResponse{
		Distance: satelite.Distance,
		Message:  satelite.Message,
	}

	return c.JSON(http.StatusOK, sateliteResponse)
}

// PostSatelite is the handler for the topsecret_split resource
func (h *sateliteHandler) PostSatelite(c echo.Context) error {
	sateliteName := c.Param("satellite_name")

	satelitesRequest := GetSateliteResponse{}

	validate := validator.New()

	if err := c.Bind(&satelitesRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := validate.Struct(satelitesRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sateliteData := models.Satellite{
		Name:     sateliteName,
		Distance: satelitesRequest.Distance,
		Message:  satelitesRequest.Message,
	}

	err = h.sateliteService.SaveSatelite(sateliteData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	SateliteResponse := GetSateliteResponse{
		Distance: sateliteData.Distance,
		Message:  sateliteData.Message,
	}

	return c.JSON(http.StatusOK, SateliteResponse)
}

// GetSateliteSplit is the handler for the topsecret_split resource
func (h *sateliteHandler) GetSateliteSplit(c echo.Context) error {

	sateliteName := c.Param("satellite_name")

	satelite, err := h.sateliteService.GetSateliteByName(sateliteName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	checkWork := 0

	for _, message := range satelite.Message {
		if message != "" {
			checkWork++
		}
	}

	if checkWork < 4 {
		return echo.NewHTTPError(http.StatusInternalServerError, "Not enough data")
	}

	satelitesData, err := h.sateliteService.GetAllSatelites()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(satelitesData) < 3 {
		return echo.NewHTTPError(http.StatusInternalServerError, "Not enough data")
	}

	sort.Slice(satelitesData, func(i, j int) bool {
		return satelitesData[i].Distance < satelitesData[j].Distance
	})

	locationAndMessage, err := h.sateliteService.ShowMessagesAndLocation(satelitesData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sateliteMessage := strings.Join(satelite.Message, " ")

	locationAndMessage.Message = sateliteMessage

	return c.JSON(http.StatusOK, locationAndMessage)
}
