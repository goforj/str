package str

import "testing"

// TestUcFirst guards its covered contract against regressions.
func TestUcFirst(t *testing.T) {
	t.Parallel()

	if got := Of("gopher").UcFirst().String(); got != "Gopher" {
		t.Fatalf("UcFirst = %q", got)
	}
	if got := Of("Go").UcFirst().String(); got != "Go" {
		t.Fatalf("UcFirst already upper = %q", got)
	}
	if got := Of("").UcFirst().String(); got != "" {
		t.Fatalf("UcFirst empty = %q", got)
	}
}
