package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestParseBenchmarkSnapshot verifies CPU-suffixed benchmark names, median timings, and stable allocation metrics.
func TestParseBenchmarkSnapshot(t *testing.T) {
	t.Parallel()

	snapshot, err := parseBenchmarkSnapshot(benchmarkFixture())
	if err != nil {
		t.Fatalf("parseBenchmarkSnapshot() error = %v", err)
	}
	if snapshot.goVersion != "go version go1.24.5 linux/arm64" {
		t.Fatalf("goVersion = %q", snapshot.goVersion)
	}
	if snapshot.command != benchmarkCommandDisplay() {
		t.Fatalf("command = %q, want %q", snapshot.command, benchmarkCommandDisplay())
	}
	if snapshot.goos != "linux" || snapshot.goarch != "arm64" {
		t.Fatalf("platform = %s/%s, want linux/arm64", snapshot.goos, snapshot.goarch)
	}
	if len(snapshot.results) != len(benchmarkDefinitions) {
		t.Fatalf("len(results) = %d, want %d", len(snapshot.results), len(benchmarkDefinitions))
	}

	first := snapshot.results[0]
	if first.name != "TrimComparison" || first.label != "Trim" {
		t.Fatalf("first result = %#v", first)
	}
	if first.standardLibrary.nanoseconds != 15.5 || first.standardLibrary.bytes != 16 || first.standardLibrary.allocations != 1 {
		t.Fatalf("standard-library measurement = %#v", first.standardLibrary)
	}
	if first.fluent.nanoseconds != 25.5 || first.fluent.bytes != 32 || first.fluent.allocations != 2 {
		t.Fatalf("fluent measurement = %#v", first.fluent)
	}
	if first.standardLibrary.samples != benchmarkSampleCount || first.fluent.samples != benchmarkSampleCount {
		t.Fatalf("sample counts = %d and %d", first.standardLibrary.samples, first.fluent.samples)
	}

	withoutSuffix := strings.ReplaceAll(string(benchmarkFixture()), "-1\t", "\t")
	if _, err := parseBenchmarkSnapshot([]byte(withoutSuffix)); err != nil {
		t.Fatalf("parseBenchmarkSnapshot() without name suffix error = %v", err)
	}
}

// TestParseBenchmarkSnapshotRejectsInvalidComparisons verifies that incomplete or unstable data cannot enter the README.
func TestParseBenchmarkSnapshotRejectsInvalidComparisons(t *testing.T) {
	t.Parallel()

	valid := string(benchmarkFixture())
	tests := []struct {
		name     string
		snapshot string
		want     string
	}{
		{
			name:     "missing pair",
			snapshot: removeLinesContaining(valid, "BenchmarkToLowerComparison/Fluent-1"),
			want:     "must contain StandardLibrary and Fluent",
		},
		{
			name:     "unequal samples",
			snapshot: strings.Replace(valid, "BenchmarkTrimComparison/Fluent-1\t1000\t30 ns/op\t32 B/op\t2 allocs/op\n", "", 1),
			want:     "sample counts differ",
		},
		{
			name:     "unstable allocations",
			snapshot: strings.Replace(valid, "BenchmarkTrimComparison/StandardLibrary-1\t1000\t20 ns/op\t16 B/op\t1 allocs/op", "BenchmarkTrimComparison/StandardLibrary-1\t1000\t20 ns/op\t17 B/op\t1 allocs/op", 1),
			want:     "allocation metrics are unstable",
		},
		{
			name:     "changed name suffix",
			snapshot: strings.Replace(valid, "BenchmarkTrimComparison/StandardLibrary-1", "BenchmarkTrimComparison/StandardLibrary-2", 1),
			want:     "name suffix changed",
		},
		{
			name: "workload order",
			snapshot: strings.Replace(
				valid,
				"BenchmarkToLowerComparison/StandardLibrary-1",
				"BenchmarkTrimComparison/StandardLibrary-1",
				1,
			),
			want: "out of order",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := parseBenchmarkSnapshot([]byte(test.snapshot))
			if err == nil {
				t.Fatal("parseBenchmarkSnapshot() error = nil, want validation error")
			}
			if !strings.Contains(err.Error(), test.want) {
				t.Fatalf("parseBenchmarkSnapshot() error = %q, want text %q", err, test.want)
			}
		})
	}
}

// TestRenderPerformance verifies absolute metrics, stable workload order, and interpretation guidance.
func TestRenderPerformance(t *testing.T) {
	t.Parallel()

	snapshot, err := parseBenchmarkSnapshot(benchmarkFixture())
	if err != nil {
		t.Fatalf("parseBenchmarkSnapshot() error = %v", err)
	}

	got := renderPerformance(snapshot)
	wantRows := []string{
		"| Trim | 15.5 ns/op · 16 B/op · 1 allocs/op | 25.5 ns/op · 32 B/op · 2 allocs/op |",
		"| ToLower | 115.5 ns/op · 32 B/op · 2 allocs/op | 125.5 ns/op · 64 B/op · 3 allocs/op |",
		"| NormalizeSpace (Fields + Join) | 215.5 ns/op · 48 B/op · 3 allocs/op | 225.5 ns/op · 96 B/op · 4 allocs/op |",
		"| Trim → ToLower | 315.5 ns/op · 64 B/op · 4 allocs/op | 325.5 ns/op · 128 B/op · 5 allocs/op |",
		"| ReplaceAll × 3 | 415.5 ns/op · 80 B/op · 5 allocs/op | 425.5 ns/op · 160 B/op · 6 allocs/op |",
	}
	lastPosition := -1
	for _, row := range wantRows {
		position := strings.Index(got, row)
		if position < 0 {
			t.Fatalf("renderPerformance() missing row %q\n%s", row, got)
		}
		if position <= lastPosition {
			t.Fatalf("renderPerformance() row %q is out of order", row)
		}
		lastPosition = position
	}
	for _, phrase := range []string{"Recorded with `go1.24.5` on `linux/arm64` using `-cpu=1` (`GOMAXPROCS=1`)", "Timing is machine-specific", "Treat small timing differences within the raw sample spread as noise", "wrapping and unwrapping added no heap allocations", "builds a field slice", "ordinary README generation only renders that frozen snapshot"} {
		if !strings.Contains(got, phrase) {
			t.Fatalf("renderPerformance() missing guidance %q", phrase)
		}
	}
}

// TestBenchmarkCommandDisplay verifies that the snapshot records a precise, copyable comparison command.
func TestBenchmarkCommandDisplay(t *testing.T) {
	t.Parallel()

	want := "go test -run '^$' -bench '^(BenchmarkTrimComparison|BenchmarkToLowerComparison|BenchmarkNormalizeSpaceComparison|BenchmarkNormalizationPipeline|BenchmarkReplaceAllPipeline)$' -benchmem -count=10 -benchtime=500ms -cpu=1 ."
	if got := benchmarkCommandDisplay(); got != want {
		t.Fatalf("benchmarkCommandDisplay() = %q, want %q", got, want)
	}
}

// TestCommittedBenchmarkSnapshot verifies that the checked-in source for the README remains complete and parseable.
func TestCommittedBenchmarkSnapshot(t *testing.T) {
	t.Parallel()

	content, err := os.ReadFile("benchmarks.txt")
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}
	snapshot, err := parseBenchmarkSnapshot(content)
	if err != nil {
		t.Fatalf("parseBenchmarkSnapshot() error = %v", err)
	}
	if len(snapshot.results) != len(benchmarkDefinitions) {
		t.Fatalf("len(results) = %d, want %d", len(snapshot.results), len(benchmarkDefinitions))
	}
	for _, result := range snapshot.results {
		if result.standardLibrary.samples != benchmarkSampleCount || result.fluent.samples != benchmarkSampleCount {
			t.Fatalf("%s sample counts = %d and %d, want %d each", result.name, result.standardLibrary.samples, result.fluent.samples, benchmarkSampleCount)
		}
	}
}

// TestAtomicWriteFile verifies that generated snapshots replace existing content without leaving temporary files.
func TestAtomicWriteFile(t *testing.T) {
	t.Parallel()

	directory := t.TempDir()
	path := filepath.Join(directory, "benchmarks.txt")
	if err := os.WriteFile(path, []byte("old"), 0o600); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}
	if err := atomicWriteFile(path, []byte("new\n"), 0o644); err != nil {
		t.Fatalf("atomicWriteFile() error = %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}
	if string(content) != "new\n" {
		t.Fatalf("content = %q, want %q", content, "new\\n")
	}
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatalf("os.ReadDir() error = %v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "benchmarks.txt" {
		t.Fatalf("directory entries = %#v", entries)
	}
}

// benchmarkFixture builds complete deterministic raw output without running performance-sensitive tests.
func benchmarkFixture() []byte {
	var output strings.Builder
	fmt.Fprintf(&output, "%s%s\n", benchmarkVersionPrefix, "go version go1.24.5 linux/arm64")
	fmt.Fprintf(&output, "%s%s\n\n", benchmarkCommandPrefix, benchmarkCommandDisplay())
	output.WriteString("goos: linux\ngoarch: arm64\npkg: github.com/goforj/str/v2\ncpu: fixture\n")

	timingOrder := []int{10, 1, 8, 3, 6, 2, 9, 4, 7, 5}
	for workloadIndex, definition := range benchmarkDefinitions {
		for implementationIndex, implementation := range benchmarkImplementations {
			base := workloadIndex*100 + implementationIndex*10 + 10
			bytesPerOperation := (workloadIndex + 1) * (implementationIndex + 1) * 16
			allocations := workloadIndex + implementationIndex + 1
			for _, timing := range timingOrder {
				fmt.Fprintf(
					&output,
					"Benchmark%s/%s-1\t1000\t%d ns/op\t%d B/op\t%d allocs/op\n",
					definition.name,
					implementation,
					base+timing,
					bytesPerOperation,
					allocations,
				)
			}
		}
	}
	output.WriteString("PASS\nok\tgithub.com/goforj/str/v2\t1.000s\n")

	return []byte(output.String())
}

// removeLinesContaining removes fixture lines matching text while preserving the rest byte-for-byte.
func removeLinesContaining(input, text string) string {
	var output strings.Builder
	for _, line := range strings.SplitAfter(input, "\n") {
		if !strings.Contains(line, text) {
			output.WriteString(line)
		}
	}
	return output.String()
}
