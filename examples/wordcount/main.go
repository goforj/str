//go:build ignore
// +build ignore

// Command wordcount is generated as a standalone program so the documented WordCount example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// WordCount returns the number of detected words.
	// Similar: SplitWords.

	// Example: count words
	v := str.Of("Hello, world!").WordCount()
	println(v)
	// #int 2
}
