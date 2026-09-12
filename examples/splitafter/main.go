// Command splitafter is generated as a standalone program so the documented SplitAfter example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// SplitAfter returns substrings including their trailing separator.
	// An empty sep splits after each UTF-8 sequence.

	// Example: SplitAfter
	v := str.Of("a,b,c").SplitAfter(",")
	fmt.Println(v)
	// #[]string [a, b, c]
}
