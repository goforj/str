// Command pascal is generated as a standalone program so the documented Pascal example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Pascal converts the string to PascalCase.
	// Similar: Camel.

	// Example: pascal case
	v := str.Of("foo_bar baz").Pascal().String()
	fmt.Println(v)
	// #string FooBarBaz
}
