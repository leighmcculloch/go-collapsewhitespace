// +build !go1.10

package collapsewhitespace

import (
	"bytes"
	"unicode/utf8"
)

// String collapses any whitespace that is not a single standalone space, into
// a single space.
//
// Whitespace is:
//
//   '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func String(s string) string {
	buf := bytes.Buffer{}
	buf.Grow(len(s))
	lastWasWhitespace := false
	for len(s) > 0 {
		var size int
		b := s[0]
		if b < utf8.RuneSelf {
			size = 1
			switch b {
			case ' ', '\t', '\n', '\v', '\f', '\r', '\u0085', '\u00A0':
				if !lastWasWhitespace {
					buf.WriteByte(' ')
					lastWasWhitespace = true
				}
			default:
				buf.WriteByte(b)
				lastWasWhitespace = false
			}
		} else {
			var r rune
			r, size = utf8.DecodeRuneInString(s)
			switch r {
			case ' ', '\t', '\n', '\v', '\f', '\r', '\u0085', '\u00A0':
				if !lastWasWhitespace {
					buf.WriteByte(' ')
					lastWasWhitespace = true
				}
			default:
				buf.WriteString(s[:size])
				lastWasWhitespace = false
			}
		}

		s = s[size:]
	}
	return buf.String()
}
