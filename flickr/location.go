package flickr

type (
	Place struct {
		PlaceID        string `json:"place_id"`
		WhereOnEarthID string `json:"woeid"`
	}

	Location struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		Accuracy  int     `json:"accuracy"`
		Context   int     `json:"context"`
		County    Place   `json:"county"`
		Region    Place   `json:"region"`
		Country   Place   `json:"country"`
	}

	FarmLocation struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
		Server string `json:"server"`
		Farm   uint   `json:"farm"`
	}

	
)
