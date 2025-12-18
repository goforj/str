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
	if got := Of("go").PadLeft(4, "").String(); got != "  go" {
		t.Fatalf("PadLeft default pad")
	}
	if got := Of("go").PadLeft(0, "*").String(); got != "" {
		t.Fatalf("PadLeft length<=0")
	}
	if got := Of("go").PadLeft(2, "*").String(); got != "go" {
		t.Fatalf("PadLeft no pad when short add")
	}
	if got := Of("go").PadBoth(5, "*").String(); got != "*go**" {
		t.Fatalf("PadBoth odd spread")
	}
	if got := Of("go").PadRight(4, "").String(); got != "go  " {
		t.Fatalf("PadRight default pad")
	}
}
