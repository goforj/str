package str

// Index returns the rune index of the first occurrence of sub, or -1 if not found.
// Similar: LastIndex.
// @group Search
//
// Example: first rune index
//
//	v := str.Of("héllo").Index("llo")
//	println(v)
//	// #int 2
func (s String) Index(sub string) int {
	return runeIndex(s.s, sub, false)
}

// LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.
// Similar: Index.
// @group Search
//
// Example: last rune index
//
//	v := str.Of("go gophers go").LastIndex("go")
//	println(v)
//	// #int 11
func (s String) LastIndex(sub string) int {
	return runeIndex(s.s, sub, true)
}

// Slice returns the substring between rune offsets [start:end).
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
