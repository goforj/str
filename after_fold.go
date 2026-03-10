package str

// AfterFold returns the substring after the first occurrence of sep using Unicode-aware
// case-insensitive comparison. If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice after marker (case-insensitive)
//
//	v := str.Of("gopher::GO-team").AfterFold("::go").String()
//	println(v)
//	// #string -team
func (s String) AfterFold(sep string) String {
	_, end, ok := foldMatchRange(s.s, sep, false)
	if !ok {
		return s
	}

	return String{s: s.s[end:]}
}
