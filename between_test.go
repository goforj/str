package str

import "testing"

func TestBetween(t *testing.T) {
	t.Parallel()

	if got := Of("This is my name").Between("This", "name").String(); got != " is my " {
		t.Fatalf("Between = %q", got)
	}
	if got := Of("abc").Between("a", "z").String(); got != "" {
		t.Fatalf("Between missing end")
	}
	if got := Of("abc").Between("bc", "a").String(); got != "" {
		t.Fatalf("Between overlapping order")
	}
	if got := Of("abc").Between("", "c").String(); got != "" {
		t.Fatalf("Between empty start should be empty")
	}
}

func TestBetweenFirst(t *testing.T) {
	t.Parallel()

	if got := Of("[a] bc [d]").BetweenFirst("[", "]").String(); got != "a" {
		t.Fatalf("BetweenFirst = %q", got)
	}
	if got := Of("abc").BetweenFirst("a", "x").String(); got != "" {
		t.Fatalf("BetweenFirst missing end")
	}
	if got := Of("abc").BetweenFirst("x", "c").String(); got != "" {
		t.Fatalf("BetweenFirst missing start")
	}
	if got := Of("abc").BetweenFirst("", "c").String(); got != "" {
		t.Fatalf("BetweenFirst empty start")
	}
}
