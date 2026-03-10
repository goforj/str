package str

import "testing"

func TestCommonSuffix(t *testing.T) {
	t.Parallel()

	if got := Of("main_test.go").CommonSuffix("user_test.go", "api_test.go").String(); got != "_test.go" {
		t.Fatalf("CommonSuffix = %q", got)
	}
	if got := Of("café").CommonSuffix("fiancé").String(); got != "é" {
		t.Fatalf("CommonSuffix unicode = %q", got)
	}
	if got := Of("gopher").CommonSuffix().String(); got != "gopher" {
		t.Fatalf("CommonSuffix no args = %q", got)
	}
}
