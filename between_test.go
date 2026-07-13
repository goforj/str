package str

import "testing"

// TestBetween guards its covered contract against regressions.
func TestBetween(t *testing.T) {
	t.Parallel()

	if got := Of("This is my name").Between("This", "name").String(); got != " is my " {
		t.Fatalf("Between = %q", got)
	}
	if got := Of("[first] and [second]").Between("[", "]").String(); got != "first" {
		t.Fatalf("Between should use the nearest closing marker = %q", got)
	}
	if got := Of("abc").Between("a", "z").String(); got != "" {
		t.Fatalf("Between missing end")
	}
	if got := Of("abc").Between("z", "a").String(); got != "" {
		t.Fatalf("Between missing start")
	}
	if got := Of("abc").Between("bc", "a").String(); got != "" {
		t.Fatalf("Between overlapping order")
	}
	if got := Of("abc").Between("", "c").String(); got != "" {
		t.Fatalf("Between empty start should be empty")
	}
	if got := Of("abc").Between("a", "").String(); got != "" {
		t.Fatalf("Between empty end should be empty")
	}
}
