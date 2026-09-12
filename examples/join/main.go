// Command join is generated as a standalone program so the documented Join example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Join concatenates elements with sep and returns the result to the fluent chain.
	// Join starts a chain from a slice without discarding an existing receiver.
	// Similar: Split.

	// Example: join words
	v := str.Join([]string{"foo", "bar"}, "-").String()
	println(v)
	// #string foo-bar
}
