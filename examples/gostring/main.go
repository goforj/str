// Command gostring is generated as a standalone program so the documented GoString example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// GoString allows %#v formatting to print the raw string.

	// Example: fmt %#v uses GoString
	v := str.Of("go")
	println(fmt.Sprintf("%#v", v))
	// #string go
}
