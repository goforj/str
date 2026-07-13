package str

import "testing"

// TestIsEmptyBlank guards its covered contract against regressions.
func TestIsEmptyBlank(t *testing.T) {
	t.Parallel()

	if !Of("").IsEmpty() {
		t.Fatalf("IsEmpty expected true")
	}
	if !Of(" \t\n").IsBlank() {
		t.Fatalf("IsBlank expected true")
	}
	if Of("go").IsBlank() {
		t.Fatalf("IsBlank expected false")
	}
}
