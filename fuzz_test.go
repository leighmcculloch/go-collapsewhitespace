package collapsewhitespace

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func fuzzInput() string {
	b := make([]byte, 1024)
	n, err := rand.Read(b)
	if err != nil {
		panic(fmt.Errorf("fuzzInput: rand.Read: %v", err))
	}
	if n != len(b) {
		panic(fmt.Errorf("fuzzInput: rand.Read: got %d bytes, want %d bytes", n, len(b)))
	}
	return string(b)
}

func fuzzInputs() []string {
	inputs := [1024]string{}
	for i := 0; i < len(inputs); i++ {
		inputs[i] = fuzzInput()
	}
	return inputs[:]
}

func TestCompactFuzz_NoRepeatingWhitespace(t *testing.T) {
	re := regexp.MustCompile(`\s{2,}`)
	for _, input := range fuzzInputs() {
		t.Run(input, func(t *testing.T) {
			output := String(input)
			if re.MatchString(output) {
				t.Errorf("%q => %q has sequential whitespace", input, output)
				return
			}
		})
	}
}

func TestCompactFuzz_ContainsOriginalNonWhitespaceCharacters(t *testing.T) {
	re := regexp.MustCompile("[\t\n\v\f\r \u0085\u00a0]+")
	for _, input := range fuzzInputs() {
		t.Run(input, func(t *testing.T) {
			words := re.Split(input, -1)
			output := String(input)
			for _, word := range words {
				if !strings.Contains(output, word) {
					t.Errorf("%#x => %#x is missing word(s) that should have remained: %#x", []byte(input), []byte(output), []byte(word))
					return
				}
			}
		})
	}
}
