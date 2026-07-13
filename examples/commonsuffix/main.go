//go:build ignore
// +build ignore

// Command commonsuffix is generated as a standalone program so the documented CommonSuffix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// CommonSuffix returns the longest shared suffix between the string and all provided others.
	// Comparison is rune-safe. If no others are provided, the original string is returned.
	// Similar: CommonPrefix.

	// Example: longest common suffix
	v := str.Of("main_test.go").CommonSuffix("user_test.go", "api_test.go").String()
	println(v)
	// #string _test.go
}
