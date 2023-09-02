package rest

import "prueba_meli/internal/domain/models"

// SateliteResponse is the response for the Satelite endpoint
type SateliteResponse struct {
	Position models.Location `json:"position"`
	Message  string          `json:"message"`
}

// GetSateliteResponse is the response for the GetSatelite endpoint
type GetSateliteResponse struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
