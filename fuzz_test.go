package collapsewhitespace

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"testing/quick"
	"time"
)

const fuzzWhitespace = " \t\n\v\f\r\u0085\u00A0"

type spaceyString string

func (s spaceyString) Generate(r *rand.Rand, size int) reflect.Value {
	bytes := make([]byte, size)

	// between 0 and size whitespace
	whitespaceCount := r.Intn(size + 1)

	// insert random whitespace
	for i := 0; i < whitespaceCount; i++ {
		bytes[i] = fuzzWhitespace[r.Intn(len(fuzzWhitespace))]
	}

	// fill remaining with random bytes
	r.Read(bytes[whitespaceCount:])

	// shuffle the bytes
	perm := rand.Perm(len(bytes))
	for i, j := range perm {
		bytes[j], bytes[i] = bytes[i], bytes[j]
	}

	return reflect.ValueOf(spaceyString(bytes))
}

func checkString(t *testing.T, check func(input, output string) bool) {
	t.Helper()

	f := func(s spaceyString) bool {
		return check(string(s), String(string(s)))
	}

	seed := time.Now().UnixNano()
	t.Logf("Seed: %d", seed)

	c := &quick.Config{
		Rand: rand.New(rand.NewSource(seed)),
	}

	if err := quick.Check(f, c); err != nil {
		if checkErr, ok := err.(*quick.CheckError); ok {
			input := string(checkErr.In[0].(spaceyString))
			t.Error(err)
			t.Errorf("Input: %x", input)
			t.Errorf("Output: %x", String(input))
		} else {
			t.Fatal(err)
		}
	}
}

func TestCompactFuzz_NoRepeatingWhitespace(t *testing.T) {
	re := regexp.MustCompile("[" + fuzzWhitespace + "]{2,}")

	f := func(input, output string) bool {
		hasRepeatingWhitespace := re.MatchString(output)
		return !hasRepeatingWhitespace
	}

	checkString(t, f)
}

func TestCompactFuzz_ContainsOriginalNonWhitespaceCharacters(t *testing.T) {
	re := regexp.MustCompile("[" + fuzzWhitespace + "]+")

	f := func(input, output string) bool {
		words := re.Split(input, -1)
		for _, word := range words {
			if !strings.Contains(output, word) {
				return false
			}
		}
		return true
	}

	checkString(t, f)
}

func TestCompactFuzz_OutputSmallerOrEqualLengthThanInput(t *testing.T) {
	f := func(input, output string) bool {
		return len(output) <= len(input)
	}

	checkString(t, f)
}

func TestCompactFuzz_OutputNotEmpty(t *testing.T) {
	f := func(input, output string) bool {
		return len(output) > 0
	}

	checkString(t, f)
}
