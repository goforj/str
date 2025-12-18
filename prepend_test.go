package str

import "testing"

func TestPrepend(t *testing.T) {
	got := Of("World").Prepend("Hello ", "Go ").String()
	if got != "Hello Go World" {
		t.Fatalf("expected Hello Go World, got %s", got)
	}

	unchanged := Of("Go").Prepend().String()
	if unchanged != "Go" {
		t.Fatalf("expected unchanged base when no parts, got %s", unchanged)
	}
}
