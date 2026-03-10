package str

// BeforeLastFold returns the substring before the last occurrence of sep using Unicode-aware
// case-insensitive comparison. If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice before last separator (case-insensitive)
//
//	v := str.Of("pkg/Path/FILE.txt").BeforeLastFold("/path/").String()
//	println(v)
//	// #string pkg
func (s String) BeforeLastFold(sep string) String {
	start, _, ok := foldMatchRange(s.s, sep, true)
	if !ok {
		return s
	}

	return String{s: s.s[:start]}
}
