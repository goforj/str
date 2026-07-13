// Command before is generated as a standalone program so the documented Before example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Before returns the substring before the first occurrence of sep.
	// If sep is empty or not found, the original string is returned.
	// Similar: BeforeLast and After.

	// Example: slice before marker
	v := str.Of("gopher::go").Before("::").String()
	println(v)
	// #string gopher
}
