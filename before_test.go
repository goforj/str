package str

import "testing"

func TestBeforeVariants(t *testing.T) {
	t.Parallel()

	if got := Of("gopher::go").Before("::").String(); got != "gopher" {
		t.Fatalf("Before = %q", got)
	}
	if got := Of("gopher").Before("z").String(); got != "gopher" {
		t.Fatalf("Before missing sep")
	}
	if got := Of("gopher").Before("").String(); got != "gopher" {
		t.Fatalf("Before empty sep")
	}
}

func TestBeforeLastVariants(t *testing.T) {
	t.Parallel()

	if got := Of("pkg/path/file.txt").BeforeLast("/").String(); got != "pkg/path" {
		t.Fatalf("BeforeLast = %q", got)
	}
	if got := Of("gopher").BeforeLast("z").String(); got != "gopher" {
		t.Fatalf("BeforeLast missing sep")
	}
	if got := Of("gopher").BeforeLast("").String(); got != "gopher" {
		t.Fatalf("BeforeLast empty sep")
	}
}
