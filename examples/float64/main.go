//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Float64 parses the string as a float64 using strconv.ParseFloat semantics.

	// Example: parse float64
	v, err := str.Of("3.14").Float64()
	println(v, err == nil)
	// #float64 3.14
	// #bool true
}
