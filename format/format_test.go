package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trailimage.com/format"
)

func TestFraction(t *testing.T) {
	assert.Equal(t, "<sup>1</sup>&frasl;<sub>2</sub>", format.Fraction("1/2"))
}

func TestTypography(t *testing.T) {
	assert.Empty(t, format.Typography(""))

	assert.Equal(t,
		"&ldquo;He said,&rdquo; she said",
		format.Typography(`"He said," she said`))

	assert.Equal(t,
		`<a href="/page">so you &ldquo;say&rdquo;</a>`,
		format.Typography(`<a href="/page">so you "say"</a>`))
}

func TestSayNumber(t *testing.T) {
	assert.Equal(t, "two", format.SayNumber(2, false))
	assert.Equal(t, "Eighteen", format.SayNumber(18, true))
}
