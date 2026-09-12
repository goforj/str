// Command fieldsfuncseq is generated as a standalone program so the documented FieldsFuncSeq example can be run directly.
package main

import (
	"fmt"
	"slices"
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// FieldsFuncSeq returns an iterator over fields separated by runes satisfying f.
	// Each iteration starts again from the beginning of the string.
	// The predicate must return the same result for a given rune; its call order is unspecified.
	// Consecutive separators are combined; empty or separator-only input yields no fields.

	// Example: FieldsFuncSeq
	v := slices.Collect(str.Of("a b c").FieldsFuncSeq(unicode.IsSpace))
	fmt.Println(v)
	// #[]string [a b c]
}
