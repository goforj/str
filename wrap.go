package str

// Wrap surrounds the string with before and after (after defaults to before).
// @group Affixes
//
// Example: wrap string
//
//	v := str.Of("GoForj").Wrap(`"`, "").String()
//	str.Dump(v)
//	// #string "GoForj"
func (s String) Wrap(before, after string) String {
	if after == "" {
		after = before
	}
	return String{s: before + s.s + after}
}

// Unwrap removes matching before and after strings if present.
// @group Affixes
//
// Example: unwrap string
//
//	v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
//	str.Dump(v)
//	// #string GoForj
func (s String) Unwrap(before, after string) String {
	if after == "" {
		after = before
	}
	if len(s.s) >= len(before)+len(after) &&
		s.s[:len(before)] == before &&
		s.s[len(s.s)-len(after):] == after {
		return String{s: s.s[len(before) : len(s.s)-len(after)]}
	}
	return s
}
