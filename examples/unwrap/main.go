// Command unwrap is generated as a standalone program so the documented Unwrap example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Unwrap removes matching before and after strings if present.
	// Similar: Wrap.

	// Example: unwrap string
	v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
	println(v)
	// #string GoForj
}
