package str

import "testing"

// TestRepeat preserves zero counts and fails fast on negative counts.
func TestRepeat(t *testing.T) {
	t.Parallel()
	if got := Of("go").Repeat(3).String(); got != "gogogo" {
		t.Fatalf("Repeat = %q", got)
	}
	if got := Of("go").Repeat(0).String(); got != "" {
		t.Fatalf("Repeat zero = %q", got)
	}
	for _, input := range []string{"", "go"} {
		t.Run(input, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("negative Repeat did not panic")
				}
			}()
			Of(input).Repeat(-1)
		})
	}
}
