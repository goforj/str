package str

import "regexp"

// IsMatch reports whether the string matches the provided regular expression.
// @group Match
//
// Example: regex match
//
//	v := str.Of("abc123").IsMatch(regexp.MustCompile(`\d+`))
//	godump.Dump(v)
//	// #bool true
func (s String) IsMatch(re *regexp.Regexp) bool {
	if re == nil {
		return false
	}
	return re.MatchString(s.s)
}
