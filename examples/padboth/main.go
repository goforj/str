//go:build ignore
// +build ignore

// Command padboth is generated as a standalone program so the documented PadBoth example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// PadBoth pads the string on both sides to reach length runes using pad (defaults to space).
	// Widths at or below the current rune width leave the string unchanged.
	// Similar: PadLeft and PadRight.

	// Example: pad both
	v := str.Of("go").PadBoth(6, "-").String()
	println(v)
	// #string --go--
}
