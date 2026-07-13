package str

import "testing"

// TestToUpper guards its covered contract against regressions.
func TestToUpper(t *testing.T) {
	t.Parallel()

	if got := Of("GoLang").ToUpper().String(); got != "GOLANG" {
		t.Fatalf("ToUpper = %q", got)
	}
}
