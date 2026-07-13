package str

import "testing"

// TestTake guards its covered contract against regressions.
func TestTake(t *testing.T) {
	t.Parallel()

	val := Of("gophers")
	if got := val.Take(3).String(); got != "gop" {
		t.Fatalf("Take = %q", got)
	}
	if got := val.Take(10).String(); got != "gophers" {
		t.Fatalf("Take overrun returns original")
	}
	if got := val.Take(-1).String(); got != "" {
		t.Fatalf("Take negative")
	}
}

// TestTakeLast guards its covered contract against regressions.
func TestTakeLast(t *testing.T) {
	t.Parallel()

	val := Of("gophers")
	if got := val.TakeLast(4).String(); got != "hers" {
		t.Fatalf("TakeLast = %q", got)
	}
	if got := val.TakeLast(0).String(); got != "" {
		t.Fatalf("TakeLast zero")
	}
	if got := val.TakeLast(10).String(); got != "gophers" {
		t.Fatalf("TakeLast overrun")
	}
}
