package str

import "testing"

func TestLastIndexFold(t *testing.T) {
	t.Parallel()

	if got := Of("Go gopher GO").LastIndexFold("go"); got != 10 {
		t.Fatalf("LastIndexFold = %d", got)
	}
}
