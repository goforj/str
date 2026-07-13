package str

import "testing"

// TestLimit guards its covered contract against regressions.
func TestLimit(t *testing.T) {
	t.Parallel()

	val := Of("gophers")
	if got := val.Limit(4, "...").String(); got != "goph..." {
		t.Fatalf("Limit = %q", got)
	}
	if got := val.Limit(0, "...").String(); got != "..." {
		t.Fatalf("Limit zero length")
	}
	if got := val.Limit(10, "...").String(); got != "gophers" {
		t.Fatalf("Limit no truncation")
	}
	if got := val.Limit(-1, "").String(); got != "" {
		t.Fatalf("Limit negative no suffix")
	}
}
