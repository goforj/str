package str

import "testing"

func TestLen(t *testing.T) {
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
			got := Of(tt.in).Len()
			if got != tt.want {
				t.Fatalf("Len(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestRuneCountAlias(t *testing.T) {
	t.Parallel()

	val := Of("naïve café")
	if got, want := val.RuneCount(), val.Len(); got != want {
		t.Fatalf("RuneCount mismatch: got %d, want %d", got, want)
	}
}
