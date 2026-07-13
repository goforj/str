package str

import "testing"

// TestSlug guards its covered contract against regressions.
func TestSlug(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "words", in: "Go Forj Toolkit", want: "go-forj-toolkit"},
		{name: "accented letters", in: "Crème brûlée!", want: "crème-brûlée"},
		{name: "non-ASCII scripts", in: "Привет, 世界 １２", want: "привет-世界-１２"},
		{name: "Unicode lowercase", in: "Ø Ú Ý", want: "ø-ú-ý"},
		{name: "collapse separators", in: "foo _ / bar", want: "foo-bar"},
		{name: "trim separators", in: "--foo--", want: "foo"},
		{name: "digits", in: "Version 2.0", want: "version-2-0"},
		{name: "no words", in: "! -- _", want: ""},
		{name: "empty", in: "", want: ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := Of(tc.in).Slug().String(); got != tc.want {
				t.Fatalf("Slug(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
