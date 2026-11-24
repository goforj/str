package str

import (
	"math"
	"strings"
	"testing"
)

// TestExcerpt guards its covered contract against regressions.
func TestExcerpt(t *testing.T) {
	got := Of("abcde").Excerpt("c", 1, "...").String()
	if got != "...bcd..." {
		t.Fatalf("expected ...bcd..., got %s", got)
	}

	miss := Of("gopher").Excerpt("x", 2, "...").String()
	if miss != "" {
		t.Fatalf("expected empty when needle missing, got %s", miss)
	}

	defaultRadius := Of("gopher").Excerpt("oph", 0, "").String()
	if defaultRadius != "gopher" {
		t.Fatalf("expected full string when radius covers all, got %s", defaultRadius)
	}

	emptyNeedle := Of("gopher").Excerpt("", 2, "...").String()
	if emptyNeedle != "" {
		t.Fatalf("expected empty when needle empty, got %s", emptyNeedle)
	}

	prefix := Of("gopher").Excerpt("go", 2, "...").String()
	if prefix != "goph..." {
		t.Fatalf("expected goph..., got %s", prefix)
	}
}

func TestExcerptRadius(t *testing.T) {
	t.Parallel()

	longText := strings.Repeat("a", 101) + "x" + strings.Repeat("b", 101)
	defaultExcerpt := "..." + strings.Repeat("a", 100) + "x" + strings.Repeat("b", 100) + "..."
	tests := []struct {
		name   string
		input  string
		needle string
		radius int
		want   string
	}{
		{"max_radius_ascii", "abcde", "c", math.MaxInt, "abcde"},
		{"max_radius_unicode", "あいうえお", "う", math.MaxInt, "あいうえお"},
		{"exact_bounds_ascii", "abcde", "c", 2, "abcde"},
		{"exact_bounds_unicode", "あいうえお", "う", 2, "あいうえお"},
		{"right_edge", "abcde", "de", 2, "...bcde"},
		{"negative_radius", longText, "x", -1, defaultExcerpt},
		{"minimum_radius", longText, "x", math.MinInt, defaultExcerpt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Of(tt.input).Excerpt(tt.needle, tt.radius, "...").String()
			if got != tt.want {
				t.Errorf("Excerpt(%q, %d) = %q, want %q", tt.needle, tt.radius, got, tt.want)
			}
		})
	}
}
