package format

import (
	"fmt"
	"strings"

	"trailimage.com/format/re"

	"regexp"
)

type (
	replacement struct {
		re   *regexp.Regexp
		text []byte
	}
)

func with(re *regexp.Regexp, text string) replacement {
	return replacement{
		re:   re,
		text: []byte(text),
	}
}

func replaceAll(text []byte, replacements ...replacement) []byte {
	for _, r := range replacements {
		text = r.re.ReplaceAll(text, r.text)
	}
	return text
}

// Typography stylizes punctuation.
func Typography(text string) string {
	if text == "" {
		return text
	}
	raw := replaceAll([]byte(text),
		with(re.QuoteRightSingle, "$1&rsquo;"),
		with(re.QuoteLeftSingle, "&lsquo;$1"),
		with(re.QuoteRightDouble, "$1&rdquo;"),
		with(re.QuoteLeftDouble, "&ldquo;$2"),
	)

	return string(raw)
}

// Fraction returns x/y as HTML superscript and subscripts.
func Fraction(text string) string {
	return re.Fraction.ReplaceAllString(text, "<sup>$1</sup>&frasl;<sub>$2</sub>")
}

func SayNumber(n int, capitalize bool) string {
	word := fmt.Sprintf("%d", n)

	switch n {
	case 1:
		word = "One"
	case 2:
		word = "Two"
	case 3:
		word = "Three"
	case 4:
		word = "Four"
	case 5:
		word = "Five"
	case 6:
		word = "Six"
	case 7:
		word = "Seven"
	case 8:
		word = "Eight"
	case 9:
		word = "Nine"
	case 10:
		word = "Ten"
	case 11:
		word = "Eleven"
	case 12:
		word = "Twelve"
	case 13:
		word = "Thirteen"
	case 14:
		word = "Fourteen"
	case 15:
		word = "Fifteen"
	case 16:
		word = "Sixteen"
	case 17:
		word = "Seventeen"
	case 18:
		word = "Eighteen"
	case 19:
		word = "Nineteen"
	case 20:
		word = "Twenty"
	}

	if capitalize {
		return word
	}
	return strings.ToLower(word)
}
