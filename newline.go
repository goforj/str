package str

// NewLine appends a newline character to the string.
// @group Compose
//
// Example: append newline
//
//	v := str.Of("hello").NewLine().Append("world").String()
//	godump.Dump(v)
//	// #string hello\nworld
func (s String) NewLine() String {
	return s.Append("\n")
}
