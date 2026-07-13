// Command lcfirst is generated as a standalone program so the documented LcFirst example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// LcFirst returns the string with the first rune lower-cased.
	// Similar: UcFirst and ToLower.

	// Example: lowercase first rune
	v := str.Of("Gopher").LcFirst().String()
	fmt.Println(v)
	// #string gopher
}
