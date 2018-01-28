package collapsewhitespace

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(s)
	}
}

func BenchmarkComparedToFieldsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strings.Fields(s), " ")
	}
}

func BenchmarkComparedToRegex(b *testing.B) {
	re := regexp.MustCompile("[\t\n\v\f\r \u0085\u00a0]+")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.ReplaceAllString(s, " ")
	}
}

const s = "   s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s                                                                                                                        s  dj      f		  d d d d d s    "
