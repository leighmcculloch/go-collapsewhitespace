package collapsewhitespace_test

import (
	"fmt"

	"4d63.com/collapsewhitespace"
)

func ExampleString() {
	s := collapsewhitespace.String("abc def			\nghi")

	fmt.Print(s)

	// Output: abc def ghi
}
