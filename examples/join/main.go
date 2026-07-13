// Command join is generated as a standalone program so the documented Join example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Join concatenates elements with sep and returns the result to the fluent chain.
	// The receiver provides fluent access and is not included in elements.
	// Similar: Split.

	// Example: join words
	v := str.Of("").Join([]string{"foo", "bar"}, "-").String()
	println(v)
	// #string foo-bar
}
