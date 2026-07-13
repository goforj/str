package str

import "testing"

// TestIsASCII guards its covered contract against regressions.
func TestIsASCII(t *testing.T) {
	t.Parallel()

	if !Of("gopher").IsASCII() {
		t.Fatalf("IsASCII should be true")
	}
	if Of("gophers 🦫").IsASCII() {
		t.Fatalf("IsASCII should be false with emoji")
	}
}
