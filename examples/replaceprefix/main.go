// Command replaceprefix is generated as a standalone program so the documented ReplacePrefix example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplacePrefix replaces old with repl when old is a prefix of the string.
	// An empty old inserts repl at the beginning.
	// Similar: ReplaceSuffix and TrimPrefix.

	// Example: replace prefix
	v := str.Of("prefix-value").ReplacePrefix("prefix-", "new-").String()
	println(v)
	// #string new-value
}
