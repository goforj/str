//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// CommonSuffix returns the longest shared suffix between the string and all provided others.
	// Comparison is rune-safe. If no others are provided, the original string is returned.

	// Example: longest common suffix
	v := str.Of("main_test.go").CommonSuffix("user_test.go", "api_test.go").String()
	println(v)
	// #string _test.go
}
