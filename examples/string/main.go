//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// String returns the underlying raw string value.

	// Example: unwrap to plain string
	v := str.Of("go").String()
	godump.Dump(v)
	// #string go
}
