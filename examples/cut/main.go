// Command cut is generated as a standalone program so the documented Cut example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Cut splits around the first occurrence of sep.
	// An empty sep is a match. On no match it returns the original string, "", false.

	// Example: split a value
	before, after, found := str.Of("go:forj").Cut(":")
	fmt.Println(before, after, found)
	// #string go forj true
}
