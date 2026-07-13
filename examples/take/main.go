// Command take is generated as a standalone program so the documented Take example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Take returns the first length runes of the string (clamped).
	// Similar: TakeLast and Limit.

	// Example: take head
	v := str.Of("gophers").Take(3).String()
	println(v)
	// #string gop
}
