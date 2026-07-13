package str

import "testing"

// TestCharAt guards its covered contract against regressions.
func TestCharAt(t *testing.T) {
	t.Parallel()

	r, ok := Of("gopher").CharAt(2)
	if !ok || r != 'p' {
		t.Fatalf("CharAt unexpected")
	}
	if _, ok := Of("gopher").CharAt(-1); ok {
		t.Fatalf("CharAt negative should be false")
	}
	if _, ok := Of("go").CharAt(10); ok {
		t.Fatalf("CharAt out of range")
	}
}
