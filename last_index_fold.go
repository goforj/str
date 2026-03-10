package str

// LastIndexFold returns the rune index of the last occurrence of sub using Unicode-aware
// case-insensitive comparison, or -1 if not found.
// @group Search
//
// Example: last rune index (case-insensitive)
//
//	v := str.Of("Go gopher GO").LastIndexFold("go")
//	println(v)
//	// #int 10
func (s String) LastIndexFold(sub string) int {
	return runeIndexFold(s.s, sub, true)
}
