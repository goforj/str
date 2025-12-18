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
	v := str.Of("GoLang")
	godump.Dump(v.ToUpper().String())
	// #string GOLANG
}
