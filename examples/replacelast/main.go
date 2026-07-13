// Command replacelast is generated as a standalone program so the documented ReplaceLast example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceLast replaces the last occurrence of old with repl.
	// Similar: ReplaceFirst and ReplaceAll.

	// Example: replace last
	v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
	println(v)
	// #string gopher go
}
