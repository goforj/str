// Command fieldsseq is generated as a standalone program so the documented FieldsSeq example can be run directly.
package main

import (
	"fmt"
	"slices"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// FieldsSeq returns an iterator over fields separated by Unicode whitespace.
	// Each iteration starts again from the beginning of the string.
	// Consecutive separators are combined; empty or separator-only input yields no fields.

	// Example: FieldsSeq
	v := slices.Collect(str.Of("a b c").FieldsSeq())
	fmt.Println(v)
	// #[]string [a b c]
}
