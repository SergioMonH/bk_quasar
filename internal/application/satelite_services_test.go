package application_test

import (
	"encoding/json"
	"errors"
	"prueba_meli/internal/application"
	"prueba_meli/internal/domain/models"
	"prueba_meli/internal/infrastructure/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSateliteService_GetAllSatelites(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

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

	var satelites []models.Satellite
	json.Unmarshal(data, &satelites)

	sateliteRepository.On("GetAllSatelites").Return(satelites, nil)

	satelites, err := sateliteService.GetAllSatelites()
	assert.NoError(err)
	assert.Equal(3, len(satelites))

}

func TestSateliteService_GetAllSatelites_Error(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

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

	var satelites []models.Satellite
	json.Unmarshal(data, &satelites)

	sateliteRepository.On("GetAllSatelites").Return([]models.Satellite{}, errors.New("error"))

	satelites, err := sateliteService.GetAllSatelites()
	assert.Error(err)

}

func TestSateliteService_GetSateliteByName(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	data := []byte(`
	{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
	}`)

	var satelite models.Satellite
	json.Unmarshal(data, &satelite)

	sateliteRepository.On("GetSateliteByName", "kenobi").Return(satelite, nil)

	satelite, err := sateliteService.GetSateliteByName("kenobi")
	assert.NoError(err)
	assert.Equal("kenobi", satelite.Name)

}

func TestSateliteService_GetSateliteByName_Error(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	data := []byte(`
	{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
	}`)

	var satelite models.Satellite
	json.Unmarshal(data, &satelite)

	sateliteRepository.On("GetSateliteByName", "kenobi").Return(models.Satellite{}, errors.New("error"))

	satelite, err := sateliteService.GetSateliteByName("kenobi")
	assert.Error(err)

}

func TestSateliteService_SaveSatelite(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	data := []byte(`
	{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
	}`)

	var satelite models.Satellite
	json.Unmarshal(data, &satelite)

	sateliteRepository.On("SaveSatelite", satelite).Return(nil)

	err := sateliteService.SaveSatelite(satelite)
	assert.NoError(err)

}

func TestSateliteService_SaveSatelite_Error(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	data := []byte(`
	{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
	}`)

	var satelite models.Satellite
	json.Unmarshal(data, &satelite)

	sateliteRepository.On("SaveSatelite", satelite).Return(errors.New("error"))

	err := sateliteService.SaveSatelite(satelite)
	assert.Error(err)

}

func TestSateliteService_GetLocation(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	x, y := sateliteService.GetLocation(-500, -200, 100.0, 100, -100, 115.5, 500, 100, 142.7)
	assert.Equal(float32(-487.2859), x)
	assert.Equal(float32(1557.0142), y)

}

func TestSateliteService_GetLocation_Error(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	x, y := sateliteService.GetLocation(-500, -200, 100.0, 100, -100, 115.5, 500, 100)
	assert.Equal(float32(0), x)
	assert.Equal(float32(0), y)

}

func TestSateliteService_GetMessage(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

	message := sateliteService.GetMessage([]string{"este", "", "", "mensaje", ""}, []string{"", "es", "", "", "secreto"}, []string{"este", "", "un", "", ""})
	assert.Equal("este es un mensaje secreto", message)

}

func TestSateliteService_ShowMessagesAndLocation(t *testing.T) {
	assert := assert.New(t)

	sateliteRepository := mocks.NewSateliteRepository(t)
	sateliteService := application.NewSateliteService(sateliteRepository)

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

	var satelites []models.Satellite
	json.Unmarshal(data, &satelites)

	response, err := sateliteService.ShowMessagesAndLocation(satelites)
	assert.NoError(err)
	assert.Equal("este es un mensaje secreto", response.Message)

}
