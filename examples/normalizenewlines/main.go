//go:build ignore
// +build ignore

// Command normalizenewlines is generated as a standalone program so the documented NormalizeNewlines example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.
	// Similar: Lines.

	// Example: normalize newline variants
	v := str.Of("a\r\nb\u2028c").NormalizeNewlines().String()
	println(v)
	// #string a\nb\nc
}
