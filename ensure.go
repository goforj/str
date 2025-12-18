package str

// EnsurePrefix ensures the string starts with prefix, adding it if missing.
// @group Affixes
//
// Example: ensure prefix
//
//	val := str.Of("path/to")
//	godump.Dump(val.EnsurePrefix("/").String())
//
//	// #string /path/to
func (s String) EnsurePrefix(prefix string) String {
	if prefix == "" || s.StartsWith(prefix) {
		return s
	}
	return String{s: prefix + s.s}
}

// EnsureSuffix ensures the string ends with suffix, adding it if missing.
// @group Affixes
//
// Example: ensure suffix
//
//	val := str.Of("path/to")
//	godump.Dump(val.EnsureSuffix("/").String())
//
//	// #string path/to/
func (s String) EnsureSuffix(suffix string) String {
	if suffix == "" || s.EndsWith(suffix) {
		return s
	}
	return String{s: s.s + suffix}
}
