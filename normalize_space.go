package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// NormalizeSpace removes surrounding whitespace and collapses internal whitespace to single spaces.
// Similar: Trim.
// @group Cleanup
//
// Example: normalize whitespace
//
//	v := str.Of("  go   forj  ").NormalizeSpace().String()
//	println(v)
//	// #string go forj
func (s String) NormalizeSpace() String {
	trimmed := strings.TrimSpace(s.s)
	if trimmed == "" {
		return String{s: trimmed}
	}

	isASCII, needsRewrite := normalizeSpaceState(trimmed)
	if !needsRewrite {
		if len(trimmed) == len(s.s) {
			return s
		}
		// Clone prevents a small trimmed result from retaining the full source buffer.
		return String{s: strings.Clone(trimmed)}
	}
	if isASCII {
		return String{s: normalizeSpaceASCII(trimmed)}
	}
	return String{s: normalizeSpaceUnicode(trimmed)}
}

// normalizeSpaceState reports whether value is ASCII and whether its whitespace or encoding needs normalization.
func normalizeSpaceState(value string) (bool, bool) {
	isASCII := true
	previousSpace := false
	needsRewrite := false
	for i := 0; i < len(value); {
		c := value[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			r, size := utf8.DecodeRuneInString(value[i:])
			if unicode.IsSpace(r) {
				return false, true
			}
			// Rune iteration historically replaces each invalid byte with utf8.RuneError.
			if r == utf8.RuneError && size == 1 {
				return false, true
			}
			previousSpace = false
			i += size
			continue
		}
		if isNormalizeSpaceASCIIByte(c) {
			if c != ' ' || previousSpace {
				needsRewrite = true
			}
			previousSpace = true
			i++
			continue
		}
		previousSpace = false
		i++
	}
	return isASCII, needsRewrite
}

// normalizeSpaceASCII collapses whitespace in an ASCII value that is already trimmed.
func normalizeSpaceASCII(value string) string {
	var out strings.Builder
	out.Grow(len(value))
	pendingSpace := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		if isNormalizeSpaceASCIIByte(c) {
			pendingSpace = true
			continue
		}
		if pendingSpace {
			out.WriteByte(' ')
			pendingSpace = false
		}
		out.WriteByte(c)
	}
	return out.String()
}

// isNormalizeSpaceASCIIByte reports whether c is whitespace according to Unicode's ASCII subset.
func isNormalizeSpaceASCIIByte(c byte) bool {
	return c == ' ' || '\t' <= c && c <= '\r'
}

// normalizeSpaceUnicode collapses whitespace while preserving Unicode text and historical invalid-UTF-8 normalization.
func normalizeSpaceUnicode(value string) string {
	var out strings.Builder
	out.Grow(len(value))
	pendingSpace := false

	for _, r := range value {
		if unicode.IsSpace(r) {
			pendingSpace = true
			continue
		}
		if pendingSpace {
			out.WriteByte(' ')
			pendingSpace = false
		}
		out.WriteRune(r)
	}

	return out.String()
}
