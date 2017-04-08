package linkdata

type Coordinates struct {
	Thing
	Elevation  float32 `json:"elevation,omitempty"`
	LonTgitude float32 `json:"longitude,omitempty"`
	Latitude   float32 `json:"latitude,omitempty"`
	PostalCode string  `json:"postalCode,omitempty"`
}
