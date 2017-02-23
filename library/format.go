package library

import (
	"fmt"
	"strings"
)

func typography(text string) string {
	if text == "" {
		return text
	}
	return ""
}

func sayNumber(n int, capitalize bool) string {
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
		return strings.ToLower(word)
	}
	return word
}
