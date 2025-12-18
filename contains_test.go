package str

import "testing"

func TestContains(t *testing.T) {
	t.Parallel()

	val := Of("Go means gophers")

	if !val.Contains("gopher") {
		t.Fatalf("expected Contains to find substring")
	}
	if val.Contains() {
		t.Fatalf("empty subs should be false for Contains")
	}
	if val.Contains("rust", "beam") {
		t.Fatalf("unexpected true when no substrings match")
	}
}

func TestContainsAll(t *testing.T) {
	t.Parallel()

	val := Of("Go means gophers")

	if !val.ContainsAll("Go", "gopher") {
		t.Fatalf("expected ContainsAll to find all substrings")
	}
	if val.ContainsAll("Go", "rust") {
		t.Fatalf("expected ContainsAll to fail when any substring missing")
	}
	if val.ContainsAll() {
		t.Fatalf("empty subs should be false for ContainsAll")
	}
	if !val.ContainsAll("Go", "") {
		t.Fatalf("ContainsAll should ignore empty substrings")
	}
}

func TestContainsFold(t *testing.T) {
	t.Parallel()

	val := Of("Go means gophers")

	if !val.ContainsFold("GOPHER") {
		t.Fatalf("expected ContainsFold to match regardless of case")
	}
	if !val.ContainsFold("") {
		t.Fatalf("empty substring should return true")
	}
	if val.ContainsFold() {
		t.Fatalf("empty subs slice should be false")
	}
	if val.ContainsFold("rust") {
		t.Fatalf("unexpected match")
	}
}

func TestContainsAllFold(t *testing.T) {
	t.Parallel()

	val := Of("Go means gophers")

	if !val.ContainsAllFold("go", "GOPHER") {
		t.Fatalf("expected ContainsAllFold to match regardless of case")
	}
	if val.ContainsAllFold("go", "rust") {
		t.Fatalf("expected ContainsAllFold to fail when a substring is missing")
	}
	if val.ContainsAllFold() {
		t.Fatalf("empty subs should be false for ContainsAllFold")
	}
	if !val.ContainsAllFold("go", "") {
		t.Fatalf("ContainsAllFold should ignore empty substrings")
	}
}
