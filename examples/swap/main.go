// Command swap is generated as a standalone program so the documented Swap example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Swap replaces multiple values in one pass using strings.Replacer built from a map.
	// Longer keys take priority at the same position; replacements are not rescanned.
	// Empty keys follow strings.Replacer byte boundaries and can split a multibyte UTF-8 rune.
	// Use ReplaceAll for empty-search insertion at UTF-8 sequence boundaries.
	// Similar: ReplaceArray.

	// Example: swap map
	pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
	v := str.Of("Gophers are great!").Swap(pairs).String()
	println(v)
	// #string GoForj is fantastic!
}
