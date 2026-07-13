package str

import "testing"

// TestReplaceFold verifies all non-overlapping Unicode simple-fold matches are replaced safely.
func TestReplaceFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value String
		old   string
		repl  string
		want  string
	}{
		{name: "ASCII all", value: Of("go gopher GO"), old: "GO", repl: "Go", want: "Go Gopher Go"},
		{name: "Greek sigma variants", value: Of("Σ σ ς"), old: "σ", repl: "s", want: "s s s"},
		{name: "different UTF-8 widths", value: Of("aKb kb"), old: "KB", repl: "X", want: "aX X"},
		{name: "non-overlapping", value: Of("ΣΣΣ"), old: "σς", repl: "x", want: "xΣ"},
		{name: "empty replacement", value: Of("Σxς"), old: "σ", repl: "", want: "x"},
		{name: "empty old", value: Of("go"), old: "", repl: "x", want: "go"},
		{name: "empty receiver", value: Of(""), old: "go", repl: "x", want: ""},
		{name: "missing", value: Of("go"), old: "rust", repl: "x", want: "go"},
		{name: "longer old", value: Of("go"), old: "gopher", repl: "x", want: "go"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := test.value.ReplaceFold(test.old, test.repl).String(); got != test.want {
				t.Fatalf("ReplaceFold(%q, %q) = %q, want %q", test.old, test.repl, got, test.want)
			}
		})
	}
}

// TestFoldMatchRange verifies shared matching reports byte offsets for unequal-width fold pairs.
func TestFoldMatchRange(t *testing.T) {
	t.Parallel()

	if got, ok := replaceFoldAll("go", "", "x"); ok || got != "go" {
		t.Fatalf("replaceFoldAll empty old = %q, %v", got, ok)
	}
	start, end, ok := foldMatchRange("xKΣy", "kς", 0)
	if !ok || start != 1 || end != 6 {
		t.Fatalf("foldMatchRange = (%d, %d, %v), want (1, 6, true)", start, end, ok)
	}
	if _, _, ok := foldMatchRange("go", "", 0); ok {
		t.Fatal("foldMatchRange matched an empty substring")
	}
}
