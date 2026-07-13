// Command slug is generated as a standalone program so the documented Slug example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Slug returns a lowercase Unicode slug separated by hyphens.
	// Unicode letters and digits are preserved, while every other run is collapsed
	// to one hyphen.
	// Similar: Kebab.

	// Example: build slug
	v := str.Of("Go Forj Toolkit").Slug().String()
	println(v)
	// #string go-forj-toolkit
}
