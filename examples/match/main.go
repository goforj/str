// Command match is generated as a standalone program so the documented Match example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Match reports whether the entire string matches pattern using [path.Match] syntax.
	// A malformed pattern returns an error, and wildcards do not match a slash.

	// Example: match a shell pattern
	matched, err := str.Of("billing:reports").Match("billing:*")
	println(matched, err == nil)
	// #bool true
	// #bool true
}
