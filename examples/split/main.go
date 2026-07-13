// Command split is generated as a standalone program so the documented Split example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Split splits the string by the given separator.

	// Example: split on comma
	v := str.Of("a,b,c").Split(",")
	fmt.Println(v)
	// #[]string [a b c]
}
