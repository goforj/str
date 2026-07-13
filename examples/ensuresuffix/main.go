//go:build ignore
// +build ignore

// Command ensuresuffix is generated as a standalone program so the documented EnsureSuffix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// EnsureSuffix ensures the string ends with suffix, adding it if missing.
	// Similar: EnsurePrefix and TrimSuffix.

	// Example: ensure suffix
	v := str.Of("path/to").EnsureSuffix("/").String()
	println(v)
	// #string path/to/
}
