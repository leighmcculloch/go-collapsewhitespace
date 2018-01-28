// +build !go1.10

package collapsewhitespace

import (
	"bytes"
)

// String collapses any whitespace that is not a single standalone space, into
// a single space.
//
// Whitespace is:
//
//   '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func String(s string) string {
	b := bytes.Buffer{}
	b.Grow(len(s))
	lastWasWhitespace := false
	for _, r := range s {
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r', '\u0085', '\u00A0':
			if lastWasWhitespace {
				continue
			} else {
				r = ' '
				lastWasWhitespace = true
			}
		default:
			lastWasWhitespace = false
		}
		b.WriteRune(r)
	}
	return b.String()
}
