// Package gpx is used to parse GPX files.
package gpx

import (
	"fmt"

	"trailimage.com/geo"
	"trailimage.com/geo/GeoJSON"
)

// https://golang.org/src/encoding/xml/example_test.go
type (
	File struct {
		PrivacyCenter geo.Point
		PrivacyMiles  float64
		Name          string      `xml:"metadata>name"`
		Tracks        []*Track    `xml:"trk"`
		Routes        []*Route    `xml:"rpt"`
		Waypoints     []*Waypoint `xml:"wpt"`
	}

	Track struct {
		Name     string     `xml:"name"`
		Segments []*Segment `xml:"trkseg"`
	}

	Route struct {
		Points []*Point `xml:"rtept"`
	}

	Segment struct {
		Points []*Point `xml:"trkpt"`
	}

	Waypoint struct {
		Points []*Point `xml:"trkpt"`
	}

	Point struct {
		Latitude  float64 `xml:"lat,attr"`
		Longitude float64 `xml:"lon,attr"`
		Elevation float64 `xml:"ele"`
		Time      string  `xml:"time"`
	}
)

// ToArray converts GPX point to a point construct.
func (p *Point) ToArray(lastPoint geo.Point) geo.Point {
	time := p.Time
	ts := 0.0

	if time != "" {
		ts = 1
	}

	return geo.Point{
		p.Longitude,
		p.Latitude,
		p.Elevation * geo.FeetPerMeter,
		ts,
		speed,
	}
}

// ToLine converts all track points into a line of points.
func (f *File) ToLine(name string) []geo.Point {
	var points []geo.Point
	var lastPoint geo.Point

	for _, t := range f.Tracks {
		for _, s := range t.Segments {
			for _, p := range s.Points {
				lastPoint = p.ToArray(lastPoint)
				points = append(points, lastPoint)
			}
		}
	}

	return points
}

func (f *File) ToFeatures() *geoJSON.FeatureCollection {
	var features []*geoJSON.Feature

	for _, t := range f.Tracks {
		for _, s := range t.Segments {
			features = append(features, s.ToFeature())
		}
	}

	return &geoJSON.FeatureCollection{
		Features: features,
	}
}

func (r *Route) ToFeature() *geoJSON.Feature {
	return &geoJSON.Feature{
		Type: "Feature",
		Geometry: &geoJSON.Geometry{
			Type: "LINE",
		},
	}
}

func (s *Segment) ToFeature() *geoJSON.Feature {
	totalTime := 0.0
	totalDistance := 0.0

	geometry := &geoJSON.Geometry{
		Type: "LINE",
	}

	return &geoJSON.Feature{
		Type: "Feature",
		Properties: map[string]string{
			"topSpeed": "",
			"avgSpeed": "",
			"duration": fmt.Sprintf("%f", totalTime),
			"distance": fmt.Sprintf("%f", totalDistance),
		},
		Geometry: geometry,
	}
}
