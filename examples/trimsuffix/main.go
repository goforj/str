//go:build ignore
// +build ignore

// Command trimsuffix is generated as a standalone program so the documented TrimSuffix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimSuffix removes suffix when it appears at the end of the string.
	// Similar: TrimPrefix and EnsureSuffix.

	// Example: trim suffix
	v := str.Of("file.txt").TrimSuffix(".txt").String()
	println(v)
	// #string file
}
