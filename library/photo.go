package library

import (
	"time"

	"trailimage.com/geo/GeoJSON"
)

type (
	Photo struct {
		ID            string
		Index         int
		SourceUrl     string
		Title         string
		Caption       Caption
		DateTaken     time.Time
		IsPrimary     bool
		Tags          []string
		IsOutlierDate bool
		Size          SizeChoice
	}

	SizeChoice struct {
		Preview PhotoSize
		Normal  PhotoSize
		Big     PhotoSize
	}

	PhotoSize struct {
		Url    string
		Width  int
		Height int
	}

	Coordinate struct {
		Latitude  int
		Longitude int
	}
)

// ToGeoJSON creates a GeoJSON feature for the photo.
func (p *Photo) ToGeoJSON() *geoJSON.Feature {
	return &geoJSON.Feature{
		ID: p.ID,
	}
}

func (p *Photo) GetExif() {

}

func (p *Photo) TagList() string {
	return ""
}

func (s *PhotoSize) IsEmpty() bool {
	return s.Url == "" && s.Width == 0
}
