package str

import "testing"

func TestDoesntPredicates(t *testing.T) {
	t.Parallel()

	s := Of("gophers are great")

	if s.DoesntContain("go") {
		t.Fatalf("DoesntContain should be false when match present")
	}
	if !s.DoesntContain("rust", "beam") {
		t.Fatalf("DoesntContain should be true when none match")
	}
	if !Of("gopher").DoesntContain() {
		t.Fatalf("DoesntContain should be true with no substrings")
	}
	if !s.DoesntContainFold() {
		t.Fatalf("DoesntContainFold empty slice should be true")
	}
	if !s.DoesntContain("", "rust") {
		t.Fatalf("DoesntContain should ignore empty substrings")
	}
	if !s.DoesntContainFold("", "rust") {
		t.Fatalf("DoesntContainFold should ignore empty substrings")
	}
	if Of("Gopher").DoesntContainFold("gopher") {
		t.Fatalf("DoesntContainFold should be false when match present")
	}

	if s.DoesntStartWith("go") {
		t.Fatalf("DoesntStartWith should be false when prefix present")
	}
	if !s.DoesntStartWith("rust", "beam") {
		t.Fatalf("DoesntStartWith should be true when no prefix matches")
	}
	if s.DoesntStartWithFold("GO") {
		t.Fatalf("DoesntStartWithFold should be false with match")
	}
	if !s.DoesntStartWith() {
		t.Fatalf("DoesntStartWith should be true with no prefixes")
	}
	if !s.DoesntStartWithFold() {
		t.Fatalf("DoesntStartWithFold should be true with no prefixes")
	}
	if !s.DoesntStartWithFold("rust") {
		t.Fatalf("DoesntStartWithFold should be true when no prefix matches")
	}

	if s.DoesntEndWith("great") {
		t.Fatalf("DoesntEndWith should be false when suffix present")
	}
	if !s.DoesntEndWith("rust", "beam") {
		t.Fatalf("DoesntEndWith should be true when no suffix matches")
	}
	if s.DoesntEndWithFold("GREAT") {
		t.Fatalf("DoesntEndWithFold should be false with match")
	}
	if !s.DoesntEndWith() {
		t.Fatalf("DoesntEndWith should be true with no suffixes")
	}
	if !s.DoesntEndWithFold() {
		t.Fatalf("DoesntEndWithFold should be true with no suffixes")
	}
	if !s.DoesntEndWithFold("rust") {
		t.Fatalf("DoesntEndWithFold should be true when no suffix matches")
	}
}
