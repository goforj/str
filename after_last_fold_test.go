package str

import "testing"

func TestAfterLastFold(t *testing.T) {
	t.Parallel()

	if got := Of("pkg/Path/FILE.txt").AfterLastFold("/path/").String(); got != "FILE.txt" {
		t.Fatalf("AfterLastFold = %q", got)
	}
	if got := Of("gopher").AfterLastFold("zzz").String(); got != "gopher" {
		t.Fatalf("AfterLastFold missing = %q", got)
	}
}
