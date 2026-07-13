package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Plural returns a best-effort English plural form of the final identifier word.
// It handles common English forms and identifier boundaries without claiming to
// resolve every irregular or ambiguous noun.
// Similar: Singular.
// @group Pluralize
//
// Example: pluralize word
//
//	v := str.Of("city").Plural().String()
//	println(v)
//	// #string cities
func (s String) Plural() String {
	if s.s == "" {
		return s
	}
	return String{s: transformLastWord(s.s, pluralizeWord)}
}

// Singular returns a best-effort English singular form of the final identifier word.
// It handles common English forms and identifier boundaries without claiming to
// resolve every irregular or ambiguous noun.
// Similar: Plural.
// @group Pluralize
//
// Example: singularize word
//
//	v := str.Of("people").Singular().String()
//	println(v)
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
	"feedback":    {},
	"fish":        {},
	"hardware":    {},
	"information": {},
	"media":       {},
	"metadata":    {},
	"moose":       {},
	"money":       {},
	"news":        {},
	"offspring":   {},
	"police":      {},
	"rice":        {},
	"research":    {},
	"salmon":      {},
	"sheep":       {},
	"software":    {},
	"species":     {},
	"series":      {},
	"staff":       {},
	"traffic":     {},
}

var irregularPlurals = map[string]string{
	"analysis":    "analyses",
	"appendix":    "appendices",
	"axis":        "axes",
	"basis":       "bases",
	"cactus":      "cacti",
	"child":       "children",
	"crisis":      "crises",
	"criterion":   "criteria",
	"datum":       "data",
	"diagnosis":   "diagnoses",
	"echo":        "echoes",
	"foot":        "feet",
	"goose":       "geese",
	"hero":        "heroes",
	"index":       "indices",
	"louse":       "lice",
	"man":         "men",
	"matrix":      "matrices",
	"mouse":       "mice",
	"ox":          "oxen",
	"parenthesis": "parentheses",
	"person":      "people",
	"phenomenon":  "phenomena",
	"potato":      "potatoes",
	"prognosis":   "prognoses",
	"quiz":        "quizzes",
	"synopsis":    "synopses",
	"synthesis":   "syntheses",
	"thesis":      "theses",
	"tomato":      "tomatoes",
	"tooth":       "teeth",
	"torpedo":     "torpedoes",
	"vertex":      "vertices",
	"woman":       "women",
}

var irregularSingulars = reverseInflections(irregularPlurals)

var regularSingularSEndings = map[string]struct{}{
	"alias":      {},
	"apparatus":  {},
	"atlas":      {},
	"bias":       {},
	"bonus":      {},
	"bus":        {},
	"campus":     {},
	"canvas":     {},
	"census":     {},
	"chorus":     {},
	"cosmos":     {},
	"focus":      {},
	"gas":        {},
	"lens":       {},
	"metropolis": {},
	"octopus":    {},
	"plus":       {},
	"prospectus": {},
	"status":     {},
	"virus":      {},
	"walrus":     {},
	"yes":        {},
}

var ieSingulars = map[string]struct{}{
	"auntie":  {},
	"birdie":  {},
	"brownie": {},
	"calorie": {},
	"cookie":  {},
	"cutie":   {},
	"die":     {},
	"foodie":  {},
	"freebie": {},
	"groupie": {},
	"hippie":  {},
	"hoodie":  {},
	"junkie":  {},
	"lie":     {},
	"movie":   {},
	"newbie":  {},
	"pie":     {},
	"selfie":  {},
	"tie":     {},
	"vie":     {},
	"zombie":  {},
}

var vesInflections = map[string]string{
	"calf":  "calves",
	"dwarf": "dwarves",
	"elf":   "elves",
	"half":  "halves",
	"hoof":  "hooves",
	"knife": "knives",
	"leaf":  "leaves",
	"life":  "lives",
	"loaf":  "loaves",
	"scarf": "scarves",
	"self":  "selves",
	"shelf": "shelves",
	"thief": "thieves",
	"wife":  "wives",
	"wolf":  "wolves",
}

// reverseInflections derives the reverse lookup so irregular pairs cannot drift
// apart as practical identifier vocabulary is added.
func reverseInflections(inflections map[string]string) map[string]string {
	reversed := make(map[string]string, len(inflections))
	for singular, plural := range inflections {
		reversed[plural] = singular
	}
	return reversed
}

// pluralizeVes uses an explicit reversible vocabulary because treating every
// final f as irregular corrupts common identifiers such as ref and proof.
func pluralizeVes(word string) (string, bool) {
	for singular, plural := range vesInflections {
		if strings.HasSuffix(word, singular) {
			return strings.TrimSuffix(word, singular) + plural, true
		}
	}
	return "", false
}

// singularizeVes mirrors pluralizeVes so regular words such as archive and
// valve lose only their ordinary trailing s.
func singularizeVes(word string) (string, bool) {
	for singular, plural := range vesInflections {
		if strings.HasSuffix(word, plural) {
			return strings.TrimSuffix(word, plural) + singular, true
		}
	}
	return "", false
}

// pluralizeWord keeps exceptions ahead of broad suffix rules because English
// forms such as status and users are otherwise indistinguishable by suffix alone.
func pluralizeWord(word string) string {
	if word == "" || !hasLetter(word) {
		return word
	}

	lower := strings.ToLower(word)
	if _, ok := uncountables[lower]; ok {
		return word
	}
	if plural, ok := irregularPlurals[lower]; ok {
		return applyCase(word, plural)
	}
	if _, ok := irregularSingulars[lower]; ok {
		return word
	}
	if strings.HasSuffix(lower, "datum") && len(lower) > len("datum") {
		return applyCase(word, lower[:len(lower)-len("datum")]+"data")
	}
	if _, ok := regularSingularSEndings[lower]; ok {
		return applyCase(word, lower+"es")
	}
	if strings.HasSuffix(lower, "ss") {
		return applyCase(word, lower+"es")
	}
	if strings.HasSuffix(lower, "s") {
		return word
	}

	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		stem := lower[:len(lower)-1]
		beforeY, _ := utf8.DecodeLastRuneInString(stem)
		if !isVowel(beforeY) {
			return applyCase(word, lower[:len(lower)-1]+"ies")
		}
		return applyCase(word, lower+"s")
	}

	if strings.HasSuffix(lower, "ch") || strings.HasSuffix(lower, "sh") ||
		strings.HasSuffix(lower, "x") || strings.HasSuffix(lower, "z") {
		return applyCase(word, lower+"es")
	}

	if plural, ok := pluralizeVes(lower); ok {
		return applyCase(word, plural)
	}

	return applyCase(word, lower+"s")
}

// singularizeWord prefers reversible, known forms before broad suffix removal
// so already-singular identifiers are not shortened merely because they end in s.
func singularizeWord(word string) string {
	if word == "" || !hasLetter(word) {
		return word
	}

	lower := strings.ToLower(word)
	if _, ok := uncountables[lower]; ok {
		return word
	}
	if singular, ok := irregularSingulars[lower]; ok {
		return applyCase(word, singular)
	}
	if _, ok := irregularPlurals[lower]; ok {
		return word
	}
	if _, ok := regularSingularSEndings[lower]; ok {
		return word
	}
	if strings.HasSuffix(lower, "ss") {
		return word
	}
	if strings.HasSuffix(lower, "data") && len(lower) > len("data") {
		return applyCase(word, lower[:len(lower)-len("data")]+"datum")
	}

	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		ieCandidate := lower[:len(lower)-1]
		if _, ok := ieSingulars[ieCandidate]; ok {
			return applyCase(word, ieCandidate)
		}
		before, _ := utf8.DecodeLastRuneInString(lower[:len(lower)-3])
		if isVowel(before) {
			return applyCase(word, ieCandidate)
		}
		return applyCase(word, lower[:len(lower)-3]+"y")
	}

	if strings.HasSuffix(lower, "ches") || strings.HasSuffix(lower, "shes") ||
		strings.HasSuffix(lower, "xes") || strings.HasSuffix(lower, "zes") || strings.HasSuffix(lower, "sses") {
		return applyCase(word, lower[:len(lower)-2])
	}
	if strings.HasSuffix(lower, "ses") {
		sCandidate := lower[:len(lower)-2]
		if _, ok := regularSingularSEndings[sCandidate]; ok {
			return applyCase(word, sCandidate)
		}
		return applyCase(word, lower[:len(lower)-1])
	}

	if singular, ok := singularizeVes(lower); ok {
		return applyCase(word, singular)
	}

	if strings.HasSuffix(lower, "s") && !strings.HasSuffix(lower, "ss") && len(lower) > 1 {
		return applyCase(word, lower[:len(lower)-1])
	}

	return word
}

// transformLastWord limits inflection to the final identifier component so
// prefixes and punctuation remain byte-for-byte unchanged.
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

// hasLetter prevents numeric and punctuation-only values from receiving an
// English noun suffix.
func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// isVowel intentionally recognizes only English vowels because these suffix
// rules describe English nouns rather than general Unicode morphology.
func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

// applyCase restores the casing style of an identifier component after rules
// operate on its lowercase representation.
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

// isAllUpper distinguishes acronym-style components for applyCase.
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

// isAllLower distinguishes conventional lowercase components for applyCase.
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

// isTitleCase recognizes a single leading uppercase letter so generated type
// names retain their conventional casing.
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

// titleCase rebuilds mapped irregular forms in the title-case style recognized
// by isTitleCase.
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

// lastCamelWordSplit isolates the final camel-case word, including boundaries
// after initialisms such as APIStatus, so only the noun is inflected.
func lastCamelWordSplit(word string) string {
	runes := []rune(word)
	if len(runes) < 2 {
		return ""
	}
	for i := len(runes) - 1; i > 0; i-- {
		if !unicode.IsUpper(runes[i]) {
			continue
		}
		if unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1]) {
			return string(runes[i:])
		}
		if i+1 < len(runes) && unicode.IsUpper(runes[i-1]) && unicode.IsLower(runes[i+1]) {
			return string(runes[i:])
		}
	}
	return ""
}
