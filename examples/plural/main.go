//go:build ignore
// +build ignore

// Command plural is generated as a standalone program so the documented Plural example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Plural returns a best-effort English plural form of the final identifier word.
	// It handles common English forms and identifier boundaries without claiming to
	// resolve every irregular or ambiguous noun.
	// Similar: Singular.

	// Example: pluralize word
	v := str.Of("city").Plural().String()
	println(v)
	// #string cities
}
