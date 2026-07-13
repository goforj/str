// Command tolower is generated as a standalone program so the documented ToLower example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ToLower returns a lowercase copy of the string using Unicode rules.
	// Similar: ToUpper and LcFirst.

	// Example: lowercase text
	v := str.Of("GoLang").ToLower().String()
	println(v)
	// #string golang
}
