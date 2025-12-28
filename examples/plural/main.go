//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Plural returns a best-effort English plural form of the last word.

	// Example: pluralize word
	v := str.Of("city").Plural().String()
	str.Dump(v)
	// #string cities
}
