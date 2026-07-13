package str

import "testing"

// TestSnake guards its covered contract against regressions.
func TestSnake(t *testing.T) {
	t.Parallel()

	if got := Of("fooBar baz").Snake().String(); got != "foo_bar_baz" {
		t.Fatalf("Snake = %q", got)
	}
	if got := Of("HTTPRequestID").Snake().String(); got != "http_request_id" {
		t.Fatalf("Snake acronym = %q", got)
	}
}
