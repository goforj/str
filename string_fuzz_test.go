package str

import (
	"testing"
	"unicode/utf8"
)

// FuzzStringInvariants verifies core transformations remain stable for arbitrary valid UTF-8.
func FuzzStringInvariants(f *testing.F) {
	for _, seed := range []string{
		"",
		"GoForj",
		"  HTTPRequestID  ",
		"naïve café",
		"Σίσυφος",
		"emoji 🦫 string",
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, value string) {
		if !utf8.ValidString(value) {
			t.Skip()
		}

		wrapped := Of(value)
		if got := wrapped.Reverse().Reverse().String(); got != value {
			t.Fatalf("double Reverse = %q, want %q", got, value)
		}

		trimmed := wrapped.Trim()
		if got := trimmed.Trim().String(); got != trimmed.String() {
			t.Fatalf("Trim is not idempotent: %q then %q", trimmed.String(), got)
		}

		normalized := wrapped.NormalizeSpace()
		if got := normalized.NormalizeSpace().String(); got != normalized.String() {
			t.Fatalf("NormalizeSpace is not idempotent: %q then %q", normalized.String(), got)
		}

		if got := wrapped.ReplaceAll("", "x").String(); got != value {
			t.Fatalf("ReplaceAll with empty search = %q, want %q", got, value)
		}
		if got := wrapped.ReplaceFirst("", "x").String(); got != value {
			t.Fatalf("ReplaceFirst with empty search = %q, want %q", got, value)
		}
		if got := wrapped.ReplaceLast("", "x").String(); got != value {
			t.Fatalf("ReplaceLast with empty search = %q, want %q", got, value)
		}
	})
}
