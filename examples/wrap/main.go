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
	v := str.Of("Laravel")
	godump.Dump(v.Wrap("\"", "").String())
	// #string "Laravel"
}
