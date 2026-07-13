package main

import (
	"go/ast"
	"strings"
	"testing"
)

// TestReplaceMarkedSection verifies that generation changes only the content enclosed by the requested markers.
func TestReplaceMarkedSection(t *testing.T) {
	t.Parallel()

	document := "before<!-- start -->old<!-- end -->after"
	got, err := replaceMarkedSection(document, "<!-- start -->", "<!-- end -->", "new", "example")
	if err != nil {
		t.Fatalf("replaceMarkedSection() error = %v", err)
	}

	want := "before<!-- start -->new<!-- end -->after"
	if got != want {
		t.Fatalf("replaceMarkedSection() = %q, want %q", got, want)
	}
}

// TestReplaceMarkedSectionRejectsMalformedMarkers verifies that structural README problems produce actionable errors.
func TestReplaceMarkedSectionRejectsMalformedMarkers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		document string
		want     string
	}{
		{name: "missing start", document: "<!-- end -->", want: "start marker"},
		{name: "missing end", document: "<!-- start -->", want: "end marker"},
		{name: "repeated start", document: "<!-- start --><!-- start --><!-- end -->", want: "appears 2 times"},
		{name: "repeated end", document: "<!-- start --><!-- end --><!-- end -->", want: "appears 2 times"},
		{name: "reversed", document: "<!-- end --><!-- start -->", want: "end marker precedes start marker"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := replaceMarkedSection(test.document, "<!-- start -->", "<!-- end -->", "new", "example")
			if err == nil {
				t.Fatal("replaceMarkedSection() error = nil, want malformed marker error")
			}
			if !strings.Contains(err.Error(), test.want) {
				t.Fatalf("replaceMarkedSection() error = %q, want text %q", err, test.want)
			}
		})
	}
}

// TestRenderAPI verifies stable local links and copyable examples alongside the compact index.
func TestRenderAPI(t *testing.T) {
	t.Parallel()

	symbols := []apiSymbol{
		{
			name:        "Append",
			group:       "Compose",
			receiver:    "String",
			description: "Append adds text.\nSimilar: Prepend.",
			examples:    []apiExample{{code: `v := str.Of("Go").Append("Forj")`}},
		},
		{
			name:        "Prepend",
			group:       "Compose",
			receiver:    "String",
			description: "Prepend adds leading text.",
			examples:    []apiExample{{code: `v := str.Of("Forj").Prepend("Go")`}},
		},
		{
			name:        "Of",
			group:       "Constructor",
			description: "Of starts a fluent chain.",
			examples:    []apiExample{{code: `v := str.Of("GoForj")`}},
		},
	}

	got := renderAPI(symbols)
	want := strings.Join([]string{
		"## API index",
		"",
		"The full API and these examples are also available on [pkg.go.dev](https://pkg.go.dev/github.com/goforj/str/v2).",
		"",
		"| Group | API |",
		"| --- | --- |",
		"| Compose | [Append](#append) · [Prepend](#prepend) |",
		"| Constructor | [Of](#of) |",
		"",
		"## API examples",
		"",
		"These examples come from GoDoc and run as part of the test suite.",
		"",
		"### Compose",
		"",
		`#### <a id="append"></a>Append`,
		"",
		"Append adds text.\nSimilar: Prepend.",
		"",
		"```go",
		`v := str.Of("Go").Append("Forj")`,
		"```",
		"",
		`#### <a id="prepend"></a>Prepend`,
		"",
		"Prepend adds leading text.",
		"",
		"```go",
		`v := str.Of("Forj").Prepend("Go")`,
		"```",
		"",
		"### Constructor",
		"",
		`#### <a id="of"></a>Of`,
		"",
		"Of starts a fluent chain.",
		"",
		"```go",
		`v := str.Of("GoForj")`,
		"```",
	}, "\n")
	if got != want {
		t.Fatalf("renderAPI() =\n%s\nwant:\n%s", got, want)
	}
}

// TestExtractGroupWithSimilarReference verifies that reader-facing references can precede grouping metadata without entering the compact index.
func TestExtractGroupWithSimilarReference(t *testing.T) {
	t.Parallel()

	group := &ast.CommentGroup{List: []*ast.Comment{
		{Text: "// Append returns a string with parts added."},
		{Text: "//"},
		{Text: "// Similar: Prepend."},
		{Text: "// @group Compose"},
	}}

	if got := extractGroup(group); got != "Compose" {
		t.Fatalf("extractGroup() = %q, want %q", got, "Compose")
	}
	if got := extractDescription(group); got != "Append returns a string with parts added.\n\nSimilar: Prepend." {
		t.Fatalf("extractDescription() = %q", got)
	}
}
