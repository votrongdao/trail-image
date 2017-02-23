package gps_test

import (
	"io/ioutil"
	"testing"

	"trailimage.com/gps"

	"github.com/stretchr/testify/assert"
)

func mockResponse(t *testing.T, name string) []byte {
	dat, err := ioutil.ReadFile("test/" + name + ".gpx")
	assert.NoError(t, err)

	return dat
}

func TestSameLocation(t *testing.T) {
	p1 := gps.Point{100, 50, 20, 0, 0}
	p2 := gps.Point{100, 50, 30, 0, 0}
	p3 := gps.Point{100, 51, 30, 0, 0}

	assert.True(t, gps.SameLocation(p1, p2))
	assert.False(t, gps.SameLocation(p1, p3))
}

func TestSpeed(t *testing.T) {
	p1 := gps.Point{-122, 48, 0, 100, 0}
	// an hour later
	p2 := gps.Point{-120, 50, 30, 1000 * 60 * 60, 0}

	assert.InDelta(t, 165, gps.Speed(p1, p2), 1)
}

func TestLength(t *testing.T) {
	points := []gps.Point{
		gps.Point{-122, 48, 0, 0, 0},
		gps.Point{-121, 49, 0, 0, 0},
		gps.Point{-120, 50, 0, 0, 0},
	}
	assert.InDelta(t, 165, gps.Length(points), 1)
}
