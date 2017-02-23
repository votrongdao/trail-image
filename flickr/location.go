package flickr

type (
	Place struct {
		PlaceID        string `json:"place_id"`
		WhereOnEarthID string `json:"woeid"`
	}

	Coordinate struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		Accuracy  uint    `json:"accuracy"`
	}

	Location struct {
		Coordinate
		Context int   `json:"context"`
		County  Place `json:"county"`
		Region  Place `json:"region"`
		Country Place `json:"country"`
	}

	ServerLocation struct {
		Secret string `json:"secret"`
		Server string `json:"server"`
		Farm   uint   `json:"farm"`
	}
)
