package str

import (
	"regexp"
	"testing"
)

func TestIsMatch(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`\d+`)
	if !Of("abc123").IsMatch(re) {
		t.Fatalf("IsMatch should be true")
	}
	if Of("abc").IsMatch(re) {
		t.Fatalf("IsMatch should be false")
	}
	if Of("abc").IsMatch(nil) {
		t.Fatalf("IsMatch nil regex should be false")
	}
}
