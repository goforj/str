package str

import "testing"

// TestHasPrefix verifies singular, case-sensitive prefix matching and empty inputs.
func TestHasPrefix(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.HasPrefix("go") {
		t.Fatal("HasPrefix did not find prefix")
	}
	if val.HasPrefix("Go") {
		t.Fatal("HasPrefix ignored case")
	}
	if val.HasPrefix("") {
		t.Fatal("HasPrefix matched an empty prefix")
	}
	if Of("").HasPrefix("") {
		t.Fatal("HasPrefix matched an empty prefix in an empty receiver")
	}
}

// TestHasPrefixFold verifies Unicode simple folding at the leading byte boundary.
func TestHasPrefixFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		value  String
		prefix string
		want   bool
	}{
		{name: "ASCII", value: Of("gopher"), prefix: "GO", want: true},
		{name: "Greek sigma variants", value: Of("Σίσυφος"), prefix: "ς", want: true},
		{name: "different UTF-8 widths", value: Of("Kelvin"), prefix: "k", want: true},
		{name: "empty prefix", value: Of("gopher"), prefix: "", want: false},
		{name: "empty receiver", value: Of(""), prefix: "", want: false},
		{name: "missing", value: Of("gopher"), prefix: "CAT", want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := test.value.HasPrefixFold(test.prefix); got != test.want {
				t.Fatalf("HasPrefixFold(%q) = %v, want %v", test.prefix, got, test.want)
			}
		})
	}
}
