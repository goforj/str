package str

// Deduplicate collapses consecutive instances of char into a single instance.
// If char is zero, space is used.
// @group Cleanup
//
// Example: collapse spaces
//
//	v := str.Of("The   Laravel   Framework").Deduplicate(' ').String()
//	godump.Dump(v)
//	// #string The Laravel Framework
func (s String) Deduplicate(char rune) String {
	if char == 0 {
		char = ' '
	}

	var out []rune
	var prev rune
	for i, r := range s.s {
		if i == 0 || r != char || prev != char {
			out = append(out, r)
		}
		prev = r
	}

	return String{s: string(out)}
}
