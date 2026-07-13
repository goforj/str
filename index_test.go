package str

import "testing"

// TestIndexing guards its covered contract against regressions.
func TestIndexing(t *testing.T) {
	t.Parallel()

	val := Of("héllo gopher")
	if got := val.Index("llo"); got != 2 {
		t.Fatalf("Index = %d", got)
	}
	if got := val.LastIndex("o"); got != 7 {
		t.Fatalf("LastIndex = %d", got)
	}
	if got := val.Slice(2, 6).String(); got != "llo " {
		t.Fatalf("Slice = %q", got)
	}
	if got := val.Slice(20, 30).String(); got != "" {
		t.Fatalf("Slice clamp beyond length %q", got)
	}
	if got := val.Slice(5, 3).String(); got != "" {
		t.Fatalf("Slice start>=end %q", got)
	}
	if got := val.Index(""); got != -1 {
		t.Fatalf("Index empty = %d", got)
	}
	if got := val.LastIndex(""); got != -1 {
		t.Fatalf("LastIndex empty = %d", got)
	}
}
