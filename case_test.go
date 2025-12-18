package str

import "testing"

func TestCaseTransforms(t *testing.T) {
	t.Parallel()

	if got := Of("GoLang").ToLower().String(); got != "golang" {
		t.Fatalf("ToLower = %q", got)
	}
	if got := Of("GoLang").ToUpper().String(); got != "GOLANG" {
		t.Fatalf("ToUpper = %q", got)
	}
	if got := Of("ß").ToTitle().String(); got != "SS" {
		t.Fatalf("ToTitle = %q", got)
	}
	if got := Of("a nice title uses the correct case").Title().String(); got != "A Nice Title Uses The Correct Case" {
		t.Fatalf("Title = %q", got)
	}
	if got := Of("gopher").UcFirst().String(); got != "Gopher" {
		t.Fatalf("UcFirst = %q", got)
	}
	if got := Of("Gopher").LcFirst().String(); got != "gopher" {
		t.Fatalf("LcFirst = %q", got)
	}
	if got := Of("foo_bar baz").Camel().String(); got != "fooBarBaz" {
		t.Fatalf("Camel = %q", got)
	}
	if got := Of("foo_bar baz").Pascal().String(); got != "FooBarBaz" {
		t.Fatalf("Pascal = %q", got)
	}
	if got := Of("fooBar baz").Snake("_").String(); got != "foo_bar_baz" {
		t.Fatalf("Snake = %q", got)
	}
	if got := Of("fooBar baz").Kebab().String(); got != "foo-bar-baz" {
		t.Fatalf("Kebab = %q", got)
	}
}
