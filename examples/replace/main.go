// Command replace is generated as a standalone program so the documented Replace example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Replace replaces the first n non-overlapping occurrences of old with new.
	// A negative n replaces all matches; zero leaves the string unchanged.
	// An empty old matches at the beginning and after each UTF-8 sequence.

	// Example: Replace
	v := str.Of("go go go").Replace("go", "Go", 2).String()
	println(v)
	// #string Go Go go
}
