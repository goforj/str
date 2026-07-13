// Command padright is generated as a standalone program so the documented PadRight example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// PadRight pads the string on the right to reach length runes using pad (defaults to space).
	// Widths at or below the current rune width leave the string unchanged.
	// Similar: PadLeft and PadBoth.

	// Example: pad right
	v := str.Of("go").PadRight(5, ".").String()
	println(v)
	// #string go...
}
