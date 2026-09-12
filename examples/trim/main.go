// Command trim is generated as a standalone program so the documented Trim example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Trim removes leading and trailing runes contained in cutset.

	// Example: Trim
	v := str.Of("..GoForj!!").Trim(".!").String()
	println(v)
	// #string GoForj
}
