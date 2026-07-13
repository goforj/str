package str

// WordCount returns the number of detected words.
// Similar: SplitWords.
// @group Words
//
// Example: count words
//
//	v := str.Of("Hello, world!").WordCount()
//	println(v)
//	// #int 2
func (s String) WordCount() int {
	return len(tokenizeWords(s.s))
}
