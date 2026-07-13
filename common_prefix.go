package str

// CommonPrefix returns the longest shared prefix between the string and all provided others.
// Comparison is rune-safe. If no others are provided, the original string is returned.
// Similar: CommonSuffix.
// @group Substrings
//
// Example: longest common prefix
//
//	v := str.Of("gopher").CommonPrefix("go", "gold").String()
//	println(v)
//	// #string go
func (s String) CommonPrefix(others ...string) String {
	if len(others) == 0 {
		return s
	}

	common := []rune(s.s)
	for _, other := range others {
		otherRunes := []rune(other)
		limit := min(len(common), len(otherRunes))
		i := 0
		for i < limit && common[i] == otherRunes[i] {
			i++
		}
		common = common[:i]
		if len(common) == 0 {
			return String{s: ""}
		}
	}

	return String{s: string(common)}
}
