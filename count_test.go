package str

import "testing"

// TestCount guards its covered contract against regressions.
func TestCount(t *testing.T) {
	t.Parallel()

	if got := Of("gogophergo").Count("go"); got != 3 {
		t.Fatalf("Count = %d", got)
	}
	if got := Of("abc").Count(""); got != 4 {
		t.Fatalf("Count empty expected 4, got %d", got)
	}
}
