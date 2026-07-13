//go:build ignore
// +build ignore

// Command wrapwords is generated as a standalone program so the documented WrapWords example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// WrapWords wraps the string to the given rune width on whitespace boundaries,
	// using breakStr between lines without discarding punctuation.
	// Similar: Words.

	// Example: wrap words
	v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
	println(v)
	// #string The quick brown fox\njumped over the lazy\ndog.
}
