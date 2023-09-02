package models

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitesGroup struct {
	SatellitesList []Satellite `json:"satellites"`
}

// Location is the location of a satelite
type Location struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// MessageAndLocation is the response for the GetMessageAndLocation endpoint
type MessageAndLocation struct {
	Location Location `json:"location"`
	Message  string   `json:"message"`
}

// NewSatelite returns a new satelite
func NewSatelite(name string, dist float32, msg []string) (*Satellite, error) {
	return &Satellite{
		Name:     name,
		Distance: dist,
		Message:  msg,
	}, nil
}
