// Command isnumeric is generated as a standalone program so the documented IsNumeric example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsNumeric reports whether the string contains at least one rune and every rune is a Unicode number.

	// Example: numeric check
	v := str.Of("12345").IsNumeric()
	println(v)
	// #bool true
}
