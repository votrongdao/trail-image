package gpx_test

import (
	"io/ioutil"
	"testing"

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
	assert.NotNil(t, f)

	return f
}

func TestTrackFile(t *testing.T) {
	f := mockGPX(t, "track")

	assert.Equal(t, "track.gpx", f.Name)
	assert.Len(t, f.Tracks, 2)
	assert.Len(t, f.Tracks[0].Segments, 1)
	assert.Len(t, f.Tracks[0].Segments[0].Points, 23)

	points := f.Tracks[0].Segments[0].Points

	assert.Equal(t, 43.238334, points[0].Latitude)
	assert.Equal(t, -116.3666, points[0].Longitude)
	assert.Equal(t, 926.90, points[0].Elevation)
}

func TestBigTrackFile(t *testing.T) {
	f := mockGPX(t, "track-big")

	assert.Equal(t, "Owyhee Snow and Sand.gpx", f.Name)
	assert.Len(t, f.Tracks, 4)
	assert.Equal(t, "2014-05-19, 014140P (Segment 6)", f.Tracks[0].Name)
	assert.Len(t, f.Tracks[0].Segments, 1)
	assert.Len(t, f.Tracks[0].Segments[0].Points, 23)

	points := f.Tracks[0].Segments[0].Points

	assert.Equal(t, 43.586089, points[0].Latitude)
	assert.Equal(t, -116.174093, points[0].Longitude)
	assert.Equal(t, 835.17, points[0].Elevation)
}
