//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// AfterLastFold returns the substring after the last occurrence of sep using Unicode-aware
	// case-insensitive comparison. If sep is empty or not found, the original string is returned.

	// Example: slice after last separator (case-insensitive)
	v := str.Of("pkg/Path/FILE.txt").AfterLastFold("/path/").String()
	println(v)
	// #string FILE.txt
}
