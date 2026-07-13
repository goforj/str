package str

import "testing"

var benchmarkStringResult String

// BenchmarkAppend measures a common multi-part fluent composition.
func BenchmarkAppend(b *testing.B) {
	for b.Loop() {
		benchmarkStringResult = Of("github.com").Append("/goforj", "/str", "/v2")
	}
}

// BenchmarkNormalizeSpace measures Unicode-aware whitespace normalization.
func BenchmarkNormalizeSpace(b *testing.B) {
	value := Of("  GoForj\tbuilds\n practical\u2003Go applications  ")
	for b.Loop() {
		benchmarkStringResult = value.NormalizeSpace()
	}
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
