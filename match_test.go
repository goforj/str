package str

import "testing"

// TestMatch guards its covered contract against regressions.
func TestMatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		value   string
		pattern string
		want    bool
		wantErr bool
	}{
		{name: "wildcard", value: "billing:reports", pattern: "billing:*", want: true},
		{name: "question mark", value: "file1", pattern: "file?", want: true},
		{name: "character class", value: "file2", pattern: "file[0-9]", want: true},
		{name: "whole string", value: "prefix-value-suffix", pattern: "value", want: false},
		{name: "slash boundary", value: "api/users", pattern: "*", want: false},
		{name: "invalid", value: "value", pattern: "[", wantErr: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := Of(test.value).Match(test.pattern)
			if (err != nil) != test.wantErr {
				t.Fatalf("Match(%q) error = %v, wantErr %v", test.pattern, err, test.wantErr)
			}
			if got != test.want {
				t.Fatalf("Match(%q) = %v, want %v", test.pattern, got, test.want)
			}
		})
	}
}
