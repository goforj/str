package str

import "testing"

func TestPlural(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"city":          "cities",
		"boy":           "boys",
		"bus":           "buses",
		"match":         "matches",
		"analysis":      "analyses",
		"leaf":          "leaves",
		"knife":         "knives",
		"roof":          "roofs",
		"cafe":          "cafes",
		"datum":         "data",
		"basis":         "bases",
		"axis":          "axes",
		"cactus":        "cacti",
		"crisis":        "crises",
		"echo":          "echoes",
		"hero":          "heroes",
		"quiz":          "quizzes",
		"thesis":        "theses",
		"torpedo":       "torpedoes",
		"person":        "people",
		"people":        "people",
		"fish":          "fish",
		"hardware":      "hardware",
		"media":         "media",
		"moose":         "moose",
		"police":        "police",
		"salmon":        "salmon",
		"software":      "software",
		"staff":         "staff",
		"user_profile":  "user_profiles",
		"City":          "Cities",
		"CITY":          "CITIES",
		"tie":           "ties",
		"series":        "series",
		"123":           "123",
		"hello-world":   "hello-worlds",
		"Hello-World":   "Hello-Worlds",
		"iPhone":        "iPhones",
		"CharacterDatum": "CharacterData",
		"characterdatum": "characterdata",
	}

	for in, want := range cases {
		if got := Of(in).Plural().String(); got != want {
			t.Fatalf("Plural(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSingular(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"cities":         "city",
		"boys":           "boy",
		"buses":          "bus",
		"analyses":       "analysis",
		"leaves":         "leaf",
		"knives":         "knife",
		"roofs":          "roof",
		"cars":           "car",
		"data":           "datum",
		"bases":          "basis",
		"axes":           "axis",
		"cacti":          "cactus",
		"crises":         "crisis",
		"echoes":         "echo",
		"heroes":         "hero",
		"quizzes":        "quiz",
		"theses":         "thesis",
		"torpedoes":      "torpedo",
		"people":         "person",
		"person":         "person",
		"fish":           "fish",
		"hardware":       "hardware",
		"media":          "media",
		"moose":          "moose",
		"police":         "police",
		"salmon":         "salmon",
		"software":       "software",
		"staff":          "staff",
		"user_profiles":  "user_profile",
		"Cities":         "City",
		"CITIES":         "CITY",
		"ties":           "tie",
		"classes":        "class",
		"aies":           "aie",
		"series":         "series",
		"123":            "123",
		"hello-worlds":   "hello-world",
		"Hello-Worlds":   "Hello-World",
		"CharacterData":  "CharacterDatum",
		"characterData":  "characterDatum",
		"characterdata":  "characterdatum",
	}

	for in, want := range cases {
		if got := Of(in).Singular().String(); got != want {
			t.Fatalf("Singular(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestPluralAndSingularEdges(t *testing.T) {
	t.Parallel()

	if got := Of("").Plural().String(); got != "" {
		t.Fatalf("Plural empty = %q", got)
	}
	if got := Of("").Singular().String(); got != "" {
		t.Fatalf("Singular empty = %q", got)
	}
	if got := Of("...").Plural().String(); got != "..." {
		t.Fatalf("Plural punctuation = %q", got)
	}
	if got := Of("...").Singular().String(); got != "..." {
		t.Fatalf("Singular punctuation = %q", got)
	}
	if got := titleCase(""); got != "" {
		t.Fatalf("titleCase empty = %q", got)
	}
}

func TestIsTitleCase(t *testing.T) {
	t.Parallel()

	if !isTitleCase("City") {
		t.Fatalf("isTitleCase expected true")
	}
	if isTitleCase("city") {
		t.Fatalf("isTitleCase lower expected false")
	}
	if isTitleCase("CiTy") {
		t.Fatalf("isTitleCase mixed expected false")
	}
	if isTitleCase("123") {
		t.Fatalf("isTitleCase digits expected false")
	}
}

func TestSingularizeWordGuards(t *testing.T) {
	t.Parallel()

	if got := singularizeWord(""); got != "" {
		t.Fatalf("singularizeWord empty = %q", got)
	}
	if got := singularizeWord("123"); got != "123" {
		t.Fatalf("singularizeWord digits = %q", got)
	}
	if got := singularizeWord("go"); got != "go" {
		t.Fatalf("singularizeWord unchanged = %q", got)
	}
}

func TestApplyCaseAndCamelSplit(t *testing.T) {
	t.Parallel()

	if got := applyCase("iPhone", "phones"); got != "phones" {
		t.Fatalf("applyCase mixed = %q", got)
	}
	if got := lastCamelWordSplit("X"); got != "" {
		t.Fatalf("lastCamelWordSplit short = %q", got)
	}
	if got := lastCamelWordSplit("gopher"); got != "" {
		t.Fatalf("lastCamelWordSplit none = %q", got)
	}
	if got := lastCamelWordSplit("ApiClient"); got != "Client" {
		t.Fatalf("lastCamelWordSplit camel = %q", got)
	}
}
