package str

import "testing"

func TestNewLine(t *testing.T) {
	got := Of("hello").NewLine().Append("world").String()
	if got != "hello\nworld" {
		t.Fatalf("expected hello\\nworld, got %s", got)
	}
}
