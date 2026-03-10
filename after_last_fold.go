package str

// AfterLastFold returns the substring after the last occurrence of sep using Unicode-aware
// case-insensitive comparison. If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice after last separator (case-insensitive)
//
//	v := str.Of("pkg/Path/FILE.txt").AfterLastFold("/path/").String()
//	println(v)
//	// #string FILE.txt
func (s String) AfterLastFold(sep string) String {
	_, end, ok := foldMatchRange(s.s, sep, true)
	if !ok {
		return s
	}

	return String{s: s.s[end:]}
}
