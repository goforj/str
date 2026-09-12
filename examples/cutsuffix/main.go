// Command cutsuffix is generated as a standalone program so the documented CutSuffix example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// CutSuffix removes suffix and reports whether it was present.
	// An empty suffix is a match.

	// Example: split a value
	value, found := str.Of("go:forj").CutSuffix(":forj")
	fmt.Println(value, found)
	// #string go true
}
