package str

import "testing"

func TestCount(t *testing.T) {
	t.Parallel()

	if got := Of("gogophergo").Count("go"); got != 3 {
		t.Fatalf("Count = %d", got)
	}
	if got := Of("abc").Count(""); got != 0 {
		t.Fatalf("Count empty expected 0, got %d", got)
	}
}
