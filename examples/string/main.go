// Command string is generated as a standalone program so the documented String example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// String returns the underlying raw string value.

	// Example: unwrap to plain string
	v := str.Of("go").String()
	println(v)
	// #string go
}
