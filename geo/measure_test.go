package geo_test

import (
	"testing"

	"trailimage.com/geo"

	"github.com/stretchr/testify/assert"
)

func TestSameLocation(t *testing.T) {
	p1 := geo.Point{100, 50, 20, 0, 0}
	p2 := geo.Point{100, 50, 30, 0, 0}
	p3 := geo.Point{100, 51, 30, 0, 0}

	assert.True(t, geo.SameLocation(p1, p2))
	assert.False(t, geo.SameLocation(p1, p3))
}

func TestToRadians(t *testing.T) {
	assert.InDelta(t, 0.83, geo.ToRadians(48), 1)
	assert.InDelta(t, -2.12, geo.ToRadians(-122), 0.01)
}

func TestSpeed(t *testing.T) {
	p1 := geo.Point{-122, 48, 0, 100, 0}
	// an hour later
	p2 := geo.Point{-120, 50, 30, 1000 * 60 * 60, 0}

	assert.InDelta(t, 165, geo.Speed(p1, p2), 1)
}

func TestLength(t *testing.T) {
	points := []geo.Point{
		geo.Point{-122, 48, 0, 0, 0},
		geo.Point{-121, 49, 0, 0, 0},
		geo.Point{-120, 50, 0, 0, 0},
	}
	assert.InDelta(t, 165, geo.Length(points), 1)
}
