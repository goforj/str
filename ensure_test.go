package str

import "testing"

// TestEnsurePrefixSuffix guards its covered contract against regressions.
func TestEnsurePrefixSuffix(t *testing.T) {
	t.Parallel()

	if got := Of("path/to").EnsurePrefix("/").String(); got != "/path/to" {
		t.Fatalf("EnsurePrefix = %q", got)
	}
	if got := Of("path/to").EnsureSuffix("/").String(); got != "path/to/" {
		t.Fatalf("EnsureSuffix = %q", got)
	}
	if got := Of("/path").EnsurePrefix("/").String(); got != "/path" {
		t.Fatalf("EnsurePrefix already present")
	}
	if got := Of("path/").EnsureSuffix("/").String(); got != "path/" {
		t.Fatalf("EnsureSuffix already present")
	}
}
