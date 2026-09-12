// Command totitlespecial is generated as a standalone program so the documented ToTitleSpecial example can be run directly.
package main

import (
	"unicode"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// ToTitleSpecial maps every rune to Unicode titlecase using c's language-specific case rules.

	// Example: ToTitleSpecial
	v := str.Of("i").ToTitleSpecial(unicode.TurkishCase).String()
	println(v)
	// #string İ
}
