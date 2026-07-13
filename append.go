package str

import "strings"

// Append concatenates the provided parts to the end of the string.
// Similar: Prepend.
// @group Compose
//
// Example: append text
//
//	v := str.Of("Go").Append("Forj", "!").String()
//	println(v)
//	// #string GoForj!
func (s String) Append(parts ...string) String {
	if len(parts) == 0 {
		return s
	}

	var b strings.Builder
	total := len(s.s)
	for _, part := range parts {
		total += len(part)
	}
	b.Grow(total)
	b.WriteString(s.s)
	for _, part := range parts {
		b.WriteString(part)
	}
	return String{s: b.String()}
}
