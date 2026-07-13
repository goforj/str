package str

import "testing"

// TestInitials guards its covered contract against regressions.
func TestInitials(t *testing.T) {
	t.Parallel()

	if got := Of("portableNetwork graphics").Initials().String(); got != "PNG" {
		t.Fatalf("Initials = %q", got)
	}
	if got := Of("").Initials().String(); got != "" {
		t.Fatalf("Initials empty = %q", got)
	}
	if got := Of("HTTPRequestID").Initials().String(); got != "HRI" {
		t.Fatalf("Initials acronym = %q", got)
	}
}
