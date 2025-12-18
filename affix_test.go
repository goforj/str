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
	if got := Of("gopher").ChopStart("").String(); got != "gopher" {
		t.Fatalf("ChopStart empty prefix")
	}
	if got := Of("gopher").ChopEnd(".zip").String(); got != "gopher" {
		t.Fatalf("ChopEnd no match")
	}
	if got := Of("note.md").ChopEnd(".txt", ".md").String(); got != "note" {
		t.Fatalf("ChopEnd second suffix")
	}
	if got := Of("gopher").ChopEnd().String(); got != "gopher" {
		t.Fatalf("ChopEnd empty list")
	}
	if got := Of("gopher").ChopEnd("").String(); got != "gopher" {
		t.Fatalf("ChopEnd empty suffix skip")
	}
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

func TestWrapAndUnwrap(t *testing.T) {
	t.Parallel()

	if got := Of("GoForj").Wrap("\"", "").String(); got != "\"GoForj\"" {
		t.Fatalf("Wrap = %q", got)
	}
	if got := Of("\"GoForj\"").Unwrap("\"", "\"").String(); got != "GoForj" {
		t.Fatalf("Unwrap = %q", got)
	}
	if got := Of("gopher").Unwrap("[", "]").String(); got != "gopher" {
		t.Fatalf("Unwrap no match")
	}
	if got := Of("##go##").Unwrap("#", "").String(); got != "#go#" {
		t.Fatalf("Unwrap inferred after = %q", got)
	}
}
