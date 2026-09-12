// Command splitn is generated as a standalone program so the documented SplitN example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// SplitN returns substrings separated by sep.
	// A positive n limits the result to n substrings; zero returns nil and negative n has no limit.
	// An empty sep splits after each UTF-8 sequence.

	// Example: SplitN
	v := str.Of("a,b,c").SplitN(",", 2)
	fmt.Println(v)
	// #[]string [a b,c]
}
