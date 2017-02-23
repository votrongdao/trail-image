package geoJSON

type (
	BoundingBox [4]float64

	Feature struct {
		ID         string            `json:"id"`
		Geometry   Geometry          `json:"geometry"`
		Properties map[string]string `json:"properties"`
	}

	FeatureCollection struct {
		Type        string      `json:"type"`
		BoundingBox BoundingBox `json:"bbox"`
		Features    []*Feature  `json:"features"`
	}

	Point struct {
		Geometry
	}

	Line struct {
		Geometry
	}

	MultiLine struct {
	}

	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}
)

func FeatureFromPhoto(p *library.Photo) *Feature {
	return &Feature{
		ID: p.ID,
	}
}
