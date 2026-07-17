package str

import (
	"strings"
	"testing"
)

var benchmarkStringResult String
var benchmarkRawStringResult string

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
	const value = " SELECT  users.id,   users.email\nFROM users\tWHERE users.status = ? "

	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = strings.Join(strings.Fields(value), " ")
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = Of(value).NormalizeSpace().String()
		}
	})
}

// BenchmarkNormalizationPipeline compares fluent composition with the equivalent standard-library pipeline.
func BenchmarkNormalizationPipeline(b *testing.B) {
	const value = "  AUTH_OAuth_Provider-Name.Test  "

	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = strings.ToLower(strings.TrimSpace(value))
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = Of(value).Trim().ToLower().String()
		}
	})
}

// BenchmarkReplaceAllPipeline compares ordered fluent replacements with equivalent standard-library calls.
func BenchmarkReplaceAllPipeline(b *testing.B) {
	const value = "archive-logs cold.storage"

	b.Run("StandardLibrary", func(b *testing.B) {
		for b.Loop() {
			result := strings.ReplaceAll(value, "-", "_")
			result = strings.ReplaceAll(result, " ", "_")
			benchmarkRawStringResult = strings.ReplaceAll(result, ".", "_")
		}
	})
	b.Run("Fluent", func(b *testing.B) {
		for b.Loop() {
			benchmarkRawStringResult = Of(value).
				ReplaceAll("-", "_").
				ReplaceAll(" ", "_").
				ReplaceAll(".", "_").
				String()
		}
	})
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
