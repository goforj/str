package str

import "testing"

// TestPadding guards its covered contract against regressions.
func TestPadding(t *testing.T) {
	t.Parallel()

	if got := Of("go").PadLeft(5, " ").String(); got != "   go" {
		t.Fatalf("PadLeft = %q", got)
	}
	if got := Of("go").PadRight(5, ".").String(); got != "go..." {
		t.Fatalf("PadRight = %q", got)
	}
	if got := Of("go").PadBoth(6, "-").String(); got != "--go--" {
		t.Fatalf("PadBoth = %q", got)
	}
	if got := Of("go").PadLeft(4, "").String(); got != "  go" {
		t.Fatalf("PadLeft default pad")
	}
	if got := Of("go").PadLeft(0, "*").String(); got != "go" {
		t.Fatalf("PadLeft length<=0")
	}
	if got := Of("go").PadLeft(2, "*").String(); got != "go" {
		t.Fatalf("PadLeft no pad when short add")
	}
	if got := Of("go").PadBoth(5, "*").String(); got != "*go**" {
		t.Fatalf("PadBoth odd spread")
	}
	if got := Of("go").PadRight(4, "").String(); got != "go  " {
		t.Fatalf("PadRight default pad")
	}
}

// TestPaddingNeverTruncates guards its covered contract against regressions.
func TestPaddingNeverTruncates(t *testing.T) {
	t.Parallel()

	widths := []int{-2, 0, 1, 2}
	directions := []struct {
		name string
		pad  func(String, int) String
	}{
		{name: "left", pad: func(s String, width int) String { return s.PadLeft(width, "*") }},
		{name: "right", pad: func(s String, width int) String { return s.PadRight(width, "*") }},
		{name: "both", pad: func(s String, width int) String { return s.PadBoth(width, "*") }},
	}

	for _, direction := range directions {
		for _, width := range widths {
			direction := direction
			width := width
			t.Run(direction.name, func(t *testing.T) {
				t.Parallel()

				if got := direction.pad(Of("猫犬"), width).String(); got != "猫犬" {
					t.Fatalf("padding to width %d = %q, want %q", width, got, "猫犬")
				}
			})
		}
	}
}

// TestPaddingRepeatsByRune guards its covered contract against regressions.
func TestPaddingRepeatsByRune(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		got  string
		want string
	}{
		{name: "left", got: Of("猫").PadLeft(4, "🙂界").String(), want: "🙂界🙂猫"},
		{name: "right", got: Of("猫").PadRight(4, "🙂界").String(), want: "猫🙂界🙂"},
		{name: "both", got: Of("猫").PadBoth(6, "🙂界").String(), want: "🙂界猫🙂界🙂"},
	}

	for _, tc := range cases {
		if tc.got != tc.want {
			t.Fatalf("%s padding = %q, want %q", tc.name, tc.got, tc.want)
		}
	}
}
