package str

// WordCount returns the number of word tokens (letters/digits runs).
// @group Words
//
// Example: count words
//
//	v := str.Of("Hello, world!")
//	godump.Dump(v.WordCount())
//	// #int 2
func (s String) WordCount() int {
	return len(splitWordsRunes(s.s))
}
