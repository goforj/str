package str

// ChopStart removes the first matching prefix if present.
// @group Affixes
//
// Example: chop start
//
//	v := str.Of("https://goforj.dev").ChopStart("https://", "http://").String()
//	godump.Dump(v)
//	// #string goforj.dev
func (s String) ChopStart(prefixes ...string) String {
	for _, p := range prefixes {
		if p == "" {
			continue
		}
		if len(s.s) >= len(p) && s.s[:len(p)] == p {
			return String{s: s.s[len(p):]}
		}
	}
	return s
}

// ChopEnd removes the first matching suffix if present.
// @group Affixes
//
// Example: chop end
//
//	v := str.Of("file.txt").ChopEnd(".txt", ".md").String()
//	godump.Dump(v)
//	// #string file
func (s String) ChopEnd(suffixes ...string) String {
	for _, suf := range suffixes {
		if suf == "" {
			continue
		}
		if len(s.s) >= len(suf) && s.s[len(s.s)-len(suf):] == suf {
			return String{s: s.s[:len(s.s)-len(suf)]}
		}
	}
	return s
}
