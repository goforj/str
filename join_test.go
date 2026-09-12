package str

import "testing"

// TestJoin guards its covered contract against regressions.
func TestJoin(t *testing.T) {
	t.Parallel()

	if got := Join([]string{"foo", "bar"}, "-").String(); got != "foo-bar" {
		t.Fatalf("Join = %q", got)
	}
	if got := Join(nil, "-").String(); got != "" {
		t.Fatalf("Join nil = %q", got)
	}
}
