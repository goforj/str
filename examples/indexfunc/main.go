// Command indexfunc is generated as a standalone program so the documented IndexFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// IndexFunc returns the byte offset of the first rune satisfying f, or -1 if absent.

	// Example: IndexFunc
	v := str.Of("go2").IndexFunc(unicode.IsDigit)
	println(v)
	// #int 2
}
