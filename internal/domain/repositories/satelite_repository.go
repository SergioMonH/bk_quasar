package repositories

import "prueba_meli/internal/domain/models"

// SateliteRepository is the port for the satelite repository
//
//go:generate mockery --name SateliteRepository --output ../../infrastructure/repositories/mocks --filename mock_satelite_repository.go
type SateliteRepository interface {
	SaveSatelite(satelites models.Satellite) error
	GetSateliteByName(sateliteName string) (models.Satellite, error)
	GetAllSatelites() ([]models.Satellite, error)
}
