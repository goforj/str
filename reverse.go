package str

// Reverse returns a rune-safe reversed string.
// @group Transform
//
// Example: reverse
//
//	val := str.Of("naïve")
//	godump.Dump(val.Reverse().String())
//
//	// #string evïan
func (s String) Reverse() String {
	runes := []rune(s.s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return String{s: string(runes)}
}
