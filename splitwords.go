package str

// SplitWords splits the string into Unicode words, including camel case and acronym boundaries.
// Similar: FirstWord, LastWord, WordCount, and Words.
// @group Words
//
// Example: split words
//
//	v := str.Of("one, two, three").SplitWords()
//	fmt.Println(v)
//	// #[]string [one two three]
func (s String) SplitWords() []string {
	return wordTokenValues(tokenizeWords(s.s))
}
