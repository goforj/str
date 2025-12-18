package str

// Prepend concatenates the provided parts to the beginning of the string.
// @group Compose
//
// Example: prepend text
//
//	v := str.Of("World").Prepend("Hello ", "Go ").String()
//	godump.Dump(v)
//	// #string Hello Go World
func (s String) Prepend(parts ...string) String {
	out := s.s
	for i := len(parts) - 1; i >= 0; i-- {
		out = parts[i] + out
	}
	return String{s: out}
}
