package re_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/format/re"
)

func TestEmail(t *testing.T) {
	assert.True(t, re.Email.MatchString("sombody@hello.com"))
	assert.False(t, re.Email.MatchString("sombody-hello.com"))
}

func TestFacebookID(t *testing.T) {
	assert.True(t, re.FacebookID.MatchString("296706240428897.53174"))
}

func TestIpAddress(t *testing.T) {
	assert.True(t, re.IpAddress.MatchString("129.12.113.12"))
	assert.False(t, re.IpAddress.MatchString("1b29.12.113.12"))
	assert.Equal(t, "129.12.113.12", re.IpAddress.FindString("sdfd 129.12.113.12 asdfww"))
}

func TestLinkQuote(t *testing.T) {
	assert.Equal(t, "", re.LinkAttributeQuotes(`<a href=\"/page&rdquo;>so you &ldquo;say&rdquo;</a>`))
}

func TestURL(t *testing.T) {
	assert.True(t, re.URL.MatchString("http://www.somewhere.org"))
	assert.True(t, re.URL.MatchString("http://www.somewhere.org/and/path?qus=2"))
}
