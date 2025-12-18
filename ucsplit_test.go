package str

import (
	"reflect"
	"testing"
)

func TestUcSplit(t *testing.T) {
	got := Of("HTTPRequestID").UcSplit()
	expected := []string{"HTTP", "Request", "ID"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}

	hyphens := Of("Go-Forj_Rules").UcSplit()
	expectedHyphens := []string{"Go", "Forj", "Rules"}
	if !reflect.DeepEqual(hyphens, expectedHyphens) {
		t.Fatalf("expected %v, got %v", expectedHyphens, hyphens)
	}

	mixed := Of("goForjRocks").UcSplit()
	expectedMixed := []string{"go", "Forj", "Rocks"}
	if !reflect.DeepEqual(mixed, expectedMixed) {
		t.Fatalf("expected %v, got %v", expectedMixed, mixed)
	}

	if nilResult := Of("").UcSplit(); nilResult != nil {
		t.Fatalf("expected nil for empty input, got %v", nilResult)
	}
}
