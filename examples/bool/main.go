// Command bool is generated as a standalone program so the documented Bool example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Bool parses the string as a bool using strconv.ParseBool semantics.
	// Similar: Int and Float64.

	// Example: parse bool
	v, err := str.Of("true").Bool()
	println(v, err == nil)
	// #bool true
	// #bool true
}
