package str

// LastWord returns the last detected word or an empty string.
// Similar: FirstWord and SplitWords.
// @group Words
//
// Example: last word
//
//	v := str.Of("Hello world").LastWord().String()
//	println(v)
//	// #string world
func (s String) LastWord() String {
	tokens := tokenizeWords(s.s)
	if len(tokens) == 0 {
		return String{s: ""}
	}
	return String{s: tokens[len(tokens)-1].value}
}
