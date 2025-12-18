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
	val := str.Of("Tacos are great!")
	pairs := map[string]string{"Tacos": "Burritos", "great": "fantastic"}
	godump.Dump(val.Swap(pairs).String())

	// #string Burritos are fantastic!
}
