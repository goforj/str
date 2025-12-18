package str

import "testing"

func TestExcerpt(t *testing.T) {
	got := Of("abcde").Excerpt("c", 1, "...").String()
	if got != "...bcd..." {
		t.Fatalf("expected ...bcd..., got %s", got)
	}

	miss := Of("gopher").Excerpt("x", 2, "...").String()
	if miss != "" {
		t.Fatalf("expected empty when needle missing, got %s", miss)
	}

	defaultRadius := Of("gopher").Excerpt("oph", 0, "").String()
	if defaultRadius != "gopher" {
		t.Fatalf("expected full string when radius covers all, got %s", defaultRadius)
	}

	emptyNeedle := Of("gopher").Excerpt("", 2, "...").String()
	if emptyNeedle != "" {
		t.Fatalf("expected empty when needle empty, got %s", emptyNeedle)
	}

	prefix := Of("gopher").Excerpt("go", 2, "...").String()
	if prefix != "goph..." {
		t.Fatalf("expected goph..., got %s", prefix)
	}
}
