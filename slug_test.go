package str

import "testing"

func TestSlug(t *testing.T) {
	t.Parallel()

	if got := Of("Laravel Framework").Slug("-").String(); got != "laravel-framework" {
		t.Fatalf("Slug = %q", got)
	}
	if got := Of("Crème brûlée!").Slug("-").String(); got != "creme-brulee" {
		t.Fatalf("Slug accent handling = %q", got)
	}
}
