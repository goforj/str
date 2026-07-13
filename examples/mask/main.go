// Command mask is generated as a standalone program so the documented Mask example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Mask replaces the middle of the string with the given rune, revealing revealLeft runes
	// at the start and revealRight runes at the end. Negative reveal values count from the end.
	// If the reveal counts cover the whole string, the original string is returned.

	// Example: mask email
	v := str.Of("gopher@example.com").Mask('*', 3, 4).String()
	println(v)
	// #string gop***********.com
}
