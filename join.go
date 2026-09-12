package str

import "strings"

// Join concatenates elements with sep and returns the result to the fluent chain.
// Join starts a chain from a slice without discarding an existing receiver.
// Similar: Split.
// @group Words
//
// Example: join words
//
//	v := str.Join([]string{"foo", "bar"}, "-").String()
//	println(v)
//	// #string foo-bar
func Join(elements []string, sep string) String {
	return String{s: strings.Join(elements, sep)}
}
