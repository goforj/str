// Command contains is generated as a standalone program so the documented Contains example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Contains reports whether the string contains sub using a case-sensitive comparison.
	// An empty substring is not a match.
	// Similar: ContainsFold.

	// Example: contains substring
	v := str.Of("Go means gophers").Contains("gopher")
	println(v)
	// #bool true
}
