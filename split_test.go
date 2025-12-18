package str

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Of("a,b,c").Split(",")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}

	empty := Of("").Split(",")
	if !reflect.DeepEqual(empty, []string{""}) {
		t.Fatalf("expected empty slice element, got %v", empty)
	}
}
