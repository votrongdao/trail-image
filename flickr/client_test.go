package flickr_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/flickr"
)

func mockResponse(name string) (*flickr.Response, error) {
	dat, err := ioutil.ReadFile("mocks/flickr." + name + ".json")
	if err != nil {
		return nil, err
	}
	res := &flickr.Response{}

	if err = json.Unmarshal(dat, res); err != nil {
		return nil, err
	}
	return res, nil
}

func TestCollection(t *testing.T) {
	res, err := mockResponse("collections.getTree")

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
