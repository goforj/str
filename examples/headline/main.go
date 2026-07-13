// Command headline is generated as a standalone program so the documented Headline example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Headline converts the string into a human-friendly headline:
	// splits on case/underscores/dashes/whitespace, title-cases words, and lowercases small words (except the first).
	// Similar: Title.

	// Example: headline
	v := str.Of("emailNotification_sent").Headline().String()
	println(v)
	// #string Email Notification Sent
}
