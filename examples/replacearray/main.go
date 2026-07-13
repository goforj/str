//go:build ignore
// +build ignore

// Command replacearray is generated as a standalone program so the documented ReplaceArray example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceArray replaces all occurrences of each old in olds with repl.
	// Similar: ReplaceAll and Swap.

	// Example: replace many
	v := str.Of("The---Go---Toolkit")
	println(v.ReplaceArray([]string{"---"}, "-").String())
	// #string The-Go-Toolkit
}
