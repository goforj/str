package str

import "testing"

// TestNormalizeSpace guards its covered contract against regressions.
func TestNormalizeSpace(t *testing.T) {
	t.Parallel()

	if got := Of("  go\t forj  ").NormalizeSpace().String(); got != "go forj" {
		t.Fatalf("NormalizeSpace = %q", got)
	}
	if got := Of("   ").NormalizeSpace().String(); got != "" {
		t.Fatalf("NormalizeSpace all-space = %q", got)
	}
	if got := Of("").NormalizeSpace().String(); got != "" {
		t.Fatalf("NormalizeSpace empty = %q", got)
	}
}
