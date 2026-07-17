package str

import (
	"strings"
	"testing"
)

var benchmarkStringResult String
var benchmarkRawStringResult string

const (
	benchmarkTrimInput           = "  GoForj builds practical Go applications  "
	benchmarkToLowerInput        = "GoForj Builds Practical Go Applications"
	benchmarkNormalizeSpaceInput = " SELECT  users.id,   users.email\nFROM users\tWHERE users.status = ? "
	benchmarkTrimToLowerInput    = "  AUTH_OAuth_Provider-Name.Test  "
	benchmarkReplaceAllInput     = "archive-logs cold.storage"
)

// BenchmarkAppend measures a common multi-part fluent composition.
func BenchmarkAppend(b *testing.B) {
	for b.Loop() {
		benchmarkStringResult = Of("github.com").Append("/goforj", "/str", "/v2")
	}
}

// BenchmarkTrim measures leading and trailing whitespace removal across representative inputs.
func BenchmarkTrim(b *testing.B) {
	benchmarks := []struct {
		name  string
		value string
	}{
		{name: "ASCII", value: "  GoForj builds practical Go applications  "},
		{name: "Unicode", value: "\u2003GoForj builds practical Go applications\u00a0"},
		{name: "Clean", value: "GoForj builds practical Go applications"},
	}

	for _, benchmark := range benchmarks {
		value := Of(benchmark.value)
		b.Run(benchmark.name, func(b *testing.B) {
			for b.Loop() {
				benchmarkStringResult = value.Trim()
			}
		})
	}
}

// BenchmarkTrimComparison compares fluent trimming with strings.TrimSpace.
func BenchmarkTrimComparison(b *testing.B) {
	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkTrimStandard(benchmarkTrimInput)
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkTrimFluent(benchmarkTrimInput)
		}
	})
}

// BenchmarkToLowerComparison compares fluent lowercasing with strings.ToLower.
func BenchmarkToLowerComparison(b *testing.B) {
	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkToLowerStandard(benchmarkToLowerInput)
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkToLowerFluent(benchmarkToLowerInput)
		}
	})
}

// BenchmarkNormalizeSpace measures Unicode-aware whitespace normalization across representative inputs.
func BenchmarkNormalizeSpace(b *testing.B) {
	benchmarks := []struct {
		name  string
		value string
	}{
		{name: "ASCII", value: "  GoForj\tbuilds\n practical Go applications  "},
		{name: "Unicode", value: "\u2003GoForj\tbuilds\n practical\u2003Go applications\u00a0"},
		{name: "Clean", value: "GoForj builds practical Go applications"},
		{name: "Trimmed", value: "  GoForj builds practical Go applications  "},
		{name: "Whitespace", value: " \t\n\u2003\u00a0 "},
	}

	for _, benchmark := range benchmarks {
		value := Of(benchmark.value)
		b.Run(benchmark.name, func(b *testing.B) {
			for b.Loop() {
				benchmarkStringResult = value.NormalizeSpace()
			}
		})
	}
}

// BenchmarkNormalizeSpaceComparison compares fluent normalization with the equivalent standard-library composition.
func BenchmarkNormalizeSpaceComparison(b *testing.B) {
	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkNormalizeSpaceStandard(benchmarkNormalizeSpaceInput)
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkNormalizeSpaceFluent(benchmarkNormalizeSpaceInput)
		}
	})
}

// BenchmarkNormalizationPipeline compares fluent composition with the equivalent standard-library pipeline.
func BenchmarkNormalizationPipeline(b *testing.B) {
	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkTrimToLowerStandard(benchmarkTrimToLowerInput)
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkTrimToLowerFluent(benchmarkTrimToLowerInput)
		}
	})
}

// BenchmarkReplaceAllPipeline compares ordered fluent replacements with equivalent standard-library calls.
func BenchmarkReplaceAllPipeline(b *testing.B) {
	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkReplaceAllStandard(benchmarkReplaceAllInput)
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = benchmarkReplaceAllFluent(benchmarkReplaceAllInput)
		}
	})
}

// TestBenchmarkComparisonHelpers verifies that each fluent benchmark performs the same work as its standard-library baseline.
func TestBenchmarkComparisonHelpers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		standard func(string) string
		fluent   func(string) string
	}{
		{name: "Trim", input: benchmarkTrimInput, standard: benchmarkTrimStandard, fluent: benchmarkTrimFluent},
		{name: "ToLower", input: benchmarkToLowerInput, standard: benchmarkToLowerStandard, fluent: benchmarkToLowerFluent},
		{name: "NormalizeSpace", input: benchmarkNormalizeSpaceInput, standard: benchmarkNormalizeSpaceStandard, fluent: benchmarkNormalizeSpaceFluent},
		{name: "TrimToLower", input: benchmarkTrimToLowerInput, standard: benchmarkTrimToLowerStandard, fluent: benchmarkTrimToLowerFluent},
		{name: "ReplaceAll", input: benchmarkReplaceAllInput, standard: benchmarkReplaceAllStandard, fluent: benchmarkReplaceAllFluent},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			standard := test.standard(test.input)
			fluent := test.fluent(test.input)
			if standard != fluent {
				t.Fatalf("standard result %q does not match fluent result %q", standard, fluent)
			}
		})
	}
}

// benchmarkTrimStandard gives the baseline the same B.Loop call boundary as the fluent variant.
//
//go:noinline
func benchmarkTrimStandard(value string) string {
	return strings.TrimSpace(value)
}

// benchmarkTrimFluent gives the fluent trim the same B.Loop call boundary as its baseline.
//
//go:noinline
func benchmarkTrimFluent(value string) string {
	return Of(value).Trim().String()
}

// benchmarkToLowerStandard gives the baseline the same B.Loop call boundary as the fluent variant.
//
//go:noinline
func benchmarkToLowerStandard(value string) string {
	return strings.ToLower(value)
}

// benchmarkToLowerFluent gives the fluent lowercase operation the same B.Loop call boundary as its baseline.
//
//go:noinline
func benchmarkToLowerFluent(value string) string {
	return Of(value).ToLower().String()
}

// benchmarkNormalizeSpaceStandard gives the baseline the same B.Loop call boundary as the fluent variant.
//
//go:noinline
func benchmarkNormalizeSpaceStandard(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// benchmarkNormalizeSpaceFluent gives fluent normalization the same B.Loop call boundary as its baseline.
//
//go:noinline
func benchmarkNormalizeSpaceFluent(value string) string {
	return Of(value).NormalizeSpace().String()
}

// benchmarkTrimToLowerStandard gives the baseline the same B.Loop call boundary as the fluent variant.
//
//go:noinline
func benchmarkTrimToLowerStandard(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

// benchmarkTrimToLowerFluent gives the fluent pipeline the same B.Loop call boundary as its baseline.
//
//go:noinline
func benchmarkTrimToLowerFluent(value string) string {
	return Of(value).Trim().ToLower().String()
}

// benchmarkReplaceAllStandard gives the baseline the same B.Loop call boundary as the fluent variant.
//
//go:noinline
func benchmarkReplaceAllStandard(value string) string {
	result := strings.ReplaceAll(value, "-", "_")
	result = strings.ReplaceAll(result, " ", "_")
	return strings.ReplaceAll(result, ".", "_")
}

// benchmarkReplaceAllFluent gives the fluent pipeline the same B.Loop call boundary as its baseline.
//
//go:noinline
func benchmarkReplaceAllFluent(value string) string {
	return Of(value).
		ReplaceAll("-", "_").
		ReplaceAll(" ", "_").
		ReplaceAll(".", "_").
		String()
}

// BenchmarkReplaceFold measures repeated Unicode simple-fold replacement.
func BenchmarkReplaceFold(b *testing.B) {
	value := Of("Go Σ go ς GO σ gopher")
	for b.Loop() {
		benchmarkStringResult = value.ReplaceFold("go", "Go")
	}
}

// BenchmarkSnake measures shared tokenization and identifier casing.
func BenchmarkSnake(b *testing.B) {
	value := Of("HTTPRequestIDWithÜnicode")
	for b.Loop() {
		benchmarkStringResult = value.Snake()
	}
}
