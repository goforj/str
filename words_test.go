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
	if got := val.Join([]string{"foo", "bar"}, "-").String(); got != "foo-bar" {
		t.Fatalf("Join = %q", got)
	}
	wrapped := Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
	if wrapped != "The quick brown fox\njumped over the lazy\ndog" {
		t.Fatalf("WrapWords = %q", wrapped)
	}
}
