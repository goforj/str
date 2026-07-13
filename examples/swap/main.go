// Command swap is generated as a standalone program so the documented Swap example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Swap replaces multiple values using strings.Replacer built from a map.
	// Similar: ReplaceArray.

	// Example: swap map
	pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
	v := str.Of("Gophers are great!").Swap(pairs).String()
	println(v)
	// #string GoForj is fantastic!
}
