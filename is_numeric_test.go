package str

import "testing"

// TestIsNumeric guards its covered contract against regressions.
func TestIsNumeric(t *testing.T) {
	t.Parallel()

	if !Of("12345").IsNumeric() {
		t.Fatalf("IsNumeric expected true")
	}
	if !Of("١٢٣").IsNumeric() {
		t.Fatalf("IsNumeric unicode expected true")
	}
	if Of("12.5").IsNumeric() {
		t.Fatalf("IsNumeric expected false")
	}
	if Of("").IsNumeric() {
		t.Fatalf("IsNumeric empty expected false")
	}
}
