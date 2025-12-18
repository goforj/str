package str

import "testing"

func TestBeforeVariants(t *testing.T) {
	t.Parallel()

	if got := Of("gopher::go").Before("::").String(); got != "gopher" {
		t.Fatalf("Before = %q", got)
	}
	if got := Of("pkg/path/file.txt").BeforeLast("/").String(); got != "pkg/path" {
		t.Fatalf("BeforeLast = %q", got)
	}
}

func TestLimitAndTake(t *testing.T) {
	t.Parallel()

	val := Of("gophers")
	if got := val.Take(3).String(); got != "gop" {
		t.Fatalf("Take = %q", got)
	}
	if got := val.TakeLast(4).String(); got != "hers" {
		t.Fatalf("TakeLast = %q", got)
	}
	if got := val.Limit(4, "...").String(); got != "goph..." {
		t.Fatalf("Limit = %q", got)
	}
}

func TestCharAtAndReverse(t *testing.T) {
	t.Parallel()

	r, ok := Of("gopher").CharAt(2)
	if !ok || r != 'p' {
		t.Fatalf("CharAt unexpected")
	}
	if got := Of("naïve").Reverse().String(); got != "evïan" {
		t.Fatalf("Reverse = %q", got)
	}
}
