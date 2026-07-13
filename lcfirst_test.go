package str

import "testing"

// TestLcFirst guards its covered contract against regressions.
func TestLcFirst(t *testing.T) {
	t.Parallel()

	if got := Of("Gopher").LcFirst().String(); got != "gopher" {
		t.Fatalf("LcFirst = %q", got)
	}
	if got := Of("").LcFirst().String(); got != "" {
		t.Fatalf("LcFirst empty = %q", got)
	}
}
