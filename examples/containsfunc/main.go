// Command containsfunc is generated as a standalone program so the documented ContainsFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// ContainsFunc reports whether any rune satisfies f.

	// Example: ContainsFunc
	v := str.Of("go2").ContainsFunc(unicode.IsDigit)
	println(v)
	// #bool true
}
