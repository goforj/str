//go:build ignore
// +build ignore

// Command singular is generated as a standalone program so the documented Singular example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Singular returns a best-effort English singular form of the final identifier word.
	// It handles common English forms and identifier boundaries without claiming to
	// resolve every irregular or ambiguous noun.
	// Similar: Plural.

	// Example: singularize word
	v := str.Of("people").Singular().String()
	println(v)
	// #string person
}
