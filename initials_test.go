package str

import "testing"

func TestInitials(t *testing.T) {
	t.Parallel()

	if got := Of("portableNetwork graphics").Initials().String(); got != "PNG" {
		t.Fatalf("Initials = %q", got)
	}
	if got := Of("").Initials().String(); got != "" {
		t.Fatalf("Initials empty = %q", got)
	}
}
