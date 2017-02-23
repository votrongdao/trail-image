package gpx_test

import (
	"io/ioutil"
	"testing"

	"trailimage.com/geo"
	"trailimage.com/geo/gpx"

	"encoding/xml"

	"github.com/stretchr/testify/assert"
)

func mockGPX(t *testing.T, name string) gpx.File {
	dat, err := ioutil.ReadFile("test/" + name + ".gpx")
	assert.NoError(t, err)

	f := gpx.File{}

	err = xml.Unmarshal(dat, &f)
	assert.NoError(t, err)

	return f
}

func TestFirstNode(t *testing.T) {
	p1 := geo.Point{100, 50, 20, 0, 0}
	p2 := geo.Point{100, 50, 30, 0, 0}
	p3 := geo.Point{100, 51, 30, 0, 0}

	assert.True(t, geo.SameLocation(p1, p2))
	assert.False(t, geo.SameLocation(p1, p3))
}

func TestNodeContent(t *testing.T) {
	p1 := geo.Point{-122, 48, 0, 100, 0}
	// an hour later
	p2 := geo.Point{-120, 50, 30, 1000 * 60 * 60, 0}

	assert.InDelta(t, 165, geo.Speed(p1, p2), 1)
}

func TestGeoJSON(t *testing.T) {
	points := []geo.Point{
		geo.Point{-122, 48, 0, 0, 0},
		geo.Point{-121, 49, 0, 0, 0},
		geo.Point{-120, 50, 0, 0, 0},
	}
	assert.InDelta(t, 165, geo.Length(points), 1)
}
