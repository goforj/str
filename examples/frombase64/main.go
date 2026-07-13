//go:build ignore
// +build ignore

// Command frombase64 is generated as a standalone program so the documented FromBase64 example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// FromBase64 decodes a standard Base64 string.
	// Similar: ToBase64.

	// Example: base64 decode
	v, err := str.Of("Z29waGVy").FromBase64()
	println(v.String(), err == nil)
	// #string gopher
	// #bool true
}
