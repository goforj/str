package str

import "strings"

// PadLeft pads the string on the left to reach length runes using pad (defaults to space).
// @group Padding
//
// Example: pad left
//
//	val := str.Of("go")
//	godump.Dump(val.PadLeft(5, " ").String())
//
//	// #string \u00a0\u00a0\u00a0go
func (s String) PadLeft(length int, pad string) String {
	return padInternal(s.s, length, pad, true, false)
}

// PadRight pads the string on the right to reach length runes using pad (defaults to space).
// @group Padding
//
// Example: pad right
//
//	val := str.Of("go")
//	godump.Dump(val.PadRight(5, ".").String())
//
//	// #string go...
func (s String) PadRight(length int, pad string) String {
	return padInternal(s.s, length, pad, false, true)
}

// PadBoth pads the string on both sides to reach length runes using pad (defaults to space).
// @group Padding
//
// Example: pad both
//
//	val := str.Of("go")
//	godump.Dump(val.PadBoth(6, "-").String())
//
//	// #string --go--
func (s String) PadBoth(length int, pad string) String {
	return padInternal(s.s, length, pad, true, true)
}

func padInternal(s string, length int, pad string, left, right bool) String {
	if length <= 0 {
		return String{s: ""}
	}
	runes := []rune(s)
	if len(runes) >= length {
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
