package str

import "testing"

func TestIndexFold(t *testing.T) {
	t.Parallel()

	val := Of("héllo gopher")
	if got := val.IndexFold("HÉL"); got != 0 {
		t.Fatalf("IndexFold = %d", got)
	}
	if got := val.IndexFold("zzz"); got != -1 {
		t.Fatalf("IndexFold missing = %d", got)
	}
}
