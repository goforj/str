package str

import "testing"

func TestCountFold(t *testing.T) {
	t.Parallel()

	if got := Of("GoGOgophergo").CountFold("go"); got != 4 {
		t.Fatalf("CountFold = %d", got)
	}
	if got := Of("gopher").CountFold(""); got != 0 {
		t.Fatalf("CountFold empty = %d", got)
	}
	if got := Of("gopher").CountFold("zzz"); got != 0 {
		t.Fatalf("CountFold missing = %d", got)
	}
}
