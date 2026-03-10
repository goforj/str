package str

// BeforeFold returns the substring before the first occurrence of sep using Unicode-aware
// case-insensitive comparison. If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice before marker (case-insensitive)
//
//	v := str.Of("GoPHER::go").BeforeFold("::GO").String()
//	println(v)
//	// #string GoPHER
func (s String) BeforeFold(sep string) String {
	start, _, ok := foldMatchRange(s.s, sep, false)
	if !ok {
		return s
	}

	return String{s: s.s[:start]}
}
