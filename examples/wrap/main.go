// Command wrap is generated as a standalone program so the documented Wrap example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Wrap surrounds the string with before and after.
	// Similar: Unwrap.

	// Example: wrap string
	v := str.Of("GoForj").Wrap(`"`, `"`).String()
	println(v)
	// #string "GoForj"
}
