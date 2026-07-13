// Command after is generated as a standalone program so the documented After example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// After returns the substring after the first occurrence of sep.
	// If sep is empty or not found, the original string is returned.
	// Similar: AfterLast and Before.

	// Example: slice after marker
	v := str.Of("gopher::go").After("::").String()
	println(v)
	// #string go
}
