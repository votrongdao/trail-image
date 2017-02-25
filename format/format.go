package format

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"trailimage.com/format/re"

	"regexp"
)

type (
	replacement struct {
		re   *regexp.Regexp
		text []byte
	}
	submatchReplacer func(match []byte, submatch [][]byte) []byte
)

var (
	notWord   = regexp.MustCompile(`\W`)
	protocol  = regexp.MustCompile(`https?://`)
	closeLink = []byte("</a>")
)

// with
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

// replaceEverySubmatch iterates over all match groups and passes sub-
// matches to a replacement function.
//
// The first member of a match group is the text the matched the whole
// expression and subsequent members are the parentheses sub-matches.
func replaceEverySubmatch(text []byte, r *regexp.Regexp, fn submatchReplacer) []byte {
	matchGroups := r.FindAllSubmatch(text, -1)
	if matchGroups == nil || len(matchGroups) == 0 || len(matchGroups[0]) < 2 {
		return text
	}

	for _, g := range matchGroups {
		text = bytes.Replace(text, g[0], fn(g[0], g[1:]), -1)
	}
	return text

}

// Typography replaces punctuation with HTML styled variants.
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

	// restore HTML attribute quotes
	raw = replaceEverySubmatch(raw, re.LinkAttributeQuotes, func(match []byte, submatch [][]byte) []byte {
		return append(
			re.QuoteHTML.ReplaceAll(submatch[0], []byte(`"`)),
			append(submatch[1], closeLink...)...)
	})
	return string(raw)
}

func FixMalformedLink(link string) string {
	index := 0
	raw := replaceEverySubmatch([]byte(link), re.TruncatedLink, func(match []byte, submatch [][]byte) []byte {
		index += 1
		return append(submatch[0], closeLink...)
	})

	// if index > 0 {
	// 	raw = replaceEverySubmatch(raw, re.EllipsisLink, func(submatches [][]byte) []byte {
	// 		return submatches[0]
	// 	})
	// } else {
	// 	return re.EllipsisLink.ReplaceAllString(string(raw), `<a href="$1$2$3">$2$3</a>`)
	// }

	// text = text.replace(re.tag.truncatedLink, (match, missedPart, i) => {
	//    index = i;
	//    return missedPart + '</a>';
	// });

	// if (index > 0) {
	//    const protocol = /https?:\/\//;
	//    const oldLink = text.substring(text.lastIndexOf('<a', index), text.indexOf('</a>', index) + 4);
	//    const newLink = oldLink.replace(re.tag.link, (match, url, name) => {
	//       // add protocol if missing
	//       if (!protocol.test(name)) { name = 'http://' + name; }
	//       return format('<a href="{0}">{1}</a>', name, decodeURI(name.remove(protocol)));
	//    });
	//    text = text.replace(oldLink, newLink);
	// } else {
	//    text = text.replace(re.tag.ellipsisLink, '<a href="$1$2$3">$2$3</a>');
	// }
	return string(raw)
}

// Fraction returns x/y as HTML superscript and subscripts.
func Fraction(text string) string {
	return re.Fraction.ReplaceAllString(text, "<sup>$1</sup>&frasl;<sub>$2</sub>")
}

// PhotoTagList converts a list of photo tag slugs into HTML links.
func PhotoTagList(tagList []string) string {
	links := ""
	link := `<a href="/photo-tag/%s" rel="tag">%s</a> `
	sort.Strings(tagList)
	for _, t := range tagList {
		links += fmt.Sprintf(link, tagSlug(t), t)
	}
	return links
}

func tagSlug(tag string) string {
	return notWord.ReplaceAllString(strings.ToLower(tag), "")
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
