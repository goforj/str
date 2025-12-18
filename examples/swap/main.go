//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Swap replaces multiple values using strings.Replacer built from a map.

	// Example: swap map
	v := str.Of("Gophers are great!")
	pairs := map[string]string{"Gophers": "GoForj", "great": "fantastic"}
	godump.Dump(v.Swap(pairs).String())
	// #string GoForj are fantastic!
}
