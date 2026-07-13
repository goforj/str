//go:build ignore
// +build ignore

// Command padleft is generated as a standalone program so the documented PadLeft example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// PadLeft pads the string on the left to reach length runes using pad (defaults to space).
	// Widths at or below the current rune width leave the string unchanged.
	// Similar: PadRight and PadBoth.

	// Example: pad left
	v := str.Of("go").PadLeft(5, " ").String()
	println(v)
	// #string \u0020\u0020\u0020go
}
