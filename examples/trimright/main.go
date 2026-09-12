// Command trimright is generated as a standalone program so the documented TrimRight example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimRight removes trailing runes contained in cutset.

	// Example: TrimRight
	v := str.Of("..GoForj!!").TrimRight(".!").String()
	println(v)
	// #string ..GoForj
}
