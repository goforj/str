// Command totitle is generated as a standalone program so the documented ToTitle example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ToTitle maps every rune to Unicode titlecase.

	// Example: ToTitle
	v := str.Of("go").ToTitle().String()
	println(v)
	// #string GO
}
