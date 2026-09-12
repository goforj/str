// Command count is generated as a standalone program so the documented Count example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Count counts non-overlapping occurrences of sub.
	// An empty sub matches before and after each UTF-8 sequence.

	// Example: Count
	v := str.Of("gogophergo").Count("go")
	println(v)
	// #int 3
}
