package str

import (
	"strings"
	"unicode"
)

var smallWords = map[string]struct{}{
	"a": {}, "an": {}, "and": {}, "as": {}, "at": {}, "but": {}, "by": {}, "for": {},
	"in": {}, "of": {}, "on": {}, "or": {}, "the": {}, "to": {}, "with": {},
}

// Headline converts the string into a human-friendly headline:
// splits on case/underscores/dashes/whitespace, title-cases words, and lowercases small words (except the first).
// @group Case
//
// Example: headline
//
//	v := str.Of("emailNotification_sent").Headline().String()
//	godump.Dump(v)
//	// #string Email Notification Sent
func (s String) Headline() String {
	words := splitWordsRunes(s.s)
	for i, w := range words {
		lower := strings.ToLower(w)
		if i != 0 {
			if _, ok := smallWords[lower]; ok {
				words[i] = lower
				continue
			}
		}
		runes := []rune(lower)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return String{s: strings.Join(words, " ")}
}
