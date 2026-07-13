package str

import "testing"

// TestStringWrapper guards its covered contract against regressions.
func TestStringWrapper(t *testing.T) {
	t.Parallel()

	s := Of("gopher")
	if s.String() != "gopher" {
		t.Fatalf("String mismatch")
	}
	if s.GoString() != "gopher" {
		t.Fatalf("GoString mismatch")
	}
}
