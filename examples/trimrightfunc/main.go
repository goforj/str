// Command trimrightfunc is generated as a standalone program so the documented TrimRightFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimRightFunc removes trailing runes satisfying f.

	// Example: TrimRightFunc
	v := str.Of("12Go34").TrimRightFunc(unicode.IsDigit).String()
	println(v)
	// #string 12Go
}
