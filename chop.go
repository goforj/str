package str

import "strings"

// TrimPrefix removes prefix when it appears at the start of the string.
// Similar: TrimSuffix and EnsurePrefix.
// @group Affixes
//
// Example: trim prefix
//
//	v := str.Of("https://goforj.dev").TrimPrefix("https://").String()
//	println(v)
//	// #string goforj.dev
func (s String) TrimPrefix(prefix string) String {
	if prefix == "" {
		return s
	}
	return String{s: strings.TrimPrefix(s.s, prefix)}
}

// TrimSuffix removes suffix when it appears at the end of the string.
// Similar: TrimPrefix and EnsureSuffix.
// @group Affixes
//
// Example: trim suffix
//
//	v := str.Of("file.txt").TrimSuffix(".txt").String()
//	println(v)
//	// #string file
func (s String) TrimSuffix(suffix string) String {
	if suffix == "" {
		return s
	}
	return String{s: strings.TrimSuffix(s.s, suffix)}
}
