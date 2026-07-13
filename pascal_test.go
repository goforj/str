package str

import "testing"

// TestPascal guards its covered contract against regressions.
func TestPascal(t *testing.T) {
	t.Parallel()

	if got := Of("foo_bar baz").Pascal().String(); got != "FooBarBaz" {
		t.Fatalf("Pascal = %q", got)
	}
	if got := Of("foo bar baz").Pascal().String(); got != "FooBarBaz" {
		t.Fatalf("Pascal multi = %q", got)
	}
	if got := Of("fooBAR").Pascal().String(); got != "FooBar" {
		t.Fatalf("Pascal mixed = %q", got)
	}
	if got := Of("HTTPRequestID").Pascal().String(); got != "HttpRequestId" {
		t.Fatalf("Pascal acronym = %q", got)
	}
}
