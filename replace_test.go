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
	if got := val.ReplaceArray([]string{"gopher", "go"}, "x").String(); got != "x x" {
		t.Fatalf("ReplaceArray = %q", got)
	}

	re := regexp.MustCompile(`go\w+`)
	if got := val.ReplaceMatches(re, strings.ToUpper).String(); got != "GOPHER GOPHER" {
		t.Fatalf("ReplaceMatches = %q", got)
	}
}

func TestSwapRemoveCountRepeat(t *testing.T) {
	t.Parallel()

	val := Of("Tacos are great!")
	pairs := map[string]string{"Tacos": "Burritos", "great": "fantastic"}
	if got := val.Swap(pairs).String(); got != "Burritos are fantastic!" {
		t.Fatalf("Swap = %q", got)
	}

	if got := Of("The Laravel Framework").Remove("Laravel ").String(); got != "The Framework" {
		t.Fatalf("Remove = %q", got)
	}

	if got := Of("gogophergo").Count("go"); got != 3 {
		t.Fatalf("Count = %d", got)
	}

	if got := Of("go").Repeat(3).String(); got != "gogogo" {
		t.Fatalf("Repeat = %q", got)
	}
}
