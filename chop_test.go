package str

import "testing"

// TestTrimPrefixSuffix guards its covered contract against regressions.
func TestTrimPrefixSuffix(t *testing.T) {
	t.Parallel()

	if got := Of("https://goforj.dev").TrimPrefix("https://").String(); got != "goforj.dev" {
		t.Fatalf("TrimPrefix = %q", got)
	}
	if got := Of("file.txt").TrimSuffix(".txt").String(); got != "file" {
		t.Fatalf("TrimSuffix = %q", got)
	}
	if got := Of("gopher").TrimPrefix("").String(); got != "gopher" {
		t.Fatalf("TrimPrefix empty = %q", got)
	}
	if got := Of("gopher").TrimSuffix("").String(); got != "gopher" {
		t.Fatalf("TrimSuffix empty = %q", got)
	}
	if got := Of("gopher").TrimSuffix(".zip").String(); got != "gopher" {
		t.Fatalf("TrimSuffix missing = %q", got)
	}
}
