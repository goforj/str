// Command tobase64 is generated as a standalone program so the documented ToBase64 example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ToBase64 encodes the string using standard Base64.
	// Similar: FromBase64.

	// Example: base64 encode
	v := str.Of("gopher").ToBase64().String()
	println(v)
	// #string Z29waGVy
}
