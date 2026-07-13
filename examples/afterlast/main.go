//go:build ignore
// +build ignore

// Command afterlast is generated as a standalone program so the documented AfterLast example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// AfterLast returns the substring after the last occurrence of sep.
	// If sep is empty or not found, the original string is returned.
	// Similar: After and BeforeLast.

	// Example: slice after last separator
	v := str.Of("pkg/path/file.txt").AfterLast("/").String()
	println(v)
	// #string file.txt
}
