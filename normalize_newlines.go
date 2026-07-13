package str

import (
	"strings"
	"unicode/utf8"
)

// NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.
// Similar: Lines.
// @group Cleanup
//
// Example: normalize newline variants
//
//	v := str.Of("a\r\nb\u2028c").NormalizeNewlines().String()
//	println(v)
//	// #string a\nb\nc
func (s String) NormalizeNewlines() String {
	if !strings.ContainsAny(s.s, "\r\u0085\u2028\u2029") {
		return s
	}

	var b strings.Builder
	b.Grow(len(s.s))

	for i := 0; i < len(s.s); {
		r, size := utf8.DecodeRuneInString(s.s[i:])
		switch r {
		case '\r':
			if i+size < len(s.s) && s.s[i+size] == '\n' {
				i += size + 1
			} else {
				i += size
			}
			b.WriteByte('\n')
		case '\u0085', '\u2028', '\u2029':
			b.WriteByte('\n')
			i += size
		default:
			b.WriteRune(r)
			i += size
		}
	}

	return String{s: b.String()}
}
