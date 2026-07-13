// Command replacesuffix is generated as a standalone program so the documented ReplaceSuffix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceSuffix replaces old with repl when old is a suffix of the string.
	// Similar: ReplacePrefix and TrimSuffix.

	// Example: replace suffix
	v := str.Of("file.old").ReplaceSuffix(".old", ".new").String()
	println(v)
	// #string file.new
}
