package str

import "testing"

// TestKebab guards its covered contract against regressions.
func TestKebab(t *testing.T) {
	t.Parallel()

	if got := Of("fooBar baz").Kebab().String(); got != "foo-bar-baz" {
		t.Fatalf("Kebab = %q", got)
	}
	if got := Of("HTTPRequestID").Kebab().String(); got != "http-request-id" {
		t.Fatalf("Kebab acronym = %q", got)
	}
}
