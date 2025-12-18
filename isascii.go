package str

// IsASCII reports whether the string consists solely of 7-bit ASCII runes.
// @group Checks
//
// Example: ASCII check
//
//	v := str.Of("gopher").IsASCII()
//	godump.Dump(v)
//	// #bool true
func (s String) IsASCII() bool {
	for _, r := range s.s {
		if r > 127 {
			return false
		}
	}
	return true
}
