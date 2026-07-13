//go:build ignore
// +build ignore

// Command toupper is generated as a standalone program so the documented ToUpper example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ToUpper returns an uppercase copy of the string using Unicode rules.
	// Similar: ToLower and UcFirst.

	// Example: uppercase text
	v := str.Of("GoLang").ToUpper().String()
	println(v)
	// #string GOLANG
}
