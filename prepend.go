package str

import "strings"

// Prepend concatenates the provided parts to the beginning of the string.
// Similar: Append.
// @group Compose
//
// Example: prepend text
//
//	v := str.Of("World").Prepend("Hello ", "Go ").String()
//	println(v)
//	// #string Hello Go World
func (s String) Prepend(parts ...string) String {
	if len(parts) == 0 {
		return s
	}

	var b strings.Builder
	total := len(s.s)
	for _, part := range parts {
		total += len(part)
	}
	b.Grow(total)
	for _, part := range parts {
		b.WriteString(part)
	}
	b.WriteString(s.s)
	return String{s: b.String()}
}
