//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/str"
)

func main() {
	// LcFirst returns the string with the first rune lower-cased.

	// Example: lowercase first rune
	v := str.Of("Gopher").LcFirst().String()
	fmt.Println(v)
	// #string gopher
}
