package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Plural returns a best-effort English plural form of the last word.
// @group Pluralize
//
// Example: pluralize word
//
//	v := str.Of("city").Plural().String()
//	str.Dump(v)
//	// #string cities
func (s String) Plural() String {
	if s.s == "" {
		return s
	}
	return String{s: transformLastWord(s.s, pluralizeWord)}
}

// Singular returns a best-effort English singular form of the last word.
// @group Pluralize
//
// Example: singularize word
//
//	v := str.Of("people").Singular().String()
//	str.Dump(v)
//	// #string person
func (s String) Singular() String {
	if s.s == "" {
		return s
	}
	return String{s: transformLastWord(s.s, singularizeWord)}
}

var uncountables = map[string]struct{}{
	"aircraft":    {},
	"deer":        {},
	"equipment":   {},
	"fish":        {},
	"hardware":    {},
	"information": {},
	"media":       {},
	"moose":       {},
	"money":       {},
	"news":        {},
	"police":      {},
	"rice":        {},
	"salmon":      {},
	"sheep":       {},
	"software":    {},
	"species":     {},
	"series":      {},
	"staff":       {},
}

var irregularSingular = map[string]string{
	"analysis": "analyses",
	"axis":   "axes",
	"basis":  "bases",
	"cactus": "cacti",
	"child":  "children",
	"crisis": "crises",
	"datum":  "data",
	"echo":   "echoes",
	"foot":   "feet",
	"goose":  "geese",
	"hero":   "heroes",
	"louse":  "lice",
	"mouse":  "mice",
	"ox":     "oxen",
	"man":    "men",
	"person": "people",
	"quiz":   "quizzes",
	"thesis": "theses",
	"tooth":  "teeth",
	"torpedo": "torpedoes",
	"woman":  "women",
}

var irregularPlural = map[string]string{
	"analyses": "analysis",
	"axes":     "axis",
	"bases":    "basis",
	"cacti":    "cactus",
	"children": "child",
	"crises":   "crisis",
	"data":     "datum",
	"echoes":   "echo",
	"feet":     "foot",
	"geese":    "goose",
	"heroes":   "hero",
	"lice":     "louse",
	"dies":     "die",
	"lies":     "lie",
	"men":      "man",
	"mice":     "mouse",
	"oxen":     "ox",
	"pies":     "pie",
	"people":   "person",
	"quizzes":  "quiz",
	"teeth":    "tooth",
	"theses":   "thesis",
	"ties":     "tie",
	"torpedoes": "torpedo",
	"vies":     "vie",
	"women":    "woman",
}

var fToVesExceptions = map[string]struct{}{
	"belief": {},
	"cafe":   {},
	"chef":   {},
	"chief":  {},
	"cliff":  {},
	"proof":  {},
	"roof":   {},
	"safe":   {},
}

var feSingulars = map[string]struct{}{
	"knife": {},
	"life":  {},
	"wife":  {},
}

func pluralizeWord(word string) string {
	if word == "" || !hasLetter(word) {
		return word
	}

	lower := strings.ToLower(word)
	if _, ok := uncountables[lower]; ok {
		return word
	}
	if plural, ok := irregularSingular[lower]; ok {
		return applyCase(word, plural)
	}
	if _, ok := irregularPlural[lower]; ok {
		return word
	}
	if strings.HasSuffix(lower, "datum") && len(lower) > len("datum") {
		return applyCase(word, lower[:len(lower)-len("datum")]+"data")
	}

	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		if !isVowel(rune(lower[len(lower)-2])) {
			return applyCase(word, lower[:len(lower)-1]+"ies")
		}
		return applyCase(word, lower+"s")
	}

	if strings.HasSuffix(lower, "ch") || strings.HasSuffix(lower, "sh") ||
		strings.HasSuffix(lower, "s") || strings.HasSuffix(lower, "x") || strings.HasSuffix(lower, "z") {
		return applyCase(word, lower+"es")
	}

	if strings.HasSuffix(lower, "fe") {
		if _, ok := fToVesExceptions[lower]; ok {
			return applyCase(word, lower+"s")
		}
		return applyCase(word, lower[:len(lower)-2]+"ves")
	}
	if strings.HasSuffix(lower, "f") {
		if _, ok := fToVesExceptions[lower]; ok {
			return applyCase(word, lower+"s")
		}
		return applyCase(word, lower[:len(lower)-1]+"ves")
	}

	return applyCase(word, lower+"s")
}

func singularizeWord(word string) string {
	if word == "" || !hasLetter(word) {
		return word
	}

	lower := strings.ToLower(word)
	if _, ok := uncountables[lower]; ok {
		return word
	}
	if singular, ok := irregularPlural[lower]; ok {
		return applyCase(word, singular)
	}
	if _, ok := irregularSingular[lower]; ok {
		return word
	}
	if strings.HasSuffix(lower, "data") && len(lower) > len("data") {
		return applyCase(word, lower[:len(lower)-len("data")]+"datum")
	}

	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		before := rune(lower[len(lower)-4])
		if isVowel(before) {
			return applyCase(word, lower[:len(lower)-1])
		}
		return applyCase(word, lower[:len(lower)-3]+"y")
	}

	if strings.HasSuffix(lower, "ches") || strings.HasSuffix(lower, "shes") ||
		strings.HasSuffix(lower, "xes") || strings.HasSuffix(lower, "zes") || strings.HasSuffix(lower, "ses") {
		return applyCase(word, lower[:len(lower)-2])
	}

	if strings.HasSuffix(lower, "ves") && len(lower) > 3 {
		base := lower[:len(lower)-3] + "f"
		if _, ok := feSingulars[base+"e"]; ok {
			return applyCase(word, base+"e")
		}
		return applyCase(word, base)
	}

	if strings.HasSuffix(lower, "s") && !strings.HasSuffix(lower, "ss") && len(lower) > 1 {
		return applyCase(word, lower[:len(lower)-1])
	}

	return word
}

func transformLastWord(s string, fn func(string) string) string {
	i := len(s)

	for i > 0 {
		r, size := utf8.DecodeLastRuneInString(s[:i])
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			end := i
			i -= size
			for i > 0 {
				r, size = utf8.DecodeLastRuneInString(s[:i])
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					break
				}
				i -= size
			}
			start := i
			word := s[start:end]
			if split := lastCamelWordSplit(word); split != "" {
				prefix := word[:len(word)-len(split)]
				return s[:start] + prefix + fn(split) + s[end:]
			}
			return s[:start] + fn(word) + s[end:]
		}
		i -= size
	}

	return s
}

func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func applyCase(original, replacement string) string {
	if isAllUpper(original) {
		return strings.ToUpper(replacement)
	}
	if isAllLower(original) {
		return strings.ToLower(replacement)
	}
	if isTitleCase(original) {
		return titleCase(replacement)
	}
	return replacement
}

func isAllUpper(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetter
}

func isAllLower(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetter
}

func isTitleCase(s string) bool {
	first := true
	for _, r := range s {
		if unicode.IsLetter(r) {
			if first {
				if !unicode.IsUpper(r) {
					return false
				}
				first = false
				continue
			}
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return !first
}

func titleCase(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func lastCamelWordSplit(word string) string {
	runes := []rune(word)
	if len(runes) < 2 {
		return ""
	}
	for i := len(runes) - 1; i > 0; i-- {
		if unicode.IsUpper(runes[i]) && (unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1])) {
			return string(runes[i:])
		}
	}
	return ""
}
