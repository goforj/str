package str

import "testing"

func TestIsASCII(t *testing.T) {
	t.Parallel()

	if !Of("gopher").IsASCII() {
		t.Fatalf("IsASCII should be true")
	}
	if Of("gophers 🦫").IsASCII() {
		t.Fatalf("IsASCII should be false with emoji")
	}
}
