// Command float64 is generated as a standalone program so the documented Float64 example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Float64 parses the string as a float64 using strconv.ParseFloat semantics.
	// Similar: Bool and Int.

	// Example: parse float64
	v, err := str.Of("3.14").Float64()
	fmt.Println(v, err == nil)
	// #float64 3.14
	// #bool true
}
