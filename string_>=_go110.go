// +build go1.10

package collapsewhitespace

import (
	"strings"
	"unicode/utf8"
)

// String collapses any whitespace that is not a single standalone space, into
// a single space.
//
// Whitespace is:
//
//   '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func String(s string) string {
	b := strings.Builder{}
	b.Grow(len(s))
	lastWasWhitespace := false
	for len(s) > 0 {
		r, rLen := utf8.DecodeRuneInString(s)

		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r', '\u0085', '\u00A0':
			if !lastWasWhitespace {
				b.WriteRune(' ')
				lastWasWhitespace = true
			}
		default:
			b.WriteString(s[:rLen])
			lastWasWhitespace = false
		}

		s = s[rLen:]
	}
	return b.String()
}
