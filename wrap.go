package str

// Wrap surrounds the string with before and after.
// Similar: Unwrap.
// @group Affixes
//
// Example: wrap string
//
//	v := str.Of("GoForj").Wrap(`"`, `"`).String()
//	println(v)
//	// #string "GoForj"
func (s String) Wrap(before, after string) String {
	return String{s: before + s.s + after}
}

// Unwrap removes matching before and after strings if present.
// Similar: Wrap.
// @group Affixes
//
// Example: unwrap string
//
//	v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
//	println(v)
//	// #string GoForj
func (s String) Unwrap(before, after string) String {
	if len(s.s) >= len(before)+len(after) &&
		s.s[:len(before)] == before &&
		s.s[len(s.s)-len(after):] == after {
		return String{s: s.s[len(before) : len(s.s)-len(after)]}
	}
	return s
}
