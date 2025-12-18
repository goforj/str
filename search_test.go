package str

import "testing"

func TestEqualsAndStartsEnds(t *testing.T) {
	t.Parallel()

	val := Of("gopher")
	if !val.Equals("gopher") || val.Equals("Gopher") {
		t.Fatalf("Equals mismatch")
	}
	if !val.EqualsFold("GOPHER") {
		t.Fatalf("EqualsFold mismatch")
	}
	if !val.StartsWith("go") || val.StartsWith("rust") {
		t.Fatalf("StartsWith mismatch")
	}
	if val.StartsWith() {
		t.Fatalf("StartsWith empty slice should be false")
	}
	if !val.StartsWithFold("GO") {
		t.Fatalf("StartsWithFold mismatch")
	}
	if !val.EndsWith("her") || val.EndsWith("cat") {
		t.Fatalf("EndsWith mismatch")
	}
	if val.EndsWith() {
		t.Fatalf("EndsWith empty slice should be false")
	}
	if !val.EndsWithFold("HER") {
		t.Fatalf("EndsWithFold mismatch")
	}
	if val.EndsWithFold() {
		t.Fatalf("EndsWithFold empty slice should be false")
	}
	if val.EndsWithFold("CAT") {
		t.Fatalf("EndsWithFold non-match should be false")
	}
	if val.StartsWithFold() {
		t.Fatalf("StartsWithFold empty slice should be false")
	}
	if val.StartsWithFold("CAT") {
		t.Fatalf("StartsWithFold non-match should be false")
	}
}

func TestIndexing(t *testing.T) {
	t.Parallel()

	val := Of("héllo gopher")
	if got := val.Index("llo"); got != 2 {
		t.Fatalf("Index = %d", got)
	}
	if got := val.LastIndex("o"); got != 7 {
		t.Fatalf("LastIndex = %d", got)
	}
	if got := val.Slice(2, 6).String(); got != "llo " {
		t.Fatalf("Slice = %q", got)
	}
	if got := val.Slice(20, 30).String(); got != "" {
		t.Fatalf("Slice clamp beyond length %q", got)
	}
	if got := val.Slice(5, 3).String(); got != "" {
		t.Fatalf("Slice start>=end %q", got)
	}
}

func TestBetweenVariants(t *testing.T) {
	t.Parallel()

	if got := Of("This is my name").Between("This", "name").String(); got != " is my " {
		t.Fatalf("Between = %q", got)
	}
	if got := Of("[a] bc [d]").BetweenFirst("[", "]").String(); got != "a" {
		t.Fatalf("BetweenFirst = %q", got)
	}
}
