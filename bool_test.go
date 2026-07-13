package str

import "testing"

// TestBool guards its covered contract against regressions.
func TestBool(t *testing.T) {
	t.Parallel()

	got, err := Of("true").Bool()
	if err != nil {
		t.Fatalf("Bool unexpected error: %v", err)
	}
	if !got {
		t.Fatalf("Bool = %v", got)
	}
	if _, err := Of(" true ").Bool(); err == nil {
		t.Fatalf("Bool expected strict parse error")
	}
	if _, err := Of("maybe").Bool(); err == nil {
		t.Fatalf("Bool expected error")
	}
}
