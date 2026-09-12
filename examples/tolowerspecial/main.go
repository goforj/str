// Command tolowerspecial is generated as a standalone program so the documented ToLowerSpecial example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// ToLowerSpecial maps every rune to lowercase using c's language-specific case rules.

	// Example: ToLowerSpecial
	v := str.Of("I").ToLowerSpecial(unicode.TurkishCase).String()
	println(v)
	// #string ı
}
