package str

import "testing"

// TestContains verifies singular, case-sensitive substring matching and the empty-substring rule.
func TestContains(t *testing.T) {
	t.Parallel()

	val := Of("Go means gophers")
	if !val.Contains("gopher") {
		t.Fatal("Contains did not find substring")
	}
	if val.Contains("") {
		t.Fatal("Contains matched an empty substring")
	}
	if val.Contains("rust") {
		t.Fatal("Contains found a missing substring")
	}
	if Of("").Contains("") {
		t.Fatal("Contains matched an empty substring in an empty receiver")
	}
}

// TestContainsFold verifies Unicode simple folding without relying on equal UTF-8 byte lengths.
func TestContainsFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value String
		sub   string
		want  bool
	}{
		{name: "ASCII", value: Of("Go means gophers"), sub: "GOPHER", want: true},
		{name: "Greek sigma variants", value: Of("AΣB"), sub: "ςb", want: true},
		{name: "different UTF-8 widths", value: Of("xKy"), sub: "kY", want: true},
		{name: "empty substring", value: Of("gopher"), sub: "", want: false},
		{name: "empty receiver", value: Of(""), sub: "go", want: false},
		{name: "missing", value: Of("gopher"), sub: "RUST", want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := test.value.ContainsFold(test.sub); got != test.want {
				t.Fatalf("ContainsFold(%q) = %v, want %v", test.sub, got, test.want)
			}
		})
	}
}
