package str

import "testing"

// TestWrapAndUnwrap guards its covered contract against regressions.
func TestWrapAndUnwrap(t *testing.T) {
	t.Parallel()

	if got := Of("GoForj").Wrap(`"`, `"`).String(); got != `"GoForj"` {
		t.Fatalf("Wrap = %q", got)
	}
	if got := Of("GoForj").Wrap("<", "").String(); got != "<GoForj" {
		t.Fatalf("Wrap should use affixes literally = %q", got)
	}
	if got := Of(`"GoForj"`).Unwrap(`"`, `"`).String(); got != "GoForj" {
		t.Fatalf("Unwrap = %q", got)
	}
	if got := Of("gopher").Unwrap("[", "]").String(); got != "gopher" {
		t.Fatalf("Unwrap no match")
	}
	if got := Of("##go##").Unwrap("#", "").String(); got != "#go##" {
		t.Fatalf("Unwrap should use affixes literally = %q", got)
	}
}
