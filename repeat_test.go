package str

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()

	if got := Of("go").Repeat(3).String(); got != "gogogo" {
		t.Fatalf("Repeat = %q", got)
	}
	if got := Of("go").Repeat(0).String(); got != "" {
		t.Fatalf("Repeat zero")
	}
	if got := Of("go").Repeat(-1).String(); got != "" {
		t.Fatalf("Repeat negative")
	}
}
