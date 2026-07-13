package str

import "testing"

// TestReverse guards its covered contract against regressions.
func TestReverse(t *testing.T) {
	t.Parallel()

	if got := Of("naïve").Reverse().String(); got != "evïan" {
		t.Fatalf("Reverse = %q", got)
	}
}
