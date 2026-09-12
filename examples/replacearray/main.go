// Command replacearray is generated as a standalone program so the documented ReplaceArray example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceArray replaces all occurrences of each old in olds with repl.
	// Entries are applied sequentially, including replacements produced by earlier entries.
	// Empty entries insert repl at UTF-8 boundaries, like ReplaceAll.
	// Similar: ReplaceAll and Swap.

	// Example: replace many
	v := str.Of("The---Go---Toolkit")
	println(v.ReplaceArray([]string{"---"}, "-").String())
	// #string The-Go-Toolkit
}
