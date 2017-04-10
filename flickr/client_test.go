package flickr_test

import (
	"testing"

	"google.golang.org/appengine/aetest"

	"github.com/stretchr/testify/assert"

	"trailimage.com/flickr"
)

// See https://cloud.google.com/appengine/docs/standard/go/tools/localunittesting/
func TestGetSetInfo(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	f := flickr.Configure()

	set, err := f.GetSetInfo("setID")

	assert.NoError(t, err)
	assert.NotNil(t, set)
}
