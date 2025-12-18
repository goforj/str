package str

// ChopStart removes the first matching prefix if present.
// @group Affixes
//
// Example: chop start
//
//	val := str.Of("https://goforj.dev")
//	godump.Dump(val.ChopStart("https://", "http://").String())
//
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
//	val := str.Of("file.txt")
//	godump.Dump(val.ChopEnd(".txt", ".md").String())
//
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
