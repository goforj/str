package str

import (
	"regexp"
	"strings"
	"testing"
)

func TestReplaceVariants(t *testing.T) {
	t.Parallel()

	val := Of("gopher gopher")
	if got := val.ReplaceFirst("gopher", "go").String(); got != "go gopher" {
		t.Fatalf("ReplaceFirst = %q", got)
	}
	if got := val.ReplaceLast("gopher", "go").String(); got != "gopher go" {
		t.Fatalf("ReplaceLast = %q", got)
	}
	if got := val.ReplaceLast("rust", "go").String(); got != "gopher gopher" {
		t.Fatalf("ReplaceLast no-op = %q", got)
	}
	if got := val.ReplaceArray([]string{"gopher", "go"}, "x").String(); got != "x x" {
		t.Fatalf("ReplaceArray = %q", got)
	}
	if got := val.ReplaceArray([]string{""}, "x").String(); got != "gopher gopher" {
		t.Fatalf("ReplaceArray empty skip = %q", got)
	}

	re := regexp.MustCompile(`go\w+`)
	if got := val.ReplaceMatches(re, strings.ToUpper).String(); got != "GOPHER GOPHER" {
		t.Fatalf("ReplaceMatches = %q", got)
	}
}

func TestSwapRemoveCountRepeat(t *testing.T) {
	t.Parallel()

	val := Of("Gophers are great!")
	pairs := map[string]string{"Gophers": "GoForj", "great": "fantastic"}
	if got := val.Swap(pairs).String(); got != "GoForj are fantastic!" {
		t.Fatalf("Swap = %q", got)
	}
	if got := val.Swap(map[string]string{}).String(); got != "Gophers are great!" {
		t.Fatalf("Swap empty map should no-op")
	}

	if got := Of("The Go Toolkit").Remove("Go ").String(); got != "The Toolkit" {
		t.Fatalf("Remove = %q", got)
	}
	if got := Of("gopher").Remove("").String(); got != "gopher" {
		t.Fatalf("Remove empty substring should no-op")
	}
	if got := Of("gopher gopher").ReplaceArray([]string{"gopher"}, "").String(); got != " " {
		t.Fatalf("ReplaceArray to empty = %q", got)
	}

	if got := Of("gogophergo").Count("go"); got != 3 {
		t.Fatalf("Count = %d", got)
	}

	if got := Of("go").Repeat(3).String(); got != "gogogo" {
		t.Fatalf("Repeat = %q", got)
	}
}
