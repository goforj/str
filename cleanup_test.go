package str

import "testing"

func TestDeduplicateSquish(t *testing.T) {
	t.Parallel()

	if got := Of("The   Go   Playground").Deduplicate(' ').String(); got != "The Go Playground" {
		t.Fatalf("Deduplicate = %q", got)
	}
	if got := Of("a  b").Deduplicate(0).String(); got != "a b" {
		t.Fatalf("Deduplicate default char = %q", got)
	}
	if got := Of("   go   forj  ").Squish().String(); got != "go forj" {
		t.Fatalf("Squish = %q", got)
	}
}

func TestTrimAndBlank(t *testing.T) {
	t.Parallel()

	if got := Of("  GoForj  ").Trim("").String(); got != "GoForj" {
		t.Fatalf("Trim = %q", got)
	}
	if got := Of("  GoForj  ").TrimLeft("").String(); got != "GoForj  " {
		t.Fatalf("TrimLeft = %q", got)
	}
	if got := Of("  GoForj  ").TrimRight("").String(); got != "  GoForj" {
		t.Fatalf("TrimRight = %q", got)
	}
	if got := Of("xxgo").Trim("x").String(); got != "go" {
		t.Fatalf("Trim cutset")
	}
	if got := Of("xxgo").TrimLeft("x").String(); got != "go" {
		t.Fatalf("TrimLeft cutset")
	}
	if got := Of("goxx").TrimRight("x").String(); got != "go" {
		t.Fatalf("TrimRight cutset")
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
