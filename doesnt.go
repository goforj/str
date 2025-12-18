package str

import "strings"

// DoesntContain reports true if none of the substrings are found (case-sensitive).
// Empty substrings are ignored.
// @group Search
//
// Example: doesn't contain any
//
//	v := str.Of("gophers are great")
//	godump.Dump(v.DoesntContain("rust", "beam"))
//	// #bool true
func (s String) DoesntContain(subs ...string) bool {
	if len(subs) == 0 {
		return true
	}
	for _, sub := range subs {
		if sub == "" {
			continue
		}
		if strings.Contains(s.s, sub) {
			return false
		}
	}
	return true
}

// DoesntContainFold reports true if none of the substrings are found using Unicode case folding.
// Empty substrings are ignored.
// @group Search
//
// Example: doesn't contain (case-insensitive)
//
//	v := str.Of("gophers are great")
//	godump.Dump(v.DoesntContainFold("GOPHER"))
//	// #bool false
func (s String) DoesntContainFold(subs ...string) bool {
	if len(subs) == 0 {
		return true
	}
	base := strings.ToLower(s.s)
	for _, sub := range subs {
		if sub == "" {
			continue
		}
		if strings.Contains(base, strings.ToLower(sub)) {
			return false
		}
	}
	return true
}

// DoesntStartWith reports true if the string starts with none of the provided prefixes (case-sensitive).
// @group Search
//
// Example: doesn't start with any
//
//	v := str.Of("gopher")
//	godump.Dump(v.DoesntStartWith("go", "rust"))
//	// #bool false
func (s String) DoesntStartWith(prefixes ...string) bool {
	if len(prefixes) == 0 {
		return true
	}
	for _, p := range prefixes {
		if strings.HasPrefix(s.s, p) {
			return false
		}
	}
	return true
}

// DoesntStartWithFold reports true if the string starts with none of the provided prefixes using Unicode case folding.
// @group Search
//
// Example: doesn't start with (case-insensitive)
//
//	v := str.Of("gopher")
//	godump.Dump(v.DoesntStartWithFold("GO"))
//	// #bool false
func (s String) DoesntStartWithFold(prefixes ...string) bool {
	if len(prefixes) == 0 {
		return true
	}
	lower := strings.ToLower(s.s)
	for _, p := range prefixes {
		if strings.HasPrefix(lower, strings.ToLower(p)) {
			return false
		}
	}
	return true
}

// DoesntEndWith reports true if the string ends with none of the provided suffixes (case-sensitive).
// @group Search
//
// Example: doesn't end with any
//
//	v := str.Of("gopher")
//	godump.Dump(v.DoesntEndWith("her", "cat"))
//	// #bool false
func (s String) DoesntEndWith(suffixes ...string) bool {
	if len(suffixes) == 0 {
		return true
	}
	for _, suf := range suffixes {
		if strings.HasSuffix(s.s, suf) {
			return false
		}
	}
	return true
}

// DoesntEndWithFold reports true if the string ends with none of the provided suffixes using Unicode case folding.
// @group Search
//
// Example: doesn't end with (case-insensitive)
//
//	v := str.Of("gopher")
//	godump.Dump(v.DoesntEndWithFold("HER"))
//	// #bool false
func (s String) DoesntEndWithFold(suffixes ...string) bool {
	if len(suffixes) == 0 {
		return true
	}
	lower := strings.ToLower(s.s)
	for _, suf := range suffixes {
		if strings.HasSuffix(lower, strings.ToLower(suf)) {
			return false
		}
	}
	return true
}
