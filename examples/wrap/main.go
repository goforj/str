//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Wrap surrounds the string with before and after (after defaults to before).

	// Example: wrap string
	v := str.Of("GoForj").Wrap(`"`, "").String()
	str.Dump(v)
	// #string "GoForj"
}
