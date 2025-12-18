package str

import "strings"

// Contains reports whether the string contains any of the provided substrings (case-sensitive).
// Empty substrings return true to match strings.Contains semantics.
// @group Search
//
// Example: contains any
//
//	v := str.Of("Go means gophers").Contains("rust", "gopher")
//	godump.Dump(v)
//	// #bool true
func (s String) Contains(subs ...string) bool {
	if len(subs) == 0 {
		return false
	}

	for _, sub := range subs {
		if sub == "" || strings.Contains(s.s, sub) {
			return true
		}
	}
	return false
}

// ContainsAll reports whether the string contains all provided substrings (case-sensitive).
// Empty substrings are ignored.
// @group Search
//
// Example: contains all
//
//	v := str.Of("Go means gophers").ContainsAll("Go", "gopher")
//	godump.Dump(v)
//	// #bool true
func (s String) ContainsAll(subs ...string) bool {
	if len(subs) == 0 {
		return false
	}

	for _, sub := range subs {
		if sub == "" {
			continue
		}
		if !strings.Contains(s.s, sub) {
			return false
		}
	}
	return true
}

// ContainsFold reports whether the string contains any of the provided substrings, using Unicode
// case folding for comparisons.
// Empty substrings return true.
// @group Search
//
// Example: contains any (case-insensitive)
//
//	v := str.Of("Go means gophers").ContainsFold("GOPHER", "rust")
//	godump.Dump(v)
//	// #bool true
func (s String) ContainsFold(subs ...string) bool {
	if len(subs) == 0 {
		return false
	}

	fold := strings.ToLower(s.s)

	for _, sub := range subs {
		if sub == "" || strings.Contains(fold, strings.ToLower(sub)) {
			return true
		}
	}
	return false
}

// ContainsAllFold reports whether the string contains all provided substrings, using Unicode
// case folding for comparisons.
// Empty substrings are ignored.
// @group Search
//
// Example: contains all (case-insensitive)
//
//	v := str.Of("Go means gophers").ContainsAllFold("go", "GOPHER")
//	godump.Dump(v)
//	// #bool true
func (s String) ContainsAllFold(subs ...string) bool {
	if len(subs) == 0 {
		return false
	}

	fold := strings.ToLower(s.s)

	for _, sub := range subs {
		if sub == "" {
			continue
		}
		if !strings.Contains(fold, strings.ToLower(sub)) {
			return false
		}
	}
	return true
}
