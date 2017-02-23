package gpx

// https://golang.org/src/encoding/xml/example_test.go
type (
	File struct {
		Name   string   `xml:"metadata>name"`
		Tracks []*Track `xml:"trk>"`
	}

	Track struct {
		Name     string     `xml:"name>"`
		Segments []*Segment `xml:"trkseg>"`
	}

	Segment struct {
		Points []*Point `xml:"trkpt>"`
	}

	Point struct {
		Latitude  float64 `xml:"lat,attr"`
		Longitude float64 `xml:"lon,attr"`
		Elevation float64 `xml:"ele>"`
		Time      string  `xml:"time>"`
	}
)
