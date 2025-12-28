//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Singular returns a best-effort English singular form of the last word.

	// Example: singularize word
	v := str.Of("people").Singular().String()
	str.Dump(v)
	// #string person
}
