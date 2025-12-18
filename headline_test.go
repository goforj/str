package str

import "testing"

func TestHeadline(t *testing.T) {
	got := Of("emailNotification_sent").Headline().String()
	if got != "Email Notification Sent" {
		t.Fatalf("expected Email Notification Sent, got %s", got)
	}

	smallWords := Of("go-and-the-gophers").Headline().String()
	if smallWords != "Go and the Gophers" {
		t.Fatalf("expected Go and the Gophers, got %s", smallWords)
	}

	leadingSmall := Of("and_with_go").Headline().String()
	if leadingSmall != "And with Go" {
		t.Fatalf("expected And with Go, got %s", leadingSmall)
	}
}
