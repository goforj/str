package str

import "testing"

func TestFloat64(t *testing.T) {
	t.Parallel()

	got, err := Of("3.14").Float64()
	if err != nil {
		t.Fatalf("Float64 unexpected error: %v", err)
	}
	if got != 3.14 {
		t.Fatalf("Float64 = %v", got)
	}
	if _, err := Of(" 3.14 ").Float64(); err == nil {
		t.Fatalf("Float64 expected strict parse error")
	}
	if _, err := Of("nope").Float64(); err == nil {
		t.Fatalf("Float64 expected error")
	}
}
