package str

// Words limits the string to count words, preserving the source through the
// selected word boundary and appending suffix if truncated.
// Similar: SplitWords and WrapWords.
// @group Words
//
// Example: limit words
//
//	v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
//	println(v)
//	// #string Perfectly balanced, as >>>
func (s String) Words(count int, suffix string) String {
	if count <= 0 {
		return String{s: ""}
	}
	tokens := tokenizeWords(s.s)
	if len(tokens) <= count {
		return s
	}
	return String{s: s.s[:tokens[count-1].end] + suffix}
}
