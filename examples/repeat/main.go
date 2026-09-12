// Command repeat is generated as a standalone program so the documented Repeat example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Repeat repeats the string count times.
	// It panics if count is negative or the result length overflows int.

	// Example: Repeat
	v := str.Of("go").Repeat(3).String()
	println(v)
	// #string gogogo
}
