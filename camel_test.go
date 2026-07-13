package str

import "testing"

// TestCamel guards its covered contract against regressions.
func TestCamel(t *testing.T) {
	t.Parallel()

	if got := Of("foo_bar baz").Camel().String(); got != "fooBarBaz" {
		t.Fatalf("Camel = %q", got)
	}
	if got := Of("123").Camel().String(); got != "123" {
		t.Fatalf("Camel digits = %q", got)
	}
	if got := Of("foo bar").Camel().String(); got != "fooBar" {
		t.Fatalf("Camel multi = %q", got)
	}
	if got := Of("HTTPRequestID").Camel().String(); got != "httpRequestId" {
		t.Fatalf("Camel acronym = %q", got)
	}
}
