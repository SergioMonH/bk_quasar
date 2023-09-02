package repositories_test

import (
	"encoding/json"
	"prueba_meli/internal/domain/models"
	repositories "prueba_meli/internal/infrastructure/repositories/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonRepository_SaveSatelite(t *testing.T) {
	assert := assert.New(t)

	repository := repositories.NewJsonSateliteRepository("../satelite_test.json")

	data := []byte(`
	{
		
			"name": "kenobi",
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]
		
	}`)

	var satelite models.Satellite
	json.Unmarshal(data, &satelite)

	err := repository.SaveSatelite(satelite)
	assert.NoError(err)

}

func TestJsonRepository_GetSateliteByName(t *testing.T) {
	assert := assert.New(t)

	repository := repositories.NewJsonSateliteRepository("../satelite_test.json")

	satelite, err := repository.GetSateliteByName("kenobi")
	assert.NoError(err)
	assert.Equal("kenobi", satelite.Name)
}

func TestJsonRepository_GetAllSatelites(t *testing.T) {
	assert := assert.New(t)

	repository := repositories.NewJsonSateliteRepository("../satelite_test.json")

	satelites, err := repository.GetAllSatelites()
	assert.NoError(err)
	assert.Equal(3, len(satelites))
}
