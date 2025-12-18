//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Join joins the provided words with sep.

	// Example: join words
	val := str.Of("unused")
	godump.Dump(val.Join([]string{"foo", "bar"}, "-").String())

	// #string foo-bar
}
