package str

import "strings"

// Join concatenates elements with sep and returns the result to the fluent chain.
// The receiver provides fluent access and is not included in elements.
// Similar: Split.
// @group Words
//
// Example: join words
//
//	v := str.Of("").Join([]string{"foo", "bar"}, "-").String()
//	println(v)
//	// #string foo-bar
func (s String) Join(elements []string, sep string) String {
	return String{s: strings.Join(elements, sep)}
}
