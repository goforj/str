package str

import "testing"

// TestCommonPrefix guards its covered contract against regressions.
func TestCommonPrefix(t *testing.T) {
	t.Parallel()

	if got := Of("gopher").CommonPrefix("go", "gold").String(); got != "go" {
		t.Fatalf("CommonPrefix = %q", got)
	}
	if got := Of("naïve").CommonPrefix("naïveté").String(); got != "naïve" {
		t.Fatalf("CommonPrefix unicode = %q", got)
	}
	if got := Of("gopher").CommonPrefix().String(); got != "gopher" {
		t.Fatalf("CommonPrefix no args = %q", got)
	}
}
