package str

import "testing"

func TestWordsSuite(t *testing.T) {
	t.Parallel()

	val := Of("Perfectly balanced, as all things should be.")
	if got := val.WordCount(); got != 7 {
		t.Fatalf("WordCount = %d", got)
	}
	if got := val.Words(3, " >>>").String(); got != "Perfectly balanced as >>>" {
		t.Fatalf("Words = %q", got)
	}
	parts := val.SplitWords()
	if len(parts) != 7 || parts[0] != "Perfectly" || parts[len(parts)-1] != "be" {
		t.Fatalf("SplitWords unexpected: %v", parts)
	}
	if got := val.FirstWord().String(); got != "Perfectly" {
		t.Fatalf("FirstWord = %q", got)
	}
	if got := val.LastWord().String(); got != "be" {
		t.Fatalf("LastWord = %q", got)
	}
	if got := Of("").FirstWord().String(); got != "" {
		t.Fatalf("FirstWord empty")
	}
	if got := Of("").LastWord().String(); got != "" {
		t.Fatalf("LastWord empty")
	}
	if got := val.Join([]string{"foo", "bar"}, "-").String(); got != "foo-bar" {
		t.Fatalf("Join = %q", got)
	}
	if got := val.Words(0, "...").String(); got != "" {
		t.Fatalf("Words zero")
	}
	if got := val.Words(10, "...").String(); got != val.String() {
		t.Fatalf("Words no truncation")
	}
	wrapped := Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
	if wrapped != "The quick brown fox\njumped over the lazy\ndog" {
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
}
