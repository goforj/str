package str

import (
	"strings"
	"testing"
)

// TestApplicationEmptySearch keeps folded predicates and replacement helpers consistent at boundaries.
func TestApplicationEmptySearch(t *testing.T) {
	t.Parallel()
	for _, input := range []string{"", "abc", "é🦫", "\xffa"} {
		value := Of(input)
		if !value.ContainsFold("") || !value.HasPrefixFold("") || !value.HasSuffixFold("") {
			t.Fatalf("empty folded search did not match %q", input)
		}
		for _, replacement := range []string{"", "-", "é"} {
			all := strings.ReplaceAll(input, "", replacement)
			checks := map[string]struct{ got, want string }{
				"first":  {value.ReplaceFirst("", replacement).String(), replacement + input},
				"last":   {value.ReplaceLast("", replacement).String(), input + replacement},
				"prefix": {value.ReplacePrefix("", replacement).String(), replacement + input},
				"suffix": {value.ReplaceSuffix("", replacement).String(), input + replacement},
				"fold":   {value.ReplaceFold("", replacement).String(), all},
				"array":  {value.ReplaceArray([]string{"", "a"}, replacement).String(), strings.ReplaceAll(all, "a", replacement)},
				"swap":   {value.Swap(map[string]string{"": replacement}).String(), strings.NewReplacer("", replacement).Replace(input)},
				"mixed":  {value.Swap(map[string]string{"a": "A", "": replacement}).String(), strings.NewReplacer("a", "A", "", replacement).Replace(input)},
			}
			for name, check := range checks {
				if check.got != check.want {
					t.Fatalf("%s(%q, %q) = %q, want %q", name, input, replacement, check.got, check.want)
				}
			}
		}
	}
}

// TestFoldCandidateProgress protects internal scans from zero-width matches while public methods handle empty searches.
func TestFoldCandidateProgress(t *testing.T) {
	t.Parallel()
	if _, ok := foldMatchAt("go", "", 0); ok {
		t.Fatal("empty candidate must not enter a scanning loop")
	}
	if _, ok := foldSuffixStart("go", ""); ok {
		t.Fatal("empty suffix must be handled at the public boundary")
	}
}
