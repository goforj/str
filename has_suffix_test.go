package str

import "testing"

// TestHasSuffix verifies singular, case-sensitive suffix matching and empty inputs.
func TestHasSuffix(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.HasSuffix("her") {
		t.Fatal("HasSuffix did not find suffix")
	}
	if val.HasSuffix("HER") {
		t.Fatal("HasSuffix ignored case")
	}
	if val.HasSuffix("") {
		t.Fatal("HasSuffix matched an empty suffix")
	}
	if Of("").HasSuffix("") {
		t.Fatal("HasSuffix matched an empty suffix in an empty receiver")
	}
}

// TestHasSuffixFold verifies Unicode simple folding at the trailing byte boundary.
func TestHasSuffixFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		value  String
		suffix string
		want   bool
	}{
		{name: "ASCII", value: Of("gopher"), suffix: "HER", want: true},
		{name: "Greek sigma variants", value: Of("final Σ"), suffix: "ς", want: true},
		{name: "different UTF-8 widths", value: Of("unit K"), suffix: "k", want: true},
		{name: "empty suffix", value: Of("gopher"), suffix: "", want: false},
		{name: "empty receiver", value: Of(""), suffix: "", want: false},
		{name: "too long", value: Of("K"), suffix: "xxk", want: false},
		{name: "missing", value: Of("gopher"), suffix: "CAT", want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := test.value.HasSuffixFold(test.suffix); got != test.want {
				t.Fatalf("HasSuffixFold(%q) = %v, want %v", test.suffix, got, test.want)
			}
		})
	}
}
