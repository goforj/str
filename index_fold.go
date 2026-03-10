package str

// IndexFold returns the rune index of the first occurrence of sub using Unicode-aware
// case-insensitive comparison, or -1 if not found.
// @group Search
//
// Example: first rune index (case-insensitive)
//
//	v := str.Of("Go gopher GO").IndexFold("go")
//	println(v)
//	// #int 0
func (s String) IndexFold(sub string) int {
	return runeIndexFold(s.s, sub, false)
}
