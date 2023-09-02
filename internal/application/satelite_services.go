package application

import (
	"prueba_meli/internal/domain/models"
	"prueba_meli/internal/domain/repositories"
	"prueba_meli/internal/util"
)

type SateliteService struct {
	repository repositories.SateliteRepository
}

// GetAllSatelites returns all satelites
func (s *SateliteService) GetAllSatelites() ([]models.Satellite, error) {
	satelites, err := s.repository.GetAllSatelites()
	if err != nil {
		return nil, err
	}

	return satelites, err
}

// GetSateiteByName returns a satelite by name
func (s *SateliteService) GetSateliteByName(sateliteName string) (models.Satellite, error) {
	satelite, err := s.repository.GetSateliteByName(sateliteName)
	if err != nil {
		return models.Satellite{}, err
	}

	return satelite, err
}

// SaveSatelite saves a satelite at json file
func (s *SateliteService) SaveSatelite(satelites models.Satellite) error {
	err := s.repository.SaveSatelite(satelites)
	if err != nil {
		return err
	}

	return nil
}

// GetLocation returns the location of the ship by trilateration
func (s *SateliteService) GetLocation(distances ...float32) (x, y float32) {
	if len(distances) != 9 {
		return 0, 0
	}

	x1 := distances[0]
	y1 := distances[1]
	r1 := distances[2]

	x2 := distances[3]
	y2 := distances[4]
	r2 := distances[5]

	x3 := distances[6]
	y3 := distances[7]
	r3 := distances[8]

	x, y = util.Trilateration(x1, y1, r1, x2, y2, r2, x3, y3, r3)

	return x, y

}

// GetMesage returns the message of the ship
func (s *SateliteService) GetMessage(messages ...[]string) string {
	message := util.DiscoverMessage(messages...)

	return message
}

// ShowMessagesAndLocation returns the location and message of the ship
func (s *SateliteService) ShowMessagesAndLocation(satelites []models.Satellite) (models.MessageAndLocation, error) {
	var paramters []float32
	var messageParts [][]string

	paramters = append(paramters, -500, -200, satelites[0].Distance, 100, -100, satelites[1].Distance, 500, 100, satelites[2].Distance)

	x, y := s.GetLocation(paramters...)

	response := models.MessageAndLocation{
		Location: models.Location{
			X: x,
			Y: y,
		},
		Message: "",
	}

	for _, satelite := range satelites {
		messageParts = append(messageParts, satelite.Message)
	}

	response.Message = s.GetMessage(messageParts...)

	return response, nil

}

// NewSateliteService creates a new instance of sateliteService
func NewSateliteService(repository repositories.SateliteRepository) *SateliteService {
	return &SateliteService{
		repository: repository,
	}
}
