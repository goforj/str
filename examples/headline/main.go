//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Headline converts the string into a human-friendly headline:
	// splits on case/underscores/dashes/whitespace, title-cases words, and lowercases small words (except the first).

	// Example: headline
	v := str.Of("emailNotification_sent").Headline().String()
	godump.Dump(v)
	// #string Email Notification Sent
}
