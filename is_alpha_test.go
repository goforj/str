package str

import "testing"

// TestIsAlpha guards its covered contract against regressions.
func TestIsAlpha(t *testing.T) {
	t.Parallel()

	if !Of("Gopher").IsAlpha() {
		t.Fatalf("IsAlpha expected true")
	}
	if !Of("Göpher").IsAlpha() {
		t.Fatalf("IsAlpha unicode expected true")
	}
	if Of("Gopher 2").IsAlpha() {
		t.Fatalf("IsAlpha expected false")
	}
	if Of("").IsAlpha() {
		t.Fatalf("IsAlpha empty expected false")
	}
}
