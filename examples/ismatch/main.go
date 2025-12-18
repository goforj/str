//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
	"regexp"
)

func main() {
	// IsMatch reports whether the string matches the provided regular expression.

	// Example: regex match
	v := str.Of("abc123").IsMatch(regexp.MustCompile(`\d+`))
	godump.Dump(v)
	// #bool true
}
