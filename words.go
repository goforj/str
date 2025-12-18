package str

import "strings"

// Words limits the string to count words, appending suffix if truncated.
// @group Words
//
// Example: limit words
//
//	v := str.Of("Perfectly balanced, as all things should be.")
//	godump.Dump(v.Words(3, " >>>").String())
//	// #string Perfectly balanced as >>>
func (s String) Words(count int, suffix string) String {
	if count <= 0 {
		return String{s: ""}
	}
	words := splitWordsRunes(s.s)
	if len(words) <= count {
		return s
	}
	return String{s: strings.Join(words[:count], " ") + suffix}
}
