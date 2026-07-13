package str

import "testing"

// TestEqualFold verifies standard Unicode simple-fold equality, including empty strings.
func TestEqualFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value String
		other string
		want  bool
	}{
		{name: "ASCII", value: Of("gopher"), other: "GOPHER", want: true},
		{name: "Greek sigma variants", value: Of("Σσς"), other: "σςΣ", want: true},
		{name: "different UTF-8 widths", value: Of("K"), other: "k", want: true},
		{name: "empty", value: Of(""), other: "", want: true},
		{name: "different", value: Of("gopher"), other: "rust", want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := test.value.EqualFold(test.other); got != test.want {
				t.Fatalf("EqualFold(%q) = %v, want %v", test.other, got, test.want)
			}
		})
	}
}
