package str

import "testing"

func TestBeforeFold(t *testing.T) {
	t.Parallel()

	if got := Of("GoPHER::go").BeforeFold("::GO").String(); got != "GoPHER" {
		t.Fatalf("BeforeFold = %q", got)
	}
	if got := Of("gopher").BeforeFold("zzz").String(); got != "gopher" {
		t.Fatalf("BeforeFold missing = %q", got)
	}
}
