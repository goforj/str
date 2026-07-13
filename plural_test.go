package str

import "testing"

// TestPlural guards its covered contract against regressions.
func TestPlural(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"status":         "statuses",
		"statuses":       "statuses",
		"user":           "users",
		"users":          "users",
		"database":       "databases",
		"databases":      "databases",
		"address":        "addresses",
		"addresses":      "addresses",
		"class":          "classes",
		"classes":        "classes",
		"process":        "processes",
		"processes":      "processes",
		"policy":         "policies",
		"policies":       "policies",
		"query":          "queries",
		"queries":        "queries",
		"key":            "keys",
		"keys":           "keys",
		"movie":          "movies",
		"movies":         "movies",
		"cookie":         "cookies",
		"cookies":        "cookies",
		"ref":            "refs",
		"refs":           "refs",
		"archive":        "archives",
		"archives":       "archives",
		"valve":          "valves",
		"valves":         "valves",
		"city":           "cities",
		"boy":            "boys",
		"bus":            "buses",
		"match":          "matches",
		"analysis":       "analyses",
		"leaf":           "leaves",
		"knife":          "knives",
		"roof":           "roofs",
		"cafe":           "cafes",
		"datum":          "data",
		"basis":          "bases",
		"axis":           "axes",
		"cactus":         "cacti",
		"crisis":         "crises",
		"echo":           "echoes",
		"hero":           "heroes",
		"quiz":           "quizzes",
		"thesis":         "theses",
		"torpedo":        "torpedoes",
		"index":          "indices",
		"indices":        "indices",
		"matrix":         "matrices",
		"criterion":      "criteria",
		"phenomenon":     "phenomena",
		"person":         "people",
		"people":         "people",
		"fish":           "fish",
		"hardware":       "hardware",
		"media":          "media",
		"moose":          "moose",
		"police":         "police",
		"salmon":         "salmon",
		"software":       "software",
		"staff":          "staff",
		"metadata":       "metadata",
		"user_profile":   "user_profiles",
		"user_status":    "user_statuses",
		"UserStatus":     "UserStatuses",
		"UserStatuses":   "UserStatuses",
		"APIStatus":      "APIStatuses",
		"APIStatuses":    "APIStatuses",
		"City":           "Cities",
		"CITY":           "CITIES",
		"Status":         "Statuses",
		"STATUS":         "STATUSES",
		"Statuses":       "Statuses",
		"STATUSES":       "STATUSES",
		"tie":            "ties",
		"series":         "series",
		"123":            "123",
		"hello-world":    "hello-worlds",
		"Hello-World":    "Hello-Worlds",
		"iPhone":         "iPhones",
		"CharacterDatum": "CharacterData",
		"characterdatum": "characterdata",
	}

	for in, want := range cases {
		if got := Of(in).Plural().String(); got != want {
			t.Fatalf("Plural(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestSingular guards its covered contract against regressions.
func TestSingular(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"status":        "status",
		"statuses":      "status",
		"user":          "user",
		"users":         "user",
		"database":      "database",
		"databases":     "database",
		"case":          "case",
		"cases":         "case",
		"release":       "release",
		"releases":      "release",
		"response":      "response",
		"responses":     "response",
		"address":       "address",
		"addresses":     "address",
		"class":         "class",
		"process":       "process",
		"processes":     "process",
		"policy":        "policy",
		"policies":      "policy",
		"query":         "query",
		"queries":       "query",
		"key":           "key",
		"keys":          "key",
		"movie":         "movie",
		"movies":        "movie",
		"cookie":        "cookie",
		"cookies":       "cookie",
		"ref":           "ref",
		"refs":          "ref",
		"archive":       "archive",
		"archives":      "archive",
		"valve":         "valve",
		"valves":        "valve",
		"cities":        "city",
		"boys":          "boy",
		"buses":         "bus",
		"analyses":      "analysis",
		"leaves":        "leaf",
		"knives":        "knife",
		"roofs":         "roof",
		"cars":          "car",
		"data":          "datum",
		"bases":         "basis",
		"axes":          "axis",
		"cacti":         "cactus",
		"crises":        "crisis",
		"echoes":        "echo",
		"heroes":        "hero",
		"quizzes":       "quiz",
		"theses":        "thesis",
		"torpedoes":     "torpedo",
		"index":         "index",
		"indices":       "index",
		"indexes":       "index",
		"matrix":        "matrix",
		"matrices":      "matrix",
		"criterion":     "criterion",
		"criteria":      "criterion",
		"phenomenon":    "phenomenon",
		"phenomena":     "phenomenon",
		"people":        "person",
		"person":        "person",
		"fish":          "fish",
		"hardware":      "hardware",
		"media":         "media",
		"moose":         "moose",
		"police":        "police",
		"salmon":        "salmon",
		"software":      "software",
		"staff":         "staff",
		"metadata":      "metadata",
		"user_profiles": "user_profile",
		"user_status":   "user_status",
		"user_statuses": "user_status",
		"UserStatus":    "UserStatus",
		"UserStatuses":  "UserStatus",
		"APIStatus":     "APIStatus",
		"APIStatuses":   "APIStatus",
		"Cities":        "City",
		"CITIES":        "CITY",
		"Status":        "Status",
		"STATUS":        "STATUS",
		"Statuses":      "Status",
		"STATUSES":      "STATUS",
		"ties":          "tie",
		"classes":       "class",
		"aies":          "aie",
		"series":        "series",
		"123":           "123",
		"hello-worlds":  "hello-world",
		"Hello-Worlds":  "Hello-World",
		"CharacterData": "CharacterDatum",
		"characterData": "characterDatum",
		"characterdata": "characterdatum",
	}

	for in, want := range cases {
		if got := Of(in).Singular().String(); got != want {
			t.Fatalf("Singular(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestPluralAndSingularEdges guards its covered contract against regressions.
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

// TestIsTitleCase guards its covered contract against regressions.
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

// TestSingularizeWordGuards guards its covered contract against regressions.
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

// TestApplyCaseAndCamelSplit guards its covered contract against regressions.
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
