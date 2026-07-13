package str

// Take returns the first length runes of the string (clamped).
// Similar: TakeLast and Limit.
// @group Substrings
//
// Example: take head
//
//	v := str.Of("gophers").Take(3).String()
//	println(v)
//	// #string gop
func (s String) Take(length int) String {
	if length <= 0 {
		return String{s: ""}
	}
	runes := []rune(s.s)
	if length >= len(runes) {
		return s
	}
	return String{s: string(runes[:length])}
}

// TakeLast returns the last length runes of the string (clamped).
// Similar: Take.
// @group Substrings
//
// Example: take tail
//
//	v := str.Of("gophers").TakeLast(4).String()
//	println(v)
//	// #string hers
func (s String) TakeLast(length int) String {
	if length <= 0 {
		return String{s: ""}
	}
	runes := []rune(s.s)
	if length >= len(runes) {
		return s
	}
	return String{s: string(runes[len(runes)-length:])}
}
