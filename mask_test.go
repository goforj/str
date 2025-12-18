package str

import "testing"

func TestMask(t *testing.T) {
	t.Parallel()

	if got := Of("gopher@example.com").Mask('*', 3, 4).String(); got != "gop***********.com" {
		t.Fatalf("Mask basic = %q", got)
	}
	if got := Of("gopher").Mask('*', 10, 10).String(); got != "gopher" {
		t.Fatalf("Mask with large reveals should be original")
	}
	if got := Of("gopher").Mask('*', -4, 1).String(); got != "go***r" {
		t.Fatalf("Mask negative reveal = %q", got)
	}
	if got := Of("go").Mask('*', -10, 0).String(); got != "**" {
		t.Fatalf("Mask with reveal clamped to zero = %q", got)
	}
	if got := Of("gopher").Mask('*', 2, -10).String(); got != "go****" {
		t.Fatalf("Mask negative right reveal = %q", got)
	}
	if got := Of("gopher").Mask(0, 1, 1).String(); got != "g****r" {
		t.Fatalf("Mask default char = %q", got)
	}
}
