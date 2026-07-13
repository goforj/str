// Command trimchars is generated as a standalone program so the documented TrimChars example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimChars removes leading and trailing runes contained in cutset.
	// Similar: Trim.

	// Example: trim selected characters
	v := str.Of("..GoForj!!").TrimChars(".!").String()
	println(v)
	// #string GoForj
}
