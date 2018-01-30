package collapsewhitespace

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkASCIIString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(asciiStr)
	}
}

func BenchmarkASCIIComparedToFieldsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(strings.Fields(asciiStr), " ")
	}
}

func BenchmarkASCIIComparedToRegex(b *testing.B) {
	re := regexp.MustCompile(`\s+`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.ReplaceAllString(asciiStr, " ")
	}
}

func BenchmarkASCIIComparedToRegexPerfect(b *testing.B) {
	re := regexp.MustCompile("[\t\n\v\f\r \u0085\u00a0]+")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.ReplaceAllString(asciiStr, " ")
	}
}

func BenchmarkUnicodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(unicodeStr)
	}
}

func BenchmarkUnicodeComparedToRegex(b *testing.B) {
	re := regexp.MustCompile(`\s+`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.ReplaceAllString(unicodeStr, " ")
	}
}

func BenchmarkUnicodeComparedToRegexPerfect(b *testing.B) {
	re := regexp.MustCompile("[\t\n\v\f\r \u0085\u00a0]+")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.ReplaceAllString(unicodeStr, " ")
	}
}

const asciiStr = "   s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s      s  dj      f		  d d d d d s                                                                                                                        s  dj      f		  d d d d d s    "

const unicodeStr = "%Ǉ\x9cd;\x8b\x9b5\x9e\"\x03S\xb0\xe2\xf4Qb\x03\x94[-\xa4F\xe0\xb1^\xf2\x83L\xfc*|\x84&8\xcb/\xe9Òg\xb6\xfe\xb0\xf4\x01bFZҫ\xa3U\x9a1\x98%\xbb\x85I\xf8\x16\x10M\xc1\x97\x9e\xd5Wi\x17ǓR\xf0\r\xf3X\xb03\xae\x16\x86\xc0\x12\xfbѵr1oL\xbc\b1&\xdd&\xa7\x0e\xd6X\xeeC\xcf+\x1cT\xb8\x01\x15T\"z\x81]\xc5*\xec\xfa{\xf2\xecP1l\x91n\x13\xf3\xc5\xcd\x02\xd8\xe2Ћ\x96\xcd\xc3U[RH\xb6\xb7\b\xc8;З\x16<c[\xd9\x05\a\xf6\xe3qT\r\xac\x8d\x18\a\xbfn\xe6E-\x15\xc9\x17l\xe4\x97cH\x86H\xa0\xa3\x94\x9b$\x02Vz\x12\nU\x8f\xe4\x13t\xce\xe3}\xd7\xff\xc9\xdfӶ\xcd\xeeն\xa4\xe8e\xd8E\x94\xf6\xd6!8lD/\x10\xb0\xcfa3\xc4\xc02\x83t\xbe)D.{ָ\x01O!\x1d\x18\x14\x99\xb23pzF}(0\xf7\xea\xfd+\"\x92h\xdc\x13\x1e-\xa0\x05VӇ\x93\x976*S\f\x85nr\xa2x\x88\xf7\xa8\xc8\x0fq\xe1\n(t\xf3S\f\x15\x83Or\xf6\xac\x1d\xa3\xee\x8a=_\xf8\x8d}\x83\x9d\x9e\x03k\xcce`\x9e@\x11_\x1b\xfd\xbb*=\x19\x90Y\xce\xe9h\xfd\x1e\xf76\xa9\xe1.\x93A\x1a2\x93\xac,l\xa9(\x8c\x95\\\xa7:\x8c\br\xe0&&\x91\xf5^\x96\x96|}\r\x1cm\vy\xa2\xfbIe\xb3\xd9|\x94\x8bdnʣ\xb6\xef\xfaG\x06\xe6\x99\x18\xa2\x1es\xb7\xba\xf8\x00\xaf\xaf\xa7\xc9X\x84\xb1\xa7_\xa4\u061c,\x84\x8c\xa8\x82-OK\x93\xdd/\v\xfbq;\xce(\x96\xef5A\xb5G:\xd2\xee&?j\x87o\xbd\x15\xb1\xa0R\xc5\xe9\"\xc0\xc2\x17\xe0\x8a\x8b\xd0\xdb\xf6\xe93\xadA\xf0zxZ(S\x98\xa1\xd6Ω=\x01\xd2I\x86\x19Y\"\x8e\xe8MIO7\xba\"5\x87\x031\x16\xd0\xfcub\x19\xbb\x00`\xef{:%\xe1\x17_\xb0[y\xdeRz\x0e*\x14\xec\x84*\x8f\x88\xe16\xbdp\xa8ެ\xf3\x96\x81K\x90\x1f\xfe\xcc\xdb\xd2\xff#x\xeb6\x8f\xa4z\xb1=\u007f\x95\xdb\b\x92\xb3\x86\x9f\xa2ʊP\x17\xb3#\xd0l\aֺ(s\xa1$\xa2\a\xadS\x1b1i--dZ\x109\xaeʤ)\xbf\x13+\xc8}\xb18Q\xe6\xfb\xfa$\x81jۻ\x1d\x9c\x89\x0f\x92Z\x899\x1d\xfeIN\x1b\x96\xa1\x13Z+\\ہ\xd3a\xe9\x1d\xa6\xb4٨]a\xec\n\x11\x82\xb9:\x91\xe2eS͟:\xbbZ\xec\a>\xcew\x18\xb6\x18d\xfe4\tx<$i%=J\x19\xe5w.\xeb\xe8\xf7\xf5jADq!\x86\xf0\x88\xa5\x05\xaf\xb1\x93/\x93azT[\xf6!\xa4\x1e&~\xa5S&nץ\x9fp\x90\xe4\xf9\x10\x92\xb9\x85\xc1\xf3\x9f~'\x1d\xee\xc1\xde&\x01\x8fc6<\x9aߺ\xa9\x04S\xf896+_ǹ\xf6\xc6=\xd3.l\x8fO\xd8\x1f\xa0\xe7\xda<^\xbc\x04fN:\xa4\x06\xf3\x9fqv\xd2\xd5\xfbJ/CMHB\xf2\x99\xea*\xfc\x93\x19\x11$R7G%\xeah\x9c\xa5͜\xea\xec\x92U@\x1d$4DYFߑ(8\xbaL\xdf\xed;\xaa\x8b\x9f\x17\xcb(ц\xf3\x06v\xeap\xd1\xfb\x1dƖ7\x9cȆz\xb9\xaa\x84\x1dT\x0fG\xc4\v\xd2NuiZ\xb4\x1b\xdd@\xea\x1aC\xf4k\x86\xc1\x1dx\xc5\xfb\x90\x813\xed\xfc\x12\x8d\a5^!\x89\xedݱ\xca\xf4\x1d1\x91\x17\x9faԊ\xc7*\n\xdb\xcb\x13\xda!=\x97ҽ\xd2\x1eY\x90\x14e\x12r\xa3\x8e-T\xe3n0\x99\x86O\v=k^\xf2YH\xf5S\x9e\r*\x92i\t\bBG݃\xd3\xde[\xcfU(\xff\xd7˚/\x14ٯ\bb-1h\x03\x9b\xb4n\x86\xd3#\xf6\xc0\xd5\xd9\xea\xb6;r\xfai\xb8+#\xcajS\xbcNNc/Ⱦ\xbf\x97+ ^\xf8\xca\x00\x88\xb8\xb5J-a\xc44\x81\xfdr\xf9\x92B"
