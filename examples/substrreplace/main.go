// Command substrreplace is generated as a standalone program so the documented SubstrReplace example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// SubstrReplace replaces the rune slice in [start:end) with repl.

	// Example: replace range
	v := str.Of("naïve café").SubstrReplace("i", 2, 3).String()
	println(v)
	// #string naive café
}
