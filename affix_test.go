package str

import "testing"

func TestChopAndEnsure(t *testing.T) {
	t.Parallel()

	if got := Of("https://goforj.dev").ChopStart("https://", "http://").String(); got != "goforj.dev" {
		t.Fatalf("ChopStart = %q", got)
	}
	if got := Of("file.txt").ChopEnd(".txt", ".md").String(); got != "file" {
		t.Fatalf("ChopEnd = %q", got)
	}
	if got := Of("path/to").EnsurePrefix("/").String(); got != "/path/to" {
		t.Fatalf("EnsurePrefix = %q", got)
	}
	if got := Of("path/to").EnsureSuffix("/").String(); got != "path/to/" {
		t.Fatalf("EnsureSuffix = %q", got)
	}
}

func TestWrapAndUnwrap(t *testing.T) {
	t.Parallel()

	if got := Of("Laravel").Wrap("\"", "").String(); got != "\"Laravel\"" {
		t.Fatalf("Wrap = %q", got)
	}
	if got := Of("\"Laravel\"").Unwrap("\"", "\"").String(); got != "Laravel" {
		t.Fatalf("Unwrap = %q", got)
	}
}
