package str

import "testing"

func TestDeduplicateSquish(t *testing.T) {
	t.Parallel()

	if got := Of("The   Laravel   Framework").Deduplicate(' ').String(); got != "The Laravel Framework" {
		t.Fatalf("Deduplicate = %q", got)
	}
	if got := Of("   go   forj  ").Squish().String(); got != "go forj" {
		t.Fatalf("Squish = %q", got)
	}
}

func TestTrimAndBlank(t *testing.T) {
	t.Parallel()

	if got := Of("  Laravel  ").Trim("").String(); got != "Laravel" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("  Laravel  ").TrimLeft("").String(); got != "Laravel  " {
		t.Fatalf("TrimLeft = %q", got)
	}
	if got := Of("  Laravel  ").TrimRight("").String(); got != "  Laravel" {
		t.Fatalf("TrimRight = %q", got)
	}
	if !Of("").IsEmpty() {
		t.Fatalf("IsEmpty expected true")
	}
	if !Of(" \t\n").IsBlank() {
		t.Fatalf("IsBlank expected true")
	}
	if Of("go").IsBlank() {
		t.Fatalf("IsBlank expected false")
	}
}
