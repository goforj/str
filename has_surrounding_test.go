package str

import "testing"

func TestHasSurrounding(t *testing.T) {
	t.Parallel()

	if !Of(`"GoForj"`).HasSurrounding(`"`, "") {
		t.Fatalf("HasSurrounding expected true")
	}
	if !Of("[go]").HasSurrounding("[", "]") {
		t.Fatalf("HasSurrounding explicit expected true")
	}
	if Of("gopher").HasSurrounding(`"`, "") {
		t.Fatalf("HasSurrounding expected false")
	}
}
