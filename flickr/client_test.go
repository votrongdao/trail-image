package flickr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/flickr"
)

func GetSetInfoTest(t *testing.T) {
	f := flickr.Configure()

	set, err := f.GetSetInfo("setID")

	assert.NoError(t, err)
	assert.NotNil(t, set)
}
