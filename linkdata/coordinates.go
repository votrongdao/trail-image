package linkdata

type Coordinates struct {
	Thing
	Elevation  float32 `json:"elevation"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
	PostalCode string  `json:"postalCode"`
}
