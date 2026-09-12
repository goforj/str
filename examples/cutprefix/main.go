// Command cutprefix is generated as a standalone program so the documented CutPrefix example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// CutPrefix removes prefix and reports whether it was present.
	// An empty prefix is a match.

	// Example: split a value
	value, found := str.Of("go:forj").CutPrefix("go:")
	fmt.Println(value, found)
	// #string forj true
}
