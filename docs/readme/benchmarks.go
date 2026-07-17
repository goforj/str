package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	benchmarkSnapshotPath  = "docs/readme/benchmarks.txt"
	benchmarkSampleCount   = 10
	benchmarkVersionPrefix = "# Go version: "
	benchmarkCommandPrefix = "# Command: "
)

var (
	benchmarkDefinitions = []benchmarkDefinition{
		{name: "TrimComparison", label: "Trim"},
		{name: "ToLowerComparison", label: "ToLower"},
		{name: "NormalizeSpaceComparison", label: "NormalizeSpace (Fields + Join)"},
		{name: "NormalizationPipeline", label: "Trim → ToLower"},
		{name: "ReplaceAllPipeline", label: "ReplaceAll × 3"},
	}
	benchmarkImplementations = []string{"StandardLibrary", "Fluent"}
	benchmarkNamePattern     = regexp.MustCompile(`^Benchmark([^/]+)/(StandardLibrary|Fluent)(?:-([0-9]+))?$`)
)

// benchmarkDefinition fixes the reader-facing order and label for a comparison workload.
type benchmarkDefinition struct {
	name  string
	label string
}

// benchmarkSnapshot contains validated metadata and results from one benchmark recording.
type benchmarkSnapshot struct {
	goVersion string
	command   string
	goos      string
	goarch    string
	results   []benchmarkResult
}

// benchmarkResult contains one workload's standard-library and fluent measurements.
type benchmarkResult struct {
	name            string
	label           string
	standardLibrary benchmarkMeasurement
	fluent          benchmarkMeasurement
}

// benchmarkMeasurement contains the median timing and stable allocation metrics for one implementation.
type benchmarkMeasurement struct {
	nanoseconds float64
	bytes       uint64
	allocations uint64
	samples     int
}

// benchmarkSample contains one raw benchmark output line before repeated samples are summarized.
type benchmarkSample struct {
	nanoseconds float64
	bytes       uint64
	allocations uint64
	cpuSuffix   string
}

// benchmarkCommandArgs returns the exact argument vector used to record comparisons.
func benchmarkCommandArgs() []string {
	names := make([]string, 0, len(benchmarkDefinitions))
	for _, definition := range benchmarkDefinitions {
		names = append(names, "Benchmark"+definition.name)
	}

	return []string{
		"test",
		"-run", "^$",
		"-bench", "^(" + strings.Join(names, "|") + ")$",
		"-benchmem",
		"-count=" + strconv.Itoa(benchmarkSampleCount),
		"-benchtime=500ms",
		"-cpu=1",
		".",
	}
}

// benchmarkCommandDisplay returns a copyable shell representation of the recording command.
func benchmarkCommandDisplay() string {
	arguments := benchmarkCommandArgs()
	quoted := make([]string, 0, len(arguments)+1)
	quoted = append(quoted, "go")
	for _, argument := range arguments {
		if strings.ContainsAny(argument, "$()|*?[]{} ") {
			quoted = append(quoted, "'"+strings.ReplaceAll(argument, "'", `'\''`)+"'")
			continue
		}
		quoted = append(quoted, argument)
	}

	return strings.Join(quoted, " ")
}

// recordBenchmarkSnapshot runs only the documented comparisons and combines their raw output with reproduction metadata.
func recordBenchmarkSnapshot(root string) ([]byte, error) {
	versionCommand := exec.Command("go", "version")
	versionOutput, err := versionCommand.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("go version failed: %w\n%s", err, versionOutput)
	}

	arguments := benchmarkCommandArgs()
	command := exec.Command("go", arguments...)
	command.Dir = root
	output, err := command.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %w\n%s", benchmarkCommandDisplay(), err, output)
	}

	var snapshot bytes.Buffer
	fmt.Fprintf(&snapshot, "%s%s\n", benchmarkVersionPrefix, strings.TrimSpace(string(versionOutput)))
	fmt.Fprintf(&snapshot, "%s%s\n\n", benchmarkCommandPrefix, benchmarkCommandDisplay())
	snapshot.Write(bytes.TrimSpace(output))
	snapshot.WriteByte('\n')

	if _, err := parseBenchmarkSnapshot(snapshot.Bytes()); err != nil {
		return nil, fmt.Errorf("validate recorded output: %w", err)
	}

	return snapshot.Bytes(), nil
}

// parseBenchmarkSnapshot validates metadata, ordering, samples, and metrics before producing README results.
func parseBenchmarkSnapshot(snapshot []byte) (benchmarkSnapshot, error) {
	metadata, output, err := splitBenchmarkSnapshot(snapshot)
	if err != nil {
		return benchmarkSnapshot{}, err
	}

	goVersion := strings.TrimPrefix(metadata[0], benchmarkVersionPrefix)
	if !strings.HasPrefix(goVersion, "go version go") {
		return benchmarkSnapshot{}, fmt.Errorf("invalid Go version metadata %q", goVersion)
	}

	command := strings.TrimPrefix(metadata[1], benchmarkCommandPrefix)
	if command != benchmarkCommandDisplay() {
		return benchmarkSnapshot{}, fmt.Errorf("benchmark command = %q, want %q", command, benchmarkCommandDisplay())
	}

	samples := make(map[string]map[string][]benchmarkSample, len(benchmarkDefinitions))
	for _, definition := range benchmarkDefinitions {
		samples[definition.name] = make(map[string][]benchmarkSample, len(benchmarkImplementations))
	}

	goos, goarch, err := parseBenchmarkEnvironment(output)
	if err != nil {
		return benchmarkSnapshot{}, err
	}

	lastOrder := -1
	cpuSuffix := ""
	suffixInitialized := false
	for _, line := range strings.Split(output, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 || !strings.HasPrefix(fields[0], "Benchmark") {
			continue
		}

		name, implementation, sample, err := parseBenchmarkLine(fields)
		if err != nil {
			return benchmarkSnapshot{}, err
		}
		order, ok := benchmarkOrder(name, implementation)
		if !ok {
			return benchmarkSnapshot{}, fmt.Errorf("unexpected benchmark %s/%s", name, implementation)
		}
		if order < lastOrder {
			return benchmarkSnapshot{}, fmt.Errorf("benchmark %s/%s is out of order", name, implementation)
		}
		lastOrder = order

		if !suffixInitialized {
			cpuSuffix = sample.cpuSuffix
			suffixInitialized = true
		}
		if sample.cpuSuffix != cpuSuffix {
			return benchmarkSnapshot{}, fmt.Errorf("benchmark name suffix changed from %q to %q", cpuSuffix, sample.cpuSuffix)
		}
		samples[name][implementation] = append(samples[name][implementation], sample)
	}

	if !suffixInitialized {
		return benchmarkSnapshot{}, errors.New("snapshot contains no comparison benchmarks")
	}

	results := make([]benchmarkResult, 0, len(benchmarkDefinitions))
	for _, definition := range benchmarkDefinitions {
		standardSamples := samples[definition.name][benchmarkImplementations[0]]
		fluentSamples := samples[definition.name][benchmarkImplementations[1]]
		if len(standardSamples) == 0 || len(fluentSamples) == 0 {
			return benchmarkSnapshot{}, fmt.Errorf("benchmark %s must contain StandardLibrary and Fluent samples", definition.name)
		}
		if len(standardSamples) != len(fluentSamples) {
			return benchmarkSnapshot{}, fmt.Errorf("benchmark %s sample counts differ: StandardLibrary=%d Fluent=%d", definition.name, len(standardSamples), len(fluentSamples))
		}
		if len(standardSamples) != benchmarkSampleCount {
			return benchmarkSnapshot{}, fmt.Errorf("benchmark %s has %d samples per implementation, want %d", definition.name, len(standardSamples), benchmarkSampleCount)
		}

		standardLibrary, err := summarizeBenchmarkSamples(definition.name, benchmarkImplementations[0], standardSamples)
		if err != nil {
			return benchmarkSnapshot{}, err
		}
		fluent, err := summarizeBenchmarkSamples(definition.name, benchmarkImplementations[1], fluentSamples)
		if err != nil {
			return benchmarkSnapshot{}, err
		}

		results = append(results, benchmarkResult{
			name:            definition.name,
			label:           definition.label,
			standardLibrary: standardLibrary,
			fluent:          fluent,
		})
	}

	return benchmarkSnapshot{
		goVersion: goVersion,
		command:   command,
		goos:      goos,
		goarch:    goarch,
		results:   results,
	}, nil
}

// parseBenchmarkEnvironment extracts the target operating system and architecture from raw go test output.
func parseBenchmarkEnvironment(output string) (string, string, error) {
	values := map[string]string{}
	for _, line := range strings.Split(output, "\n") {
		for _, key := range []string{"goos", "goarch"} {
			prefix := key + ": "
			if !strings.HasPrefix(line, prefix) {
				continue
			}
			if values[key] != "" {
				return "", "", fmt.Errorf("benchmark output contains repeated %s metadata", key)
			}
			values[key] = strings.TrimSpace(strings.TrimPrefix(line, prefix))
		}
	}
	if values["goos"] == "" || values["goarch"] == "" {
		return "", "", errors.New("benchmark output must contain goos and goarch metadata")
	}

	return values["goos"], values["goarch"], nil
}

// splitBenchmarkSnapshot separates the two required metadata lines from raw go test output.
func splitBenchmarkSnapshot(snapshot []byte) ([2]string, string, error) {
	normalized := strings.ReplaceAll(string(snapshot), "\r\n", "\n")
	parts := strings.SplitN(normalized, "\n\n", 2)
	if len(parts) != 2 {
		return [2]string{}, "", errors.New("snapshot must separate metadata from benchmark output with a blank line")
	}

	metadata := strings.Split(parts[0], "\n")
	if len(metadata) != 2 {
		return [2]string{}, "", fmt.Errorf("snapshot has %d metadata lines, want 2", len(metadata))
	}
	if !strings.HasPrefix(metadata[0], benchmarkVersionPrefix) {
		return [2]string{}, "", fmt.Errorf("snapshot first line must start with %q", benchmarkVersionPrefix)
	}
	if !strings.HasPrefix(metadata[1], benchmarkCommandPrefix) {
		return [2]string{}, "", fmt.Errorf("snapshot second line must start with %q", benchmarkCommandPrefix)
	}
	if strings.TrimSpace(parts[1]) == "" {
		return [2]string{}, "", errors.New("snapshot benchmark output is empty")
	}

	return [2]string{metadata[0], metadata[1]}, parts[1], nil
}

// parseBenchmarkLine extracts the comparison name, implementation, CPU suffix, and benchmark metrics.
func parseBenchmarkLine(fields []string) (string, string, benchmarkSample, error) {
	match := benchmarkNamePattern.FindStringSubmatch(fields[0])
	if match == nil {
		return "", "", benchmarkSample{}, fmt.Errorf("invalid benchmark name %q", fields[0])
	}

	if match[3] != "" {
		cpuSuffix, err := strconv.Atoi(match[3])
		if err != nil || cpuSuffix <= 0 {
			return "", "", benchmarkSample{}, fmt.Errorf("invalid benchmark name suffix in %q", fields[0])
		}
	}
	nanoseconds, err := parseFloatMetric(fields, "ns/op")
	if err != nil {
		return "", "", benchmarkSample{}, fmt.Errorf("%s: %w", fields[0], err)
	}
	bytesPerOperation, err := parseUintMetric(fields, "B/op")
	if err != nil {
		return "", "", benchmarkSample{}, fmt.Errorf("%s: %w", fields[0], err)
	}
	allocations, err := parseUintMetric(fields, "allocs/op")
	if err != nil {
		return "", "", benchmarkSample{}, fmt.Errorf("%s: %w", fields[0], err)
	}

	return match[1], match[2], benchmarkSample{
		nanoseconds: nanoseconds,
		bytes:       bytesPerOperation,
		allocations: allocations,
		cpuSuffix:   match[3],
	}, nil
}

// parseFloatMetric finds a finite non-negative floating-point metric immediately before its unit.
func parseFloatMetric(fields []string, unit string) (float64, error) {
	value, err := metricValue(fields, unit)
	if err != nil {
		return 0, err
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil || math.IsInf(parsed, 0) || math.IsNaN(parsed) || parsed < 0 {
		return 0, fmt.Errorf("invalid %s value %q", unit, value)
	}
	return parsed, nil
}

// parseUintMetric finds an unsigned integer metric immediately before its unit.
func parseUintMetric(fields []string, unit string) (uint64, error) {
	value, err := metricValue(fields, unit)
	if err != nil {
		return 0, err
	}

	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s value %q", unit, value)
	}
	return parsed, nil
}

// metricValue returns the field immediately before a requested benchmark unit.
func metricValue(fields []string, unit string) (string, error) {
	for index, field := range fields {
		if field != unit {
			continue
		}
		if index == 0 {
			break
		}
		return fields[index-1], nil
	}
	return "", fmt.Errorf("missing %s metric", unit)
}

// benchmarkOrder maps expected workload and implementation names to their stable output position.
func benchmarkOrder(name, implementation string) (int, bool) {
	for workloadIndex, definition := range benchmarkDefinitions {
		if definition.name != name {
			continue
		}
		for implementationIndex, expected := range benchmarkImplementations {
			if expected == implementation {
				return workloadIndex*len(benchmarkImplementations) + implementationIndex, true
			}
		}
	}
	return 0, false
}

// summarizeBenchmarkSamples uses the median timing while requiring allocation metrics to remain deterministic.
func summarizeBenchmarkSamples(name, implementation string, samples []benchmarkSample) (benchmarkMeasurement, error) {
	timings := make([]float64, len(samples))
	bytesPerOperation := samples[0].bytes
	allocations := samples[0].allocations
	for index, sample := range samples {
		if sample.bytes != bytesPerOperation || sample.allocations != allocations {
			return benchmarkMeasurement{}, fmt.Errorf("benchmark %s/%s allocation metrics are unstable", name, implementation)
		}
		timings[index] = sample.nanoseconds
	}
	sort.Float64s(timings)

	median := timings[len(timings)/2]
	if len(timings)%2 == 0 {
		median = (timings[len(timings)/2-1] + median) / 2
	}

	return benchmarkMeasurement{
		nanoseconds: median,
		bytes:       bytesPerOperation,
		allocations: allocations,
		samples:     len(samples),
	}, nil
}

// renderPerformance presents absolute benchmark costs without implying that timings transfer across machines.
func renderPerformance(snapshot benchmarkSnapshot) string {
	var output strings.Builder
	output.WriteString("## Performance\n\n")
	output.WriteString("These comparisons measure equivalent standard-library and `str` operations. Each cell reports the median of ")
	output.WriteString(strconv.Itoa(benchmarkSampleCount))
	output.WriteString(" samples as `ns/op · B/op · allocs/op`.\n\n")
	goVersionFields := strings.Fields(snapshot.goVersion)
	fmt.Fprintf(
		&output,
		"Recorded with `%s` on `%s/%s` using `-cpu=1` (`GOMAXPROCS=1`).\n\n",
		goVersionFields[2],
		snapshot.goos,
		snapshot.goarch,
	)
	output.WriteString("| Workload | Standard library | `str` chain |\n")
	output.WriteString("| --- | ---: | ---: |\n")
	for _, result := range snapshot.results {
		fmt.Fprintf(
			&output,
			"| %s | %s | %s |\n",
			result.label,
			formatBenchmarkMeasurement(result.standardLibrary),
			formatBenchmarkMeasurement(result.fluent),
		)
	}

	output.WriteString("\nTiming is machine-specific; use it to understand the scale of these operations, not as a universal speed claim. Treat small timing differences within the raw sample spread as noise. Allocation counts are less sensitive to machine speed and show how much heap work each composition performs. In these workloads, wrapping and unwrapping added no heap allocations; allocations came from transformations that produced new text. `NormalizeSpace` is algorithmically different: the standard-library composition builds a field slice before joining it, while `str` uses one builder pass.\n\n")
	output.WriteString("The [benchmark source](string_benchmark_test.go) and [committed raw output](docs/readme/benchmarks.txt) record exactly what ran, including the Go version and command. Refresh the measurements explicitly with `go -C docs run ./readme -record-benchmarks`; ordinary README generation only renders that frozen snapshot.")

	return output.String()
}

// formatBenchmarkMeasurement formats one absolute result at stable one-decimal timing precision.
func formatBenchmarkMeasurement(measurement benchmarkMeasurement) string {
	return fmt.Sprintf(
		"%s ns/op · %d B/op · %d allocs/op",
		strconv.FormatFloat(measurement.nanoseconds, 'f', 1, 64),
		measurement.bytes,
		measurement.allocations,
	)
}

// atomicWriteFile replaces a generated file only after its complete content is durable in the same directory.
func atomicWriteFile(path string, content []byte, mode os.FileMode) error {
	temporary, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+"-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer func() {
		_ = temporary.Close()
		_ = os.Remove(temporaryPath)
	}()

	if err := temporary.Chmod(mode); err != nil {
		return err
	}
	if _, err := temporary.Write(content); err != nil {
		return err
	}
	if err := temporary.Sync(); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return err
	}

	return nil
}
