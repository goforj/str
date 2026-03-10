package str

import "testing"

func TestIsAlnum(t *testing.T) {
	t.Parallel()

	if !Of("Go2025").IsAlnum() {
		t.Fatalf("IsAlnum expected true")
	}
	if !Of("١٢٣abc").IsAlnum() {
		t.Fatalf("IsAlnum unicode expected true")
	}
	if Of("go-forj").IsAlnum() {
		t.Fatalf("IsAlnum expected false")
	}
}
