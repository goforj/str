package str

import "testing"

// TestLines guards its covered contract against regressions.
func TestLines(t *testing.T) {
	t.Parallel()

	got := Of("a\r\nb\nc").Lines()
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Fatalf("Lines = %#v", got)
	}
}
