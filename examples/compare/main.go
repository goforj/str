// Command compare is generated as a standalone program so the documented Compare example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Compare returns -1, 0, or 1 according to lexicographic byte order.

	// Example: Compare
	v := str.Of("go").Compare("rust")
	println(v)
	// #int -1
}
