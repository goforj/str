package str

import "unicode"

// UcWords uppercases the first rune of each word, leaving the rest unchanged.
// Words are sequences of letters/digits.
// @group Case
//
// Example: uppercase each word start
//
//	v := str.Of("hello WORLD").UcWords().String()
//	godump.Dump(v)
//	// #string Hello WORLD
func (s String) UcWords() String {
	var out []rune
	inWord := false

	for _, r := range s.s {
		isWord := unicode.IsLetter(r) || unicode.IsDigit(r)
		if isWord && !inWord {
			out = append(out, unicode.ToUpper(r))
		} else {
			out = append(out, r)
		}
		inWord = isWord
	}

	return String{s: string(out)}
}
