package str

import "testing"

func TestToBase64(t *testing.T) {
	got := Of("gopher").ToBase64().String()
	if got != "Z29waGVy" {
		t.Fatalf("expected Z29waGVy, got %s", got)
	}
}

func TestFromBase64(t *testing.T) {
	val, err := Of("Z29waGVy").FromBase64()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if val.String() != "gopher" {
		t.Fatalf("expected gopher, got %s", val.String())
	}
}

func TestFromBase64Error(t *testing.T) {
	_, err := Of("not-base64").FromBase64()
	if err == nil {
		t.Fatalf("expected error for invalid base64 input")
	}
}
