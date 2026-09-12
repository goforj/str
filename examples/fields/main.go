// Command fields is generated as a standalone program so the documented Fields example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Fields splits the string into fields separated by Unicode whitespace.
	// Consecutive separators are combined; empty or separator-only input yields no fields.

	// Example: Fields
	v := str.Of("a b c").Fields()
	fmt.Println(v)
	// #[]string [a b c]
}
