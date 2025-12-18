//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// FromBase64 decodes a standard Base64 string.

	// Example: base64 decode
	v, err := str.Of("Z29waGVy").FromBase64()
	godump.Dump(v.String(), err)
	// #string gopher
	// #error <nil>
}
