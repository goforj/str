package str

import "testing"

// TestReplaceAll guards its covered contract against regressions.
func TestReplaceAll(t *testing.T) {
	t.Parallel()

	if got := Of("go gopher go").ReplaceAll("go", "Go").String(); got != "Go Gopher Go" {
		t.Fatalf("expected Go Gopher Go, got %s", got)
	}

	if got := Of("gopher").ReplaceAll("rust", "go").String(); got != "gopher" {
		t.Fatalf("expected unchanged when no matches, got %s", got)
	}

	if got := Of("gopher").ReplaceAll("", "x").String(); got != "gopher" {
		t.Fatalf("expected unchanged when old is empty, got %s", got)
	}
}
