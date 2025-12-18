//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ToUpper returns an uppercase copy of the string using Unicode rules.

	// Example: uppercase text
	val := str.Of("GoLang")
	godump.Dump(val.ToUpper().String())

	// #string GOLANG
}
