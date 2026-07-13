// Command splitwords is generated as a standalone program so the documented SplitWords example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// SplitWords splits the string into Unicode words, including camel case and acronym boundaries.
	// Similar: FirstWord, LastWord, WordCount, and Words.

	// Example: split words
	v := str.Of("one, two, three").SplitWords()
	fmt.Println(v)
	// #[]string [one two three]
}
