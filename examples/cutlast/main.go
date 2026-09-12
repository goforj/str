// Command cutlast is generated as a standalone program so the documented CutLast example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// CutLast splits around the last occurrence of sep.
	// An empty sep is a match. On no match it returns the original string, "", false.
	// CutLast preserves the Go 1.27 contract without raising the Go 1.24 minimum.

	// Example: split a value
	before, after, found := str.Of("go:forj").CutLast(":")
	fmt.Println(before, after, found)
	// #string go forj true
}
