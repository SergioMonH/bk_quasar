package repositories

import (
	"encoding/json"

	"os"
	"prueba_meli/internal/domain/models"
	"prueba_meli/internal/util"
)

// JsonSateliteRepository is the port for the satelite repository
type JsonSateliteRepository struct {
	file string
}

// NewJsonSateliteRepository returns a new instance of JsonSateliteRepository
func NewJsonSateliteRepository(file string) *JsonSateliteRepository {
	return &JsonSateliteRepository{
		file: file,
	}
}

// SaveSatelite saves a satelite at json file
func (j *JsonSateliteRepository) SaveSatelite(satelite models.Satellite) error {
	var satellites []models.Satellite

	content, err := os.ReadFile(j.file)
	if err != nil {
		return err
	}

	json.Unmarshal(content, &satellites)

	satellitesMap := make(map[string]models.Satellite)
	for _, s := range satellites {
		newMessage := util.AppendMessage(s.Message, satelite.Message)
		s.Message = newMessage
		satellitesMap[s.Name] = s

	}

	satellitesMap[satelite.Name] = satelite

	satellites = make([]models.Satellite, 0, len(satellitesMap))

	for _, s := range satellitesMap {
		satellites = append(satellites, s)
	}

	data, err := json.MarshalIndent(satellites, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(j.file, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// GetSateliteByName returns a satelite by name
func (j *JsonSateliteRepository) GetSateliteByName(sateliteName string) (models.Satellite, error) {

	dataSatelites, err := os.ReadFile(j.file)
	if err != nil {
		return models.Satellite{}, err
	}

	var satelites []models.Satellite
	var sateliteResponse models.Satellite
	err = json.Unmarshal(dataSatelites, &satelites)
	if err != nil {
		return models.Satellite{}, err
	}

	for _, satelite := range satelites {
		if satelite.Name == sateliteName {
			sateliteResponse = satelite
		}
	}

	return sateliteResponse, nil

}

// GetAllSatelites returns all satelites
func (j *JsonSateliteRepository) GetAllSatelites() ([]models.Satellite, error) {
	dataSatelites, err := os.ReadFile(j.file)
	if err != nil {
		return nil, err
	}

	var satelites []models.Satellite
	err = json.Unmarshal(dataSatelites, &satelites)
	if err != nil {
		return nil, err
	}

	return satelites, nil
}
