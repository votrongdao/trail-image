package flickr_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/flickr"
)

func mockResponse(t *testing.T, name string) *flickr.Response {
	dat, err := ioutil.ReadFile("mocks/flickr." + name + ".json")
	assert.NoError(t, err)

	res := &flickr.Response{}
	err = json.Unmarshal(dat, res)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Okay())

	return res
}

func TestCollection(t *testing.T) {
	res := mockResponse(t, "collections.getTree")
	collections := res.Collections.List
	sets := collections[0].Collections[0].Sets

	assert.Equal(t, "When", collections[0].Title)
	assert.Len(t, sets, 5)
	assert.Equal(t, "Stanley Lake Snow Hike", sets[0].Title)
}

func TestEXIF(t *testing.T) {
	res := mockResponse(t, "photos.getExif")

	assert.Len(t, res.Photo.EXIF, 110)
	assert.Equal(t, res.Photo.EXIF[0].Label, "Image Description")
	assert.Equal(t, "NIKON CORPORATION", res.Photo.EXIF[1].Raw.Text)
}

func TestPhotoSearch(t *testing.T) {
	res := mockResponse(t, "photos.search")
	photos := res.PhotoMatch.Photos

	assert.Len(t, photos, 19)
	assert.Equal(t, "690", photos[0].Server)
	assert.Equal(t, "Sentinel", photos[1].Title)
}

func TestPhotoSizes(t *testing.T) {
	res := mockResponse(t, "photos.getSizes")
	sizes := res.Sizes.Size

	assert.Len(t, sizes, 12)
	assert.Equal(t, uint(75), sizes[0].Height)
	assert.Equal(t, uint(150), sizes[1].Width)
}

func TestUserTags(t *testing.T) {
	res := mockResponse(t, "tags.getListUserRaw")
	tags := res.TagMatch.Matches()

	assert.Len(t, tags, 1198)
	assert.Equal(t, "abbott", tags[2].Slug)
	assert.Equal(t, "Aerial", tags[4].Name())
}
