// Command trimfunc is generated as a standalone program so the documented TrimFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimFunc removes leading and trailing runes satisfying f.

	// Example: TrimFunc
	v := str.Of("12Go34").TrimFunc(unicode.IsDigit).String()
	println(v)
	// #string Go
}
