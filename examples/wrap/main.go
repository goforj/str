//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Wrap surrounds the string with before and after (after defaults to before).

	// Example: wrap string
	val := str.Of("Laravel")
	godump.Dump(val.Wrap("\"", "").String())

	// #string "Laravel"
}
