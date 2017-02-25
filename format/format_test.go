package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/format"
)

func TestFixMalformedLink(t *testing.T) {
	assert.Equal(t,
		`<a href="http://www.motoidaho.com/sites/default/files/IAMC%20Newsletter%20(4-2011%20Issue%202).pdf">www.motoidaho.com/sites/default/files/IAMC Newsletter (4-2011 Issue 2).pdf</a>`,
		format.FixMalformedLink(`<a href="http://www.motoidaho.com/sites/default/files/IAMC%20Newsletter%20" rel="nofollow">www.motoidaho.com/sites/default/files/IAMC%20Newsletter%20</a>(4-2011%20Issue%202).pdf`))
}

func TestFraction(t *testing.T) {
	assert.Equal(t, "<sup>1</sup>&frasl;<sub>2</sub>", format.Fraction("1/2"))
}

func TestPhotoTagList(t *testing.T) {
	assert.Equal(t,
		`<a href="/photo-tag/first" rel="tag">First</a> <a href="/photo-tag/second" rel="tag">Second</a> <a href="/photo-tag/thirdandlast" rel="tag">Third and Last</a> `,
		format.PhotoTagList([]string{
			"Second", "First", "Third and Last",
		}))

}

func TestTypography(t *testing.T) {
	assert.Empty(t, format.Typography(""))

	assert.Equal(t,
		"&ldquo;He said,&rdquo; she said",
		format.Typography(`"He said," she said`))

	assert.Equal(t,
		`<a href="/page">so you &ldquo;say&rdquo;</a>`,
		format.Typography(`<a href="/page">so you "say"</a>`))

	assert.Equal(t,
		`<A HREF="/page">so you &ldquo;say&rdquo;</a>`,
		format.Typography(`<A HREF="/page">so you "say"</A>`))
}

func TestSayNumber(t *testing.T) {
	assert.Equal(t, "two", format.SayNumber(2, false))
	assert.Equal(t, "Eighteen", format.SayNumber(18, true))
}
