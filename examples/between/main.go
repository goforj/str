// Command between is generated as a standalone program so the documented Between example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Between returns the substring between the first start marker and the first end marker after it.
	// It returns an empty string when either marker is empty or missing.

	// Example: between markers
	v := str.Of("[first] and [second]").Between("[", "]").String()
	println(v)
	// #string first
}
