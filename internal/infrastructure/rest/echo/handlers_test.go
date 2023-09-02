package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"prueba_meli/internal/domain/models"
	rest "prueba_meli/internal/infrastructure/rest/echo"
	"prueba_meli/internal/infrastructure/services/mocks"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSateliteHandler_ShowMessageAndLocation(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	data := []byte(`
	{
		"satellites": [
			{
				"name": "kenobi",
				"distance": 100.0,
				"message": ["este", "", "", "mensaje", ""]
			},
			{
				"name": "skywalker",
				"distance": 115.5,
				"message": ["", "es", "", "", "secreto"]
			},
			{
				"name": "sato",
				"distance": 142.7,
				"message": ["este", "", "un", "", ""]
			}
		]
	}`)

	var satelites models.SatellitesGroup

	json.Unmarshal(data, &satelites)

	req := httptest.NewRequest(echo.POST, "/topsecret/", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.On("ShowMessagesAndLocation", satelites.SatellitesList).Return(models.MessageAndLocation{}, nil)
	err := handler.ShowMessageAndLocation(c)
	assert.NoError(err)

}

func TestSateliteHandler_ShowMessageAndLocationNoPayload(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	data := []byte(`
	{
		"satellites": [
			{
				"name": "kenobi",
				"distance": 100.0,
				"message": ["este", "", "", "mensaje", ""]
			},
			{
				"name": "skywalker",
				"distance": 115.5,
				"message": ["", "es", "", "", "secreto"]
			},
			{
				"name": "sato",
				"distance": 142.7,
				"message": ["este", "", "un", "", ""]
			}
		]
	}`)

	var satelites models.SatellitesGroup

	json.Unmarshal(data, &satelites)

	req := httptest.NewRequest(echo.POST, "/topsecret/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.ShowMessageAndLocation(c)
	assert.Error(err)

}

func TestSateliteHandler_PostSatelite(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	data := []byte(`
	{
		
			"name": "",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		
	}`)

	var satelites models.Satellite

	json.Unmarshal(data, &satelites)

	req := httptest.NewRequest(echo.POST, "/topsecret_split/?satelite_name=kenobi", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.On("SaveSatelite", satelites).Return(nil)

	err := handler.PostSatelite(c)
	assert.NoError(err)

}

func TestSateliteHandler_PostSateliteNoParams(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	data := []byte(`
	{
		
			"name": "",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		
	}`)

	var satelites models.Satellite

	json.Unmarshal(data, &satelites)

	req := httptest.NewRequest(echo.POST, "/topsecret_split/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.PostSatelite(c)
	assert.Error(err)

}

func TestSateliteHandler_GetSatelite(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	data := []byte(`
	{
		
			"name": "kenobi",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		
	}`)

	var satelite models.Satellite

	json.Unmarshal(data, &satelite)

	req := httptest.NewRequest(echo.GET, "/topsecret_split/?satelite_name=kenobi", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.On("GetSateliteByName", "").Return(satelite, nil)

	err := handler.GetSatelite(c)
	assert.NoError(err)

}

func TestSateliteHandler_GetSateliteSplit(t *testing.T) {
	assert := assert.New(t)

	service := mocks.NewSateliteService(t)

	handler := rest.NewSateliteHandler(service)

	e := echo.New()

	dataOneSatelite := []byte(`
	{
		
			"name": "kenobi",
			"distance": 100.0,
			"message": ["este", "es", "un", "mensaje", ""]
		
	}`)

	data := []byte(`
	[
		{
			"name": "kenobi",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		},
		{
			"name": "skywalker",
			"distance": 115.5,
			"message": ["", "es", "", "", "secreto"]
		},
		{
			"name": "sato",
			"distance": 142.7,
			"message": ["este", "", "un", "", ""]
		}
	]`)

	var satelite models.Satellite
	json.Unmarshal(dataOneSatelite, &satelite)

	var satelites []models.Satellite
	json.Unmarshal(data, &satelites)

	req := httptest.NewRequest(echo.POST, "/topsecret_split/?satelite_name=kenobi", bytes.NewBuffer(dataOneSatelite))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.On("GetSateliteByName", "").Return(satelite, nil)
	service.On("GetAllSatelites").Return(satelites, nil)
	service.On("ShowMessagesAndLocation", satelites).Return(models.MessageAndLocation{}, nil)

	err := handler.GetSateliteSplit(c)
	assert.NoError(err)

}
