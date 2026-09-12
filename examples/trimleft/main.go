// Command trimleft is generated as a standalone program so the documented TrimLeft example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimLeft removes leading runes contained in cutset.

	// Example: TrimLeft
	v := str.Of("..GoForj!!").TrimLeft(".!").String()
	println(v)
	// #string GoForj!!
}
