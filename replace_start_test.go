package str

import "testing"

// TestReplacePrefixSuffix guards its covered contract against regressions.
func TestReplacePrefixSuffix(t *testing.T) {
	t.Parallel()

	val := Of("prefix-value")
	if got := val.ReplacePrefix("prefix-", "new-").String(); got != "new-value" {
		t.Fatalf("ReplacePrefix = %q", got)
	}
	if got := val.ReplacePrefix("missing", "new-").String(); got != "prefix-value" {
		t.Fatalf("ReplacePrefix missing = %q", got)
	}
	if got := val.ReplacePrefix("", "new-").String(); got != "prefix-value" {
		t.Fatalf("ReplacePrefix empty old = %q", got)
	}

	val = Of("file.old")
	if got := val.ReplaceSuffix(".old", ".new").String(); got != "file.new" {
		t.Fatalf("ReplaceSuffix = %q", got)
	}
	if got := val.ReplaceSuffix(".missing", ".new").String(); got != "file.old" {
		t.Fatalf("ReplaceSuffix missing = %q", got)
	}
	if got := val.ReplaceSuffix("", ".new").String(); got != "file.old" {
		t.Fatalf("ReplaceSuffix empty old = %q", got)
	}
}
