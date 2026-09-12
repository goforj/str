package str

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

// TestLines preserves every line byte and the standard empty-input contract.
func TestLines(t *testing.T) {
	t.Parallel()
	for _, input := range []string{"", "a\r\nb\nc", "a\n", "a\rb", "a\u2028b"} {
		if got, want := slices.Collect(Of(input).Lines()), slices.Collect(strings.Lines(input)); !reflect.DeepEqual(got, want) {
			t.Fatalf("Lines(%q) = %#v, want %#v", input, got, want)
		}
	}
}
