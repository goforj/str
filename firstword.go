package str

// FirstWord returns the first detected word or an empty string.
// Similar: LastWord and SplitWords.
// @group Words
//
// Example: first word
//
//	v := str.Of("Hello world")
//	println(v.FirstWord().String())
//	// #string Hello
func (s String) FirstWord() String {
	tokens := tokenizeWords(s.s)
	if len(tokens) == 0 {
		return String{s: ""}
	}
	return String{s: tokens[0].value}
}
