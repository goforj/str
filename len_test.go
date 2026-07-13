package str

import "testing"

// TestRuneCount guards its covered contract against regressions.
func TestRuneCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want int
	}{
		{name: "ascii", in: "gopher", want: 6},
		{name: "emoji", in: "go 🦫", want: 4},
		{name: "empty", in: "", want: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Of(tt.in).RuneCount()
			if got != tt.want {
				t.Fatalf("RuneCount(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
