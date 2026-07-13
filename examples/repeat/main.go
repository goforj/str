// Command repeat is generated as a standalone program so the documented Repeat example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Repeat repeats the string count times (non-negative).

	// Example: repeat string
	v := str.Of("go").Repeat(3).String()
	println(v)
	// #string gogogo
}
