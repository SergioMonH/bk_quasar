package services

import "prueba_meli/internal/domain/models"

// SateliteService is the port for the satelite service
//
//go:generate mockery --name SateliteService --output ../../infrastructure/services/mocks --filename mock_satelite_service.go
type SateliteService interface {
	SaveSatelite(satelites models.Satellite) error
	GetSateliteByName(sateliteName string) (models.Satellite, error)
	GetLocation(distances ...float32) (x, y float32)
	ShowMessagesAndLocation(satelites []models.Satellite) (models.MessageAndLocation, error)
	GetMessage(messages ...[]string) string
	GetAllSatelites() ([]models.Satellite, error)
}
