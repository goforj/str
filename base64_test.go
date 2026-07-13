package str

import "testing"

// TestToBase64 guards its covered contract against regressions.
func TestToBase64(t *testing.T) {
	t.Parallel()

	if got := Of("gopher").ToBase64().String(); got != "Z29waGVy" {
		t.Fatalf("ToBase64 = %q", got)
	}
}

// TestFromBase64 guards its covered contract against regressions.
func TestFromBase64(t *testing.T) {
	t.Parallel()

	value, err := Of("Z29waGVy").FromBase64()
	if err != nil {
		t.Fatalf("FromBase64 returned an error: %v", err)
	}
	if got := value.String(); got != "gopher" {
		t.Fatalf("FromBase64 = %q", got)
	}
}

// TestFromBase64Error guards its covered contract against regressions.
func TestFromBase64Error(t *testing.T) {
	t.Parallel()

	if _, err := Of("not-base64").FromBase64(); err == nil {
		t.Fatal("FromBase64 accepted invalid Base64")
	}
}
