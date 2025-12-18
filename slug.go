package str

import (
	"strings"
	"unicode"
)

// Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
// Not locale-aware; intended for identifiers/URLs.
// @group Slug
//
// Example: build slug
//
//	val := str.Of("Laravel Framework")
//	godump.Dump(val.Slug("-").String())
//
//	// #string laravel-framework
func (s String) Slug(sep string) String {
	if sep == "" {
		sep = "-"
	}

	var out []rune
	lastSep := false

	for _, r := range s.s {
		if mapped, ok := transliterateBasic(r); ok {
			for _, mr := range mapped {
				if unicode.IsLetter(mr) || unicode.IsDigit(mr) {
					out = append(out, unicode.ToLower(mr))
				}
			}
			lastSep = false
			continue
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out = append(out, unicode.ToLower(r))
			lastSep = false
			continue
		}
		if !lastSep {
			out = append(out, []rune(sep)...)
			lastSep = true
		}
	}

	// Trim leading/trailing separator sequences
	result := strings.Trim(string(out), sep)
	return String{s: result}
}

// transliterateBasic handles a minimal set of accented runes to ASCII.
func transliterateBasic(r rune) (string, bool) {
	switch r {
	case 'à', 'á', 'â', 'ã', 'ä', 'å', 'æ':
		return "a", true
	case 'ç':
		return "c", true
	case 'è', 'é', 'ê', 'ë':
		return "e", true
	case 'ì', 'í', 'î', 'ï':
		return "i", true
	case 'ñ':
		return "n", true
	case 'ò', 'ó', 'ô', 'õ', 'ö', 'ø':
		return "o", true
	case 'ß':
		return "ss", true
	case 'ù', 'ú', 'û', 'ü':
		return "u", true
	case 'ý', 'ÿ':
		return "y", true
	default:
		return "", false
	}
}
