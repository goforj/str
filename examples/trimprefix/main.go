//go:build ignore
// +build ignore

// Command trimprefix is generated as a standalone program so the documented TrimPrefix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimPrefix removes prefix when it appears at the start of the string.
	// Similar: TrimSuffix and EnsurePrefix.

	// Example: trim prefix
	v := str.Of("https://goforj.dev").TrimPrefix("https://").String()
	println(v)
	// #string goforj.dev
}
