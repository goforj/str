package str

import "testing"

func TestIsWildcard(t *testing.T) {
	t.Parallel()

	if !Of("foo/bar").Is("foo/*") {
		t.Fatalf("Is wildcard expected true")
	}
	if Of("foo/bar").Is() {
		t.Fatalf("Is empty patterns should be false")
	}
	if Of("foo/bar").Is("baz*") {
		t.Fatalf("Is non-match should be false")
	}

	if Of("foo").Is("") {
		t.Fatalf("empty pattern should be ignored and return false")
	}

	if Of("go").Is("[") {
		t.Fatalf("Invalid pattern not contained should be false")
	}
	if !Of("go[forj").Is("[") {
		t.Fatalf("Invalid pattern with substring should use fallback contains")
	}

	if !matchWildcard("go[a-", "[a-") {
		t.Fatalf("matchWildcard should fall back on bad pattern")
	}
}
