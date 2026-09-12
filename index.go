package str

import "strings"

// Index returns the byte index of the first occurrence of sub, or -1 if not found.
// Similar: LastIndex.
// @group Search
//
// Example: first byte index
//
//	v := str.Of("héllo").Index("llo")
//	println(v)
//	// #int 3
func (s String) Index(sub string) int {
	return strings.Index(s.s, sub)
}

// LastIndex returns the byte index of the last occurrence of sub, or -1 if not found.
// Similar: Index.
// @group Search
//
// Example: last byte index
//
//	v := str.Of("go gophers go").LastIndex("go")
//	println(v)
//	// #int 11
func (s String) LastIndex(sub string) int {
	return strings.LastIndex(s.s, sub)
}

// Slice returns the substring between rune offsets [start:end).
// Index and LastIndex return bytes, which must not be passed directly here.
// Indices are clamped; if start >= end the result is empty.
// @group Substrings
//
// Example: rune-safe slice
//
//	v := str.Of("naïve café").Slice(3, 7).String()
//	println(v)
//	// #string ve c
func (s String) Slice(start, end int) String {
	runes := []rune(s.s)
	start, end = clampRange(start, end, len(runes))
	if start >= end {
		return String{s: ""}
	}
	return String{s: string(runes[start:end])}
}
