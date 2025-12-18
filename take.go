package str

// Take returns the first length runes of the string (clamped).
// @group Substrings
//
// Example: take head
//
//	val := str.Of("gophers")
//	godump.Dump(val.Take(3).String())
//
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
// @group Substrings
//
// Example: take tail
//
//	val := str.Of("gophers")
//	godump.Dump(val.TakeLast(4).String())
//
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
