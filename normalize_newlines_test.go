package str

import "testing"

// TestNormalizeNewlines guards its covered contract against regressions.
func TestNormalizeNewlines(t *testing.T) {
	t.Parallel()

	in := "a\r\nb\rc\u0085d\u2028e\u2029f\n"
	want := "a\nb\nc\nd\ne\nf\n"
	if got := Of(in).NormalizeNewlines().String(); got != want {
		t.Fatalf("NormalizeNewlines = %q", got)
	}
	if got := Of("a\nb").NormalizeNewlines().String(); got != "a\nb" {
		t.Fatalf("NormalizeNewlines no change = %q", got)
	}
}
