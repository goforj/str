package str

// CharAt returns the rune at the given index and true if within bounds.
// @group Substrings
//
// Example: rune lookup
//
//	v, ok := str.Of("gopher").CharAt(2)
//	godump.Dump(string(v), ok)
//	// #string p
//	// #bool true
func (s String) CharAt(idx int) (rune, bool) {
	if idx < 0 {
		return 0, false
	}
	count := 0
	for _, r := range s.s {
		if count == idx {
			return r, true
		}
		count++
	}
	return 0, false
}
