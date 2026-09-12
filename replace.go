package str

import (
	"sort"
	"strings"
)

// ReplaceFirst replaces the first occurrence of old with repl.
// An empty old inserts repl at the beginning.
// Similar: ReplaceLast and ReplaceAll.
// @group Replace
//
// Example: replace first
//
//	v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
//	println(v)
//	// #string go gopher
func (s String) ReplaceFirst(old, repl string) String {
	return String{s: strings.Replace(s.s, old, repl, 1)}
}

// ReplaceLast replaces the last occurrence of old with repl.
// An empty old inserts repl at the end.
// Similar: ReplaceFirst and ReplaceAll.
// @group Replace
//
// Example: replace last
//
//	v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
//	println(v)
//	// #string gopher go
func (s String) ReplaceLast(old, repl string) String {
	idx := strings.LastIndex(s.s, old)
	if idx == -1 {
		return s
	}
	var b strings.Builder
	b.Grow(len(s.s) - len(old) + len(repl))
	b.WriteString(s.s[:idx])
	b.WriteString(repl)
	b.WriteString(s.s[idx+len(old):])
	return String{s: b.String()}
}

// ReplaceArray replaces all occurrences of each old in olds with repl.
// Entries are applied sequentially, including replacements produced by earlier entries.
// Empty entries insert repl at UTF-8 boundaries, like ReplaceAll.
// Similar: ReplaceAll and Swap.
// @group Replace
//
// Example: replace many
//
//	v := str.Of("The---Go---Toolkit")
//	println(v.ReplaceArray([]string{"---"}, "-").String())
//	// #string The-Go-Toolkit
func (s String) ReplaceArray(olds []string, repl string) String {
	out := s.s
	for _, old := range olds {
		out = strings.ReplaceAll(out, old, repl)
	}
	return String{s: out}
}

// Swap replaces multiple values in one pass using strings.Replacer built from a map.
// Longer keys take priority at the same position; replacements are not rescanned.
// Empty keys insert values at the boundaries handled by strings.Replacer.
// Similar: ReplaceArray.
// @group Replace
//
// Example: swap map
//
//	pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
//	v := str.Of("Gophers are great!").Swap(pairs).String()
//	println(v)
//	// #string GoForj is fantastic!
func (s String) Swap(pairs map[string]string) String {
	if len(pairs) == 0 {
		return s
	}
	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

	repPairs := make([]string, 0, len(keys)*2)
	for _, k := range keys {
		repPairs = append(repPairs, k, pairs[k])
	}

	r := strings.NewReplacer(repPairs...)
	return String{s: r.Replace(s.s)}
}
