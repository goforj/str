package str

import "testing"

func TestAfterFold(t *testing.T) {
	t.Parallel()

	if got := Of("GoPHER::go").AfterFold("::GO").String(); got != "" {
		t.Fatalf("AfterFold = %q", got)
	}
	if got := Of("gopher").AfterFold("zzz").String(); got != "gopher" {
		t.Fatalf("AfterFold missing = %q", got)
	}
}
