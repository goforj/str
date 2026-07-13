// Command title is generated as a standalone program so the documented Title example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.
	// Similar: Headline.

	// Example: title case words
	v := str.Of("a nice title uses the correct case").Title().String()
	println(v)
	// #string A Nice Title Uses The Correct Case
}
