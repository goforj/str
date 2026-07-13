package str

import "testing"

// TestTrim guards its covered contract against regressions.
func TestTrim(t *testing.T) {
	t.Parallel()

	if got := Of("\u2003GoForj\u00a0").Trim().String(); got != "GoForj" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("..GoForj!!").TrimChars(".!").String(); got != "GoForj" {
		t.Fatalf("TrimChars = %q", got)
	}
	if got := Of("GoForj").TrimChars("").String(); got != "GoForj" {
		t.Fatalf("TrimChars with an empty cutset = %q", got)
	}
}

// TestTrimLeft guards its covered contract against regressions.
func TestTrimLeft(t *testing.T) {
	t.Parallel()

	if got := Of("\u2003GoForj  ").TrimLeft().String(); got != "GoForj  " {
		t.Fatalf("TrimLeft = %q", got)
	}
}

// TestTrimRight guards its covered contract against regressions.
func TestTrimRight(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj\u00a0").TrimRight().String(); got != "  GoForj" {
		t.Fatalf("TrimRight = %q", got)
	}
}
