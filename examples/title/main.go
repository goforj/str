//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.

	// Example: title case words
	val := str.Of("a nice title uses the correct case")
	godump.Dump(val.Title().String())

	// #string A Nice Title Uses The Correct Case
}
