package str

// EnsurePrefix ensures the string starts with prefix, adding it if missing.
// Similar: EnsureSuffix and TrimPrefix.
// @group Affixes
//
// Example: ensure prefix
//
//	v := str.Of("path/to").EnsurePrefix("/").String()
//	println(v)
//	// #string /path/to
func (s String) EnsurePrefix(prefix string) String {
	if prefix == "" || s.HasPrefix(prefix) {
		return s
	}
	return String{s: prefix + s.s}
}

// EnsureSuffix ensures the string ends with suffix, adding it if missing.
// Similar: EnsurePrefix and TrimSuffix.
// @group Affixes
//
// Example: ensure suffix
//
//	v := str.Of("path/to").EnsureSuffix("/").String()
//	println(v)
//	// #string path/to/
func (s String) EnsureSuffix(suffix string) String {
	if suffix == "" || s.HasSuffix(suffix) {
		return s
	}
	return String{s: s.s + suffix}
}
