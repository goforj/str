package str

import "testing"

// TestSubstrReplace guards its covered contract against regressions.
func TestSubstrReplace(t *testing.T) {
	t.Parallel()

	val := Of("naïve café")
	if got := val.SubstrReplace("i", 2, 3).String(); got != "naive café" {
		t.Fatalf("SubstrReplace = %q", got)
	}
	if got := val.SubstrReplace("X", -5, 1).String(); got != "Xaïve café" {
		t.Fatalf("SubstrReplace clamp start = %q", got)
	}
	if got := val.SubstrReplace("X", 20, 30).String(); got != "naïve caféX" {
		t.Fatalf("SubstrReplace clamp end = %q", got)
	}
	if got := val.SubstrReplace("X", 4, 2).String(); got != "naïve café" {
		t.Fatalf("SubstrReplace swapped range = %q", got)
	}
}
