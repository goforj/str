package str

// Mask replaces the middle of the string with the given rune, revealing revealLeft runes
// at the start and revealRight runes at the end. Negative reveal values count from the end.
// If the reveal counts cover the whole string, the original string is returned.
// @group Masking
//
// Example: mask email
//
//	v := str.Of("gopher@example.com").Mask('*', 3, 4).String()
//	godump.Dump(v)
//	// #string gop***********.com
func (s String) Mask(mask rune, revealLeft, revealRight int) String {
	runes := []rune(s.s)
	n := len(runes)
	if mask == 0 {
		mask = '*'
	}

	left := revealLeft
	if left < 0 {
		left = n + left
	}
	if left < 0 {
		left = 0
	}
	if left > n {
		left = n
	}

	right := revealRight
	if right < 0 {
		right = n + right
	}
	if right < 0 {
		right = 0
	}
	if right > n {
		right = n
	}

	if left+right >= n {
		return s
	}

	midLen := n - left - right
	masked := make([]rune, 0, n)
	masked = append(masked, runes[:left]...)
	for i := 0; i < midLen; i++ {
		masked = append(masked, mask)
	}
	masked = append(masked, runes[n-right:]...)

	return String{s: string(masked)}
}
