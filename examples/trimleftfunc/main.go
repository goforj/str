// Command trimleftfunc is generated as a standalone program so the documented TrimLeftFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimLeftFunc removes leading runes satisfying f.

	// Example: TrimLeftFunc
	v := str.Of("12Go34").TrimLeftFunc(unicode.IsDigit).String()
	println(v)
	// #string Go34
}
