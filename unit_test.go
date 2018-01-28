package collapsewhitespace

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput string
	}{
		{input: "", wantOutput: ""},
		{input: " \t\n\v\f\r\u0085\u00A0", wantOutput: " "},
		{input: " \t1\n2\v3\f4\r5\u00856\u00A07 ", wantOutput: " 1 2 3 4 5 6 7 "},
		{input: "a", wantOutput: "a"},
		{input: "a b", wantOutput: "a b"},
		{input: " a b ", wantOutput: " a b "},
		{input: "  a b ", wantOutput: " a b "},
		{input: "a  b	c", wantOutput: "a b c"},
		{input: "abc	def\nghi", wantOutput: "abc def ghi"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%q=>%q", c.input, c.wantOutput), func(t *testing.T) {
			output := String(c.input)
			if output == c.wantOutput {
				t.Logf("String: got %q", output)
			} else {
				t.Errorf("String: got %q, want %q", output, c.wantOutput)
			}
		})
	}
}
