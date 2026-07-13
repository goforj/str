package str

import (
	"reflect"
	"testing"
)

// TestWordsSuite guards its covered contract against regressions.
func TestWordsSuite(t *testing.T) {
	t.Parallel()

	val := Of("Perfectly balanced, as all things should be.")
	if got := val.WordCount(); got != 7 {
		t.Fatalf("WordCount = %d", got)
	}
	if got := val.Words(3, " >>>").String(); got != "Perfectly balanced, as >>>" {
		t.Fatalf("Words = %q", got)
	}
	parts := val.SplitWords()
	if len(parts) != 7 || parts[0] != "Perfectly" || parts[len(parts)-1] != "be" {
		t.Fatalf("SplitWords unexpected: %v", parts)
	}
	if got := Of("HTTPRequestID").SplitWords(); !reflect.DeepEqual(got, []string{"HTTP", "Request", "ID"}) {
		t.Fatalf("SplitWords acronym = %v", got)
	}
	if got := Of("").SplitWords(); got != nil {
		t.Fatalf("SplitWords empty = %v", got)
	}
	if got := val.FirstWord().String(); got != "Perfectly" {
		t.Fatalf("FirstWord = %q", got)
	}
	if got := val.LastWord().String(); got != "be" {
		t.Fatalf("LastWord = %q", got)
	}
	if got := Of("HTTPRequestID").FirstWord().String(); got != "HTTP" {
		t.Fatalf("FirstWord acronym = %q", got)
	}
	if got := Of("HTTPRequestID").LastWord().String(); got != "ID" {
		t.Fatalf("LastWord acronym = %q", got)
	}
	if got := Of("HTTPRequestID").WordCount(); got != 3 {
		t.Fatalf("WordCount acronym = %d", got)
	}
	if got := Of("").FirstWord().String(); got != "" {
		t.Fatalf("FirstWord empty = %q", got)
	}
	if got := Of("").LastWord().String(); got != "" {
		t.Fatalf("LastWord empty = %q", got)
	}
	if got := val.Words(0, "...").String(); got != "" {
		t.Fatalf("Words zero")
	}
	if got := val.Words(10, "...").String(); got != val.String() {
		t.Fatalf("Words no truncation")
	}
	if got := Of("  One,  two; three four").Words(3, "...").String(); got != "  One,  two; three..." {
		t.Fatalf("Words original prefix = %q", got)
	}
	if got := Of("HTTPRequestID tail").Words(2, "...").String(); got != "HTTPRequest..." {
		t.Fatalf("Words acronym boundary = %q", got)
	}
	wrapped := Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
	if wrapped != "The quick brown fox\njumped over the lazy\ndog." {
		t.Fatalf("WrapWords = %q", wrapped)
	}
	if got := Of("foo bar baz").WrapWords(3, "").String(); got != "foo\nbar\nbaz" {
		t.Fatalf("WrapWords default break %q", got)
	}
	if got := Of("foo bar").WrapWords(50, "").String(); got != "foo bar" {
		t.Fatalf("WrapWords wide %q", got)
	}
	if got := Of("foo bar").WrapWords(3, "|").String(); got != "foo|bar" {
		t.Fatalf("WrapWords narrow %q", got)
	}
	if got := Of("   ").WrapWords(10, "|").String(); got != "   " {
		t.Fatalf("WrapWords no words %q", got)
	}
	if got := Of("foo bar").WrapWords(0, "\n").String(); got != "foo bar" {
		t.Fatalf("WrapWords width<=0 %q", got)
	}
	if got := Of("Wait... what? Yes!").WrapWords(8, "|").String(); got != "Wait...|what?|Yes!" {
		t.Fatalf("WrapWords punctuation = %q", got)
	}
	if got := Of("foo/bar baz-qux.").WrapWords(8, "|").String(); got != "foo/bar|baz-qux." {
		t.Fatalf("WrapWords content = %q", got)
	}
	if got := Of("éé éé").WrapWords(5, "|").String(); got != "éé éé" {
		t.Fatalf("WrapWords rune width = %q", got)
	}
	if got := Of("éé éé").WrapWords(4, "|").String(); got != "éé|éé" {
		t.Fatalf("WrapWords rune boundary = %q", got)
	}
	if got := Of("foo  bar ").WrapWords(8, "|").String(); got != "foo  bar " {
		t.Fatalf("WrapWords whitespace = %q", got)
	}
}
