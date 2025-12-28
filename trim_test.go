package str

import "testing"

func TestTrim(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj  ").Trim("").String(); got != "GoForj" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("xxgo").Trim("x").String(); got != "go" {
		t.Fatalf("Trim cutset")
	}
}

func TestTrimLeft(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj  ").TrimLeft("").String(); got != "GoForj  " {
		t.Fatalf("TrimLeft = %q", got)
	}
	if got := Of("xxgo").TrimLeft("x").String(); got != "go" {
		t.Fatalf("TrimLeft cutset")
	}
}

func TestTrimRight(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj  ").TrimRight("").String(); got != "  GoForj" {
		t.Fatalf("TrimRight = %q", got)
	}
	if got := Of("goxx").TrimRight("x").String(); got != "go" {
		t.Fatalf("TrimRight cutset")
	}
}

func TestTrimSpace(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj  ").TrimSpace().String(); got != "GoForj" {
		t.Fatalf("TrimSpace = %q", got)
	}
	if got := Of("GoForj").TrimSpace().String(); got != "GoForj" {
		t.Fatalf("TrimSpace no-op = %q", got)
	}
	if got := Of("").TrimSpace().String(); got != "" {
		t.Fatalf("TrimSpace empty = %q", got)
	}
}
