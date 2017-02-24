package re

import "regexp"

const super = "[⁰¹²³⁴⁵⁶⁷⁸⁹]"

// https://golang.org/pkg/regexp/syntax/
// https://gobyexample.com/regular-expressions
// Caveats: https://swtch.com/~rsc/regexp/regexp3.html
//  Synatx: https://github.com/google/re2/wiki/Syntax
var (
	Domain = regexp.MustCompile(`(?i)[a-z0-9][a-z0-9\-]*[a-z0-9]\.[a-z\.]{2,6}$`)

	// EllipsisLink finds anchors named for their URL that are truncated by
	// ellipsis. Example:
	//
	//    <a href="http://idahohistory.cdmhost.com/cdm/singleitem/collection/p16281coll21/id/116/rec/2" rel="nofollow">idahohistory.cdmhost.com/cdm/singleitem/collection/p16281...</a>
	//
	//EllipsisLink = regexp.MustCompile(`(?i)<a href=["'](https?:\/\/)?([^\/]+)([^"']+)['"][^>]*>\2[^<]+\.{3}<\/a>`)

	Email      = regexp.MustCompile(`(?i)\b[A-Z0-9._%-]+@[A-Z0-9.-]+\.[A-Z]{2,4}\b`)
	FacebookID = regexp.MustCompile(`\d{15}\.\d{5}`)
	Fraction   = regexp.MustCompile(`(\d+)\/(\d+)`)
	//FootnoteNumber = regexp.MustCompile(`([^\/\s])([⁰¹²³⁴⁵⁶⁷⁸⁹]+)(?!\w)`)
	FootnoteText = regexp.MustCompile(`(?im)(^|[\r\n]+)_{3}[\r\n]*([\s\S]+)$`)
	IpAddress    = regexp.MustCompile(`\b(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\b`)

	// LinkAttributeQuotes captured for replacement if erroneously encoded to a
	// typography style.
	LinkAttributeQuotes = regexp.MustCompile(`(?i)(<a [^>]+>)([^<]+)<\/a>`)

	QuoteRightSingle = regexp.MustCompile(`(\w)'`)
	QuoteLeftSingle  = regexp.MustCompile(`\b'(\w)`)
	QuoteRightDouble = regexp.MustCompile(`([\w,])("|&quot;)`)
	QuoteLeftDouble  = regexp.MustCompile(`("|&quot;)(\w)`)
	QuoteOpen        = regexp.MustCompile(`^\s*["“]`)
	QuoteEnd         = regexp.MustCompile(`["”]\s*` + super + `?\s*$`)
	QuoteAny         = regexp.MustCompile(`["“”]`)
	QuoteCurly       = regexp.MustCompile(`[“”]`)
	QuoteHTML        = regexp.MustCompile(`(&ldquo;|&rdquo;)`)
	QuoteBlock       = regexp.MustCompile(`(\r\n|\r|\n)?(“[^”]{275,}”` + super + `*)\s*(\r\n|\r|\n)?`)

	// TruncatedLink fixes Flickr prematurely closing tags around text with
	// paranetheses by capturing the wrongly excluded part of the URL. Examples:
	//
	//    <a href="http://www.motoidaho.com/sites/default/files/IAMC%20Newsletter%20" rel="nofollow">www.motoidaho.com/sites/default/files/IAMC%20Newsletter%20</a>(4-2011%20Issue%202).pdf
	//    <a href="http://www.idahogeology.org/PDF/Technical_Reports_" rel="nofollow">www.idahogeology.org/PDF/Technical_Reports_</a>(T)/TR-81-1.pdf
	//
	TruncatedLink = regexp.MustCompile(`(?i)<\/a>(\([\w\/\.\-%\)\(]+)`)
	URL           = regexp.MustCompile(`(http:\/\/[^\s\r\n]+)`)
)
