package str

import "testing"

// TestCutLastBoundaries checks the backport independently of the pre-Go-1.27 reference implementation.
func TestCutLastBoundaries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, input, sep, before, after string
		found                           bool
	}{
		{"last match", "one:two:three", ":", "one:two", "three", true},
		{"overlapping matches", "ababa", "aba", "ab", "", true},
		{"overlapping run", "aaaaa", "aaa", "aa", "", true},
		{"whole input", "abc", "abc", "", "", true},
		{"prefix", ":abc", ":", "", "abc", true},
		{"suffix", "abc:", ":", "abc", "", true},
		{"missing", "abc", "x", "abc", "", false},
		{"long separator", "abc", "abcd", "abc", "", false},
		{"empty source", "", "x", "", "", false},
		{"both empty", "", "", "", "", true},
		{"empty separator", "é🦫", "", "é🦫", "", true},
		{"multibyte match", "a🦫b🦫c", "🦫", "a🦫b", "c", true},
		{"within encoded rune", "éé", "\xa9", "é\xc3", "", true},
		{"invalid bytes", "\xffa\xffb", "\xff", "\xffa", "b", true},
		{"NUL separator", "a\x00b\x00c", "\x00", "a\x00b", "c", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, after, found := Of(test.input).CutLast(test.sep)
			if before != test.before || after != test.after || found != test.found {
				t.Fatalf("CutLast(%q, %q) = (%q, %q, %t), want (%q, %q, %t)", test.input, test.sep, before, after, found, test.before, test.after, test.found)
			}
		})
	}
}
