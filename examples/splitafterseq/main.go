// Command splitafterseq is generated as a standalone program so the documented SplitAfterSeq example can be run directly.
package main

import (
	"fmt"
	"slices"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// SplitAfterSeq returns a single-use iterator over substrings including their trailing separator.
	// An empty sep splits after each UTF-8 sequence.

	// Example: SplitAfterSeq
	v := slices.Collect(str.Of("a,b,c").SplitAfterSeq(","))
	fmt.Println(v)
	// #[]string [a, b, c]
}
