package str

import "testing"

// TestTitle guards its covered contract against regressions.
func TestTitle(t *testing.T) {
	t.Parallel()

	if got := Of("a nice title uses the correct case").Title().String(); got != "A Nice Title Uses The Correct Case" {
		t.Fatalf("Title = %q", got)
	}
}
