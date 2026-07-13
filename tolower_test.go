package str

import "testing"

// TestToLower guards its covered contract against regressions.
func TestToLower(t *testing.T) {
	t.Parallel()

	if got := Of("GoLang").ToLower().String(); got != "golang" {
		t.Fatalf("ToLower = %q", got)
	}
}
