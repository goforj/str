// Command toupperspecial is generated as a standalone program so the documented ToUpperSpecial example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// ToUpperSpecial maps every rune to uppercase using c's language-specific case rules.

	// Example: ToUpperSpecial
	v := str.Of("i").ToUpperSpecial(unicode.TurkishCase).String()
	println(v)
	// #string İ
}
