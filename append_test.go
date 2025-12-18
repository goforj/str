package str

import "testing"

func TestAppend(t *testing.T) {
	got := Of("Go").Append("Forj", "!").String()
	if got != "GoForj!" {
		t.Fatalf("expected GoForj!, got %s", got)
	}

	unchanged := Of("Go").Append().String()
	if unchanged != "Go" {
		t.Fatalf("expected unchanged base when no parts, got %s", unchanged)
	}
}
