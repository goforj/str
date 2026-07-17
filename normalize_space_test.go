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
	if got := Of("GoForj builds practical Go applications").NormalizeSpace().String(); got != "GoForj builds practical Go applications" {
		t.Fatalf("NormalizeSpace clean = %q", got)
	}
	if got := Of("\u2003GoForj\u00a0\tbuilds\u2002Go\u2003").NormalizeSpace().String(); got != "GoForj builds Go" {
		t.Fatalf("NormalizeSpace Unicode whitespace = %q", got)
	}
	if got := Of("Go\xff  Forj").NormalizeSpace().String(); got != "Go\uFFFD Forj" {
		t.Fatalf("NormalizeSpace invalid UTF-8 = %q", got)
	}
}
