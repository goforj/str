package str

// FirstWord returns the first word token or empty string.
// @group Words
//
// Example: first word
//
//	val := str.Of("Hello world")
//	godump.Dump(val.FirstWord().String())
//
//	// #string Hello
func (s String) FirstWord() String {
	words := splitWordsRunes(s.s)
	if len(words) == 0 {
		return String{s: ""}
	}
	return String{s: words[0]}
}
