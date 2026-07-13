// Command ensureprefix is generated as a standalone program so the documented EnsurePrefix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// EnsurePrefix ensures the string starts with prefix, adding it if missing.
	// Similar: EnsureSuffix and TrimPrefix.

	// Example: ensure prefix
	v := str.Of("path/to").EnsurePrefix("/").String()
	println(v)
	// #string /path/to
}
