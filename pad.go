package str

import "strings"

// PadLeft pads the string on the left to reach length runes using pad (defaults to space).
// Widths at or below the current rune width leave the string unchanged.
// Similar: PadRight and PadBoth.
// @group Padding
//
// Example: pad left
//
//	v := str.Of("go").PadLeft(5, " ").String()
//	println(v)
//	// #string \u0020\u0020\u0020go
func (s String) PadLeft(length int, pad string) String {
	return padInternal(s.s, length, pad, true, false)
}

// PadRight pads the string on the right to reach length runes using pad (defaults to space).
// Widths at or below the current rune width leave the string unchanged.
// Similar: PadLeft and PadBoth.
// @group Padding
//
// Example: pad right
//
//	v := str.Of("go").PadRight(5, ".").String()
//	println(v)
//	// #string go...
func (s String) PadRight(length int, pad string) String {
	return padInternal(s.s, length, pad, false, true)
}

// PadBoth pads the string on both sides to reach length runes using pad (defaults to space).
// Widths at or below the current rune width leave the string unchanged.
// Similar: PadLeft and PadRight.
// @group Padding
//
// Example: pad both
//
//	v := str.Of("go").PadBoth(6, "-").String()
//	println(v)
//	// #string --go--
func (s String) PadBoth(length int, pad string) String {
	return padInternal(s.s, length, pad, true, true)
}

// padInternal centralizes width handling so every padding direction preserves
// the original value when the requested rune width has already been met.
func padInternal(s string, length int, pad string, left, right bool) String {
	runes := []rune(s)
	if length <= len(runes) {
		return String{s: s}
	}
	if pad == "" {
		pad = " "
	}
	padRunes := []rune(pad)
	add := length - len(runes)

	leftPad := 0
	rightPad := 0
	if left && right {
		leftPad = add / 2
		rightPad = add - leftPad
	} else if left {
		leftPad = add
	} else if right {
		rightPad = add
	}

	build := func(n int) string {
		if n <= 0 {
			return ""
		}
		var b strings.Builder
		b.Grow(n)
		for i := 0; i < n; i++ {
			b.WriteRune(padRunes[i%len(padRunes)])
		}
		return b.String()
	}

	return String{s: build(leftPad) + s + build(rightPad)}
}
