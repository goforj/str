package str

import "testing"

func TestLimitAndTake(t *testing.T) {
	t.Parallel()

	val := Of("gophers")
	if got := val.Take(3).String(); got != "gop" {
		t.Fatalf("Take = %q", got)
	}
	if got := val.Take(10).String(); got != "gophers" {
		t.Fatalf("Take overrun returns original")
	}
	if got := val.TakeLast(4).String(); got != "hers" {
		t.Fatalf("TakeLast = %q", got)
	}
	if got := val.Limit(4, "...").String(); got != "goph..." {
		t.Fatalf("Limit = %q", got)
	}
	if got := val.Take(-1).String(); got != "" {
		t.Fatalf("Take negative")
	}
	if got := val.TakeLast(0).String(); got != "" {
		t.Fatalf("TakeLast zero")
	}
	if got := val.TakeLast(10).String(); got != "gophers" {
		t.Fatalf("TakeLast overrun")
	}
	if got := val.Limit(0, "...").String(); got != "..." {
		t.Fatalf("Limit zero length")
	}
	if got := val.Limit(10, "...").String(); got != "gophers" {
		t.Fatalf("Limit no truncation")
	}
	if got := val.Limit(-1, "").String(); got != "" {
		t.Fatalf("Limit negative no suffix")
	}
}

func TestCharAtAndReverse(t *testing.T) {
	t.Parallel()

	r, ok := Of("gopher").CharAt(2)
	if !ok || r != 'p' {
		t.Fatalf("CharAt unexpected")
	}
	if _, ok := Of("gopher").CharAt(-1); ok {
		t.Fatalf("CharAt negative should be false")
	}
	if _, ok := Of("go").CharAt(10); ok {
		t.Fatalf("CharAt out of range")
	}
	if got := Of("naïve").Reverse().String(); got != "evïan" {
		t.Fatalf("Reverse = %q", got)
	}
}
