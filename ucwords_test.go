package str

import "testing"

func TestUcWords(t *testing.T) {
	got := Of("hello WORLD").UcWords().String()
	if got != "Hello WORLD" {
		t.Fatalf("expected Hello WORLD, got %s", got)
	}

	withDelim := Of("go-forj rules").UcWords().String()
	if withDelim != "Go-Forj Rules" {
		t.Fatalf("expected Go-Forj Rules, got %s", withDelim)
	}
}
