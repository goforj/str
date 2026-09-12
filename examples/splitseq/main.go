// Command splitseq is generated as a standalone program so the documented SplitSeq example can be run directly.
package main

import (
	"fmt"
	"slices"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// SplitSeq returns a single-use iterator over substrings separated by sep.
	// An empty sep splits after each UTF-8 sequence.

	// Example: SplitSeq
	v := slices.Collect(str.Of("a,b,c").SplitSeq(","))
	fmt.Println(v)
	// #[]string [a b c]
}
