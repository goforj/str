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
	if !val.StartsWithFold("GO") {
		t.Fatalf("StartsWithFold mismatch")
	}
	if !val.EndsWith("her") || val.EndsWith("cat") {
		t.Fatalf("EndsWith mismatch")
	}
	if !val.EndsWithFold("HER") {
		t.Fatalf("EndsWithFold mismatch")
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
