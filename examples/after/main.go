//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// After returns the substring after the first occurrence of sep.
	// If sep is empty or not found, the original string is returned.

	// Example: slice after marker
	v := str.Of("gopher::go").After("::").String()
	godump.Dump(v)
	// #string go
}
