// Command fieldsseq is generated as a standalone program so the documented FieldsSeq example can be run directly.
package main

import (
	"fmt"
	"slices"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// FieldsSeq returns a single-use iterator over fields separated by Unicode whitespace.
	// Consecutive separators are combined; empty or separator-only input yields no fields.

	// Example: FieldsSeq
	v := slices.Collect(str.Of("a b c").FieldsSeq())
	fmt.Println(v)
	// #[]string [a b c]
}
