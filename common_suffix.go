package str

// CommonSuffix returns the longest shared suffix between the string and all provided others.
// Comparison is rune-safe. If no others are provided, the original string is returned.
// @group Substrings
//
// Example: longest common suffix
//
//	v := str.Of("main_test.go").CommonSuffix("user_test.go", "api_test.go").String()
//	println(v)
//	// #string _test.go
func (s String) CommonSuffix(others ...string) String {
	if len(others) == 0 {
		return s
	}

	common := []rune(s.s)
	for _, other := range others {
		otherRunes := []rune(other)
		limit := min(len(common), len(otherRunes))
		i := 0
		for i < limit && common[len(common)-1-i] == otherRunes[len(otherRunes)-1-i] {
			i++
		}
		common = common[len(common)-i:]
		if len(common) == 0 {
			return String{s: ""}
		}
	}

	return String{s: string(common)}
}
