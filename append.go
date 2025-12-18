package str

// Append concatenates the provided parts to the end of the string.
// @group Compose
//
// Example: append text
//
//	v := str.Of("Go").Append("Forj", "!").String()
//	godump.Dump(v)
//	// #string GoForj!
func (s String) Append(parts ...string) String {
	out := s.s
	for _, p := range parts {
		out += p
	}
	return String{s: out}
}
