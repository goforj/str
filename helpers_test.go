package str

import (
	"reflect"
	"testing"
)

// TestClampAndRuneHelpers guards its covered contract against regressions.
func TestClampAndRuneHelpers(t *testing.T) {
	t.Parallel()

	if a, b := clampRange(-1, 10, 5); a != 0 || b != 5 {
		t.Fatalf("clampRange negative start got %d %d", a, b)
	}
	if a, b := clampRange(10, 1, 5); a != 5 || b != 1 {
		t.Fatalf("clampRange reversed got %d %d", a, b)
	}
	if a, b := clampRange(10, 20, 5); a != 5 || b != 5 {
		t.Fatalf("clampRange beyond length %d %d", a, b)
	}
	if a, b := clampRange(-5, -1, 5); a != 0 || b != 0 {
		t.Fatalf("clampRange both negative %d %d", a, b)
	}

	if got := runeSubstring("naïve", 1, 3); got != "aï" {
		t.Fatalf("runeSubstring segment %q", got)
	}
	if got := runeSubstring("go", 5, 1); got != "" {
		t.Fatalf("runeSubstring start>=end")
	}

	if got := runeIndex("gopher", "ph", false); got != 2 {
		t.Fatalf("runeIndex first %d", got)
	}
	if got := runeIndex("gopher gopher", "go", true); got != 7 {
		t.Fatalf("runeIndex last %d", got)
	}
	if got := runeIndex("go", "", false); got != -1 {
		t.Fatalf("runeIndex empty sub")
	}
	if got := runeIndex("go", "z", true); got != -1 {
		t.Fatalf("runeIndex missing")
	}
}

// TestTokenizeWords guards its covered contract against regressions.
func TestTokenizeWords(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "acronyms", input: "HTTPRequestID", want: []string{"HTTP", "Request", "ID"}},
		{name: "mixed delimiters", input: "Go-Forj_Rules", want: []string{"Go", "Forj", "Rules"}},
		{name: "Unicode", input: "ÉclairHTTPÜberID", want: []string{"Éclair", "HTTP", "Über", "ID"}},
		{name: "combining mark", input: "Cafe\u0301HTTP", want: []string{"Cafe\u0301", "HTTP"}},
		{name: "empty", input: "", want: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := wordTokenValues(tokenizeWords(test.input))
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("tokenizeWords(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}

	tokens := tokenizeWords("  HTTPRequestID!")
	if tokens[0].start != 2 || tokens[1].end != 13 {
		t.Fatalf("token spans = %+v", tokens)
	}
}
