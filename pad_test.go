package str

import "testing"

func TestPadding(t *testing.T) {
	t.Parallel()

	if got := Of("go").PadLeft(5, " ").String(); got != "   go" {
		t.Fatalf("PadLeft = %q", got)
	}
	if got := Of("go").PadRight(5, ".").String(); got != "go..." {
		t.Fatalf("PadRight = %q", got)
	}
	if got := Of("go").PadBoth(6, "-").String(); got != "--go--" {
		t.Fatalf("PadBoth = %q", got)
	}
}
