// Command map is generated as a standalone program so the documented Map example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Map maps each rune using mapping, dropping runes mapped to a negative value.

	// Example: Map
	v := str.Of("go").Map(unicode.ToUpper).String()
	println(v)
	// #string GO
}
