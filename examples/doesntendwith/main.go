//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntEndWith reports true if the string ends with none of the provided suffixes (case-sensitive).

	// Example: doesn't end with any
	v := str.Of("gopher")
	godump.Dump(v.DoesntEndWith("her", "cat"))
	// #bool false
}
