package str

// Limit truncates the string to length runes, appending suffix if truncation occurs.
// @group Substrings
//
// Example: limit with suffix
//
//	v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
//	println(v)
//	// #string Perfectly\u0020...
func (s String) Limit(length int, suffix string) String {
	if length <= 0 {
		if suffix != "" {
			return String{s: suffix}
		}
		return String{s: ""}
	}
	runes := []rune(s.s)
	if len(runes) <= length {
		return s
	}
	return String{s: string(runes[:length]) + suffix}
}
