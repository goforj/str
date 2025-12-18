package str

// SplitWords splits the string into words (Unicode letters/digits runs).
// @group Words
//
// Example: split words
//
//	v := str.Of("one, two, three")
//	godump.Dump(v.SplitWords())
//	// #[]string [one two three]
func (s String) SplitWords() []string {
	return splitWordsRunes(s.s)
}
