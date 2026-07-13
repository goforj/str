package str

import "testing"

// TestDeduplicate guards its covered contract against regressions.
func TestDeduplicate(t *testing.T) {
	t.Parallel()

	if got := Of("The   Go   Playground").Deduplicate(' ').String(); got != "The Go Playground" {
		t.Fatalf("Deduplicate = %q", got)
	}
	if got := Of("a  b").Deduplicate(0).String(); got != "a b" {
		t.Fatalf("Deduplicate default char = %q", got)
	}
}
