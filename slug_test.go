package str

import "testing"

func TestSlug(t *testing.T) {
	t.Parallel()

	if got := Of("Go Forj Toolkit").Slug("-").String(); got != "go-forj-toolkit" {
		t.Fatalf("Slug = %q", got)
	}
	if got := Of("Crème brûlée!").Slug("-").String(); got != "creme-brulee" {
		t.Fatalf("Slug accent handling = %q", got)
	}
	if got := Of("ø ú ý").Slug("-").String(); got != "o-u-y" {
		t.Fatalf("Slug other accents = %q", got)
	}
	if got := Of("foo   bar").Slug("-").String(); got != "foo-bar" {
		t.Fatalf("Slug collapse separators = %q", got)
	}
	if got := Of("!foo").Slug("").String(); got != "foo" {
		t.Fatalf("Slug default sep = %q", got)
	}
}
