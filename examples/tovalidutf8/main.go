// Command tovalidutf8 is generated as a standalone program so the documented ToValidUTF8 example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ToValidUTF8 replaces each run of invalid UTF-8 bytes with replacement.

	// Example: ToValidUTF8
	v := str.Of("a\xff\xfeb").ToValidUTF8("?").String()
	println(v)
	// #string a?b
}
