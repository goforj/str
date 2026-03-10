package str

import "testing"

func TestInt(t *testing.T) {
	t.Parallel()

	got, err := Of("42").Int()
	if err != nil {
		t.Fatalf("Int unexpected error: %v", err)
	}
	if got != 42 {
		t.Fatalf("Int = %d", got)
	}
	if _, err := Of(" 42 ").Int(); err == nil {
		t.Fatalf("Int expected strict parse error")
	}
	if _, err := Of("nope").Int(); err == nil {
		t.Fatalf("Int expected error")
	}
}
