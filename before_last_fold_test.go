package str

import "testing"

func TestBeforeLastFold(t *testing.T) {
	t.Parallel()

	if got := Of("pkg/Path/FILE.txt").BeforeLastFold("/path/").String(); got != "pkg" {
		t.Fatalf("BeforeLastFold = %q", got)
	}
	if got := Of("gopher").BeforeLastFold("zzz").String(); got != "gopher" {
		t.Fatalf("BeforeLastFold missing = %q", got)
	}
}
