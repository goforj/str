package str

import "testing"

func TestAfter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		sep  string
		want string
	}{
		{name: "found", in: "gopher::go", sep: "::", want: "go"},
		{name: "missing", in: "gopher", sep: "::", want: "gopher"},
		{name: "empty sep", in: "gopher", sep: "", want: "gopher"},
		{name: "multibyte sep", in: "hello—world", sep: "—", want: "world"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Of(tt.in).After(tt.sep).String()
			if got != tt.want {
				t.Fatalf("After(%q,%q) = %q, want %q", tt.in, tt.sep, got, tt.want)
			}
		})
	}
}

func TestAfterLast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		sep  string
		want string
	}{
		{name: "found", in: "a/b/c", sep: "/", want: "c"},
		{name: "missing", in: "abc", sep: "/", want: "abc"},
		{name: "empty sep", in: "abc", sep: "", want: "abc"},
		{name: "multibyte sep", in: "go→lang→best", sep: "→", want: "best"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Of(tt.in).AfterLast(tt.sep).String()
			if got != tt.want {
				t.Fatalf("AfterLast(%q,%q) = %q, want %q", tt.in, tt.sep, got, tt.want)
			}
		})
	}
}
