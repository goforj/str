// Command lastindexfunc is generated as a standalone program so the documented LastIndexFunc example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// LastIndexFunc returns the byte offset of the last rune satisfying f, or -1 if absent.

	// Example: LastIndexFunc
	v := str.Of("go2").LastIndexFunc(unicode.IsDigit)
	println(v)
	// #int 2
}
