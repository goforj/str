// Command fieldsfunc is generated as a standalone program so the documented FieldsFunc example can be run directly.
package main

import (
	"fmt"
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// FieldsFunc splits the string into fields separated by runes satisfying f.
	// Consecutive separators are combined; empty or separator-only input yields no fields.

	// Example: FieldsFunc
	v := str.Of("a b c").FieldsFunc(unicode.IsSpace)
	fmt.Println(v)
	// #[]string [a b c]
}
