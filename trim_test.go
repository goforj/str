package str

import (
	"testing"
	"unicode"
)

// TestTrim guards its covered contract against regressions.
func TestTrim(t *testing.T) {
	t.Parallel()

	if got := Of("\u2003GoForj\u00a0").TrimSpace().String(); got != "GoForj" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("..GoForj!!").Trim(".!").String(); got != "GoForj" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("GoForj").Trim("").String(); got != "GoForj" {
		t.Fatalf("Trim with an empty cutset = %q", got)
	}
}

// TestTrimLeft guards its covered contract against regressions.
func TestTrimLeft(t *testing.T) {
	t.Parallel()

	if got := Of("\u2003GoForj  ").TrimLeftFunc(unicode.IsSpace).String(); got != "GoForj  " {
		t.Fatalf("TrimLeft = %q", got)
	}
}

// TestTrimRight guards its covered contract against regressions.
func TestTrimRight(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj\u00a0").TrimRightFunc(unicode.IsSpace).String(); got != "  GoForj" {
		t.Fatalf("TrimRight = %q", got)
	}
}
