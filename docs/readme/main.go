// Command readme rebuilds generated README sections from the library's GoDoc and tests.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	apiStart         = "<!-- api:embed:start -->"
	apiEnd           = "<!-- api:embed:end -->"
	performanceStart = "<!-- performance:embed:start -->"
	performanceEnd   = "<!-- performance:embed:end -->"
	testCountStart   = "<!-- test-count:embed:start -->"
	testCountEnd     = "<!-- test-count:embed:end -->"
	documentation    = "https://pkg.go.dev/github.com/goforj/str/v2"
)

var (
	groupHeader   = regexp.MustCompile(`(?im)^\s*@group\s+(.+?)\s*$`)
	exampleHeader = regexp.MustCompile(`(?i)^\s*Example:\s*(.*)$`)
)

// apiSymbol contains one public declaration and the GoDoc content rendered in the README.
type apiSymbol struct {
	name        string
	group       string
	receiver    string
	description string
	examples    []apiExample
}

// apiExample contains one runnable GoDoc example and its source position.
type apiExample struct {
	label string
	code  string
	line  int
}

// main reports generation errors without a stack trace because this command is intended for routine documentation updates.
func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "readme generator:", err)
		os.Exit(1)
	}

	fmt.Println("README.md performance table, API index, and test count updated")
}

// run computes every generated value before writing so a failed parse or test run cannot partially update README.md.
func run(arguments []string) error {
	flags := flag.NewFlagSet("readme", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	recordBenchmarks := flags.Bool("record-benchmarks", false, "record a fresh benchmark snapshot before rebuilding README.md")
	if err := flags.Parse(arguments); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(flags.Args(), " "))
	}

	root, err := findRoot()
	if err != nil {
		return err
	}

	benchmarkPath := filepath.Join(root, benchmarkSnapshotPath)
	benchmarkSnapshot, err := os.ReadFile(benchmarkPath)
	if err != nil && !*recordBenchmarks {
		return fmt.Errorf("read benchmark snapshot: %w", err)
	}
	if *recordBenchmarks {
		benchmarkSnapshot, err = recordBenchmarkSnapshot(root)
		if err != nil {
			return fmt.Errorf("record benchmark snapshot: %w", err)
		}
	}
	benchmarks, err := parseBenchmarkSnapshot(benchmarkSnapshot)
	if err != nil {
		return fmt.Errorf("parse benchmark snapshot: %w", err)
	}

	symbols, err := parseAPISymbols(root)
	if err != nil {
		return fmt.Errorf("parse API declarations: %w", err)
	}

	readmePath := filepath.Join(root, "README.md")
	readme, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("read README.md: %w", err)
	}
	if _, _, err := markerBounds(string(readme), apiStart, apiEnd, "API index"); err != nil {
		return err
	}
	if _, _, err := markerBounds(string(readme), performanceStart, performanceEnd, "performance"); err != nil {
		return err
	}
	if _, _, err := markerBounds(string(readme), testCountStart, testCountEnd, "test count"); err != nil {
		return err
	}

	tests, err := countTests(root)
	if err != nil {
		return fmt.Errorf("count tests: %w", err)
	}

	updated, err := replaceMarkedSection(
		string(readme),
		performanceStart,
		performanceEnd,
		"\n\n"+renderPerformance(benchmarks)+"\n",
		"performance",
	)
	if err != nil {
		return err
	}

	updated, err = replaceMarkedSection(
		updated,
		apiStart,
		apiEnd,
		"\n\n"+renderAPI(symbols)+"\n",
		"API index",
	)
	if err != nil {
		return err
	}

	testBadge := fmt.Sprintf("\n    <img src=\"https://img.shields.io/badge/tests-%d-brightgreen\" alt=\"Tests\">\n", tests)
	updated, err = replaceMarkedSection(
		updated,
		testCountStart,
		testCountEnd,
		testBadge,
		"test count",
	)
	if err != nil {
		return err
	}

	readmeChanged := !bytes.Equal(readme, []byte(updated))
	if !*recordBenchmarks && !readmeChanged {
		return nil
	}

	if *recordBenchmarks {
		if err := atomicWriteFile(benchmarkPath, benchmarkSnapshot, 0o644); err != nil {
			return fmt.Errorf("write benchmark snapshot: %w", err)
		}
	}
	if readmeChanged {
		if err := atomicWriteFile(readmePath, []byte(updated), 0o644); err != nil {
			return fmt.Errorf("write README.md: %w", err)
		}
	}

	return nil
}

// parseAPISymbols reads public declarations and their examples from the package root.
func parseAPISymbols(root string) ([]apiSymbol, error) {
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(
		fset,
		root,
		func(info os.FileInfo) bool {
			return !strings.HasSuffix(info.Name(), "_test.go")
		},
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	packageName, err := selectPackage(packages)
	if err != nil {
		return nil, err
	}

	pkg, ok := packages[packageName]
	if !ok {
		return nil, fmt.Errorf("selected package %q is missing", packageName)
	}

	symbolsByAnchor := make(map[string]apiSymbol)
	for _, file := range pkg.Files {
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if !ok || function.Doc == nil || !ast.IsExported(function.Name.Name) {
				continue
			}

			receiver, err := receiverName(function)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", function.Name.Name, err)
			}

			symbol := apiSymbol{
				name:        function.Name.Name,
				group:       extractGroup(function.Doc),
				receiver:    receiver,
				description: extractDescription(function.Doc),
				examples:    extractExamples(fset, function),
			}
			anchor := symbolAnchor(symbol)
			if existing, ok := symbolsByAnchor[anchor]; ok {
				existing.examples = append(existing.examples, symbol.examples...)
				symbolsByAnchor[anchor] = existing
				continue
			}
			symbolsByAnchor[anchor] = symbol
		}
	}

	symbols := make([]apiSymbol, 0, len(symbolsByAnchor))
	for _, symbol := range symbolsByAnchor {
		sort.Slice(symbol.examples, func(i, j int) bool {
			return symbol.examples[i].line < symbol.examples[j].line
		})
		symbols = append(symbols, symbol)
	}

	sort.Slice(symbols, func(i, j int) bool {
		if symbols[i].group != symbols[j].group {
			return symbols[i].group < symbols[j].group
		}
		if symbols[i].name != symbols[j].name {
			return symbols[i].name < symbols[j].name
		}
		return symbols[i].receiver < symbols[j].receiver
	})

	return symbols, nil
}

// selectPackage prefers the largest non-main package so incidental tool packages cannot displace the library package.
func selectPackage(packages map[string]*ast.Package) (string, error) {
	if len(packages) == 0 {
		return "", errors.New("no packages found in repository root")
	}

	type candidate struct {
		name      string
		fileCount int
	}

	candidates := make([]candidate, 0, len(packages))
	for name, pkg := range packages {
		candidates = append(candidates, candidate{name: name, fileCount: len(pkg.Files)})
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].fileCount != candidates[j].fileCount {
			return candidates[i].fileCount > candidates[j].fileCount
		}
		return candidates[i].name < candidates[j].name
	})

	for _, candidate := range candidates {
		if candidate.name != "main" {
			return candidate.name, nil
		}
	}

	return candidates[0].name, nil
}

// extractGroup defaults unclassified declarations to Other so newly added APIs remain visible until their documentation is categorized.
func extractGroup(group *ast.CommentGroup) string {
	match := groupHeader.FindStringSubmatch(group.Text())
	if match == nil {
		return "Other"
	}

	return strings.TrimSpace(match[1])
}

// extractDescription returns the reader-facing prose before generator metadata and examples.
func extractDescription(group *ast.CommentGroup) string {
	var lines []string
	for _, comment := range group.List {
		line := strings.TrimSpace(commentLine(comment.Text))
		if groupHeader.MatchString(line) || exampleHeader.MatchString(line) {
			break
		}
		if len(lines) == 0 && line == "" {
			continue
		}
		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

// extractExamples returns every runnable example block attached to function.
func extractExamples(fset *token.FileSet, function *ast.FuncDecl) []apiExample {
	var examples []apiExample
	var collected []string
	var label string
	var line int
	inExample := false

	flush := func() {
		if len(collected) == 0 {
			return
		}
		examples = append(examples, apiExample{
			label: label,
			code:  strings.Join(normalizeIndent(collected), "\n"),
			line:  line,
		})
		collected = nil
		label = ""
		inExample = false
	}

	for _, comment := range function.Doc.List {
		raw := commentLine(comment.Text)
		if match := exampleHeader.FindStringSubmatch(strings.TrimSpace(raw)); match != nil {
			flush()
			inExample = true
			label = strings.TrimSpace(match[1])
			line = fset.Position(comment.Slash).Line
			continue
		}
		if inExample {
			collected = append(collected, raw)
		}
	}
	flush()

	return examples
}

// commentLine removes comment syntax while retaining code indentation.
func commentLine(text string) string {
	line := strings.TrimPrefix(text, "//")
	if strings.HasPrefix(line, " ") {
		line = line[1:]
	}
	return line
}

// normalizeIndent removes shared indentation without altering nested example code.
func normalizeIndent(lines []string) []string {
	minimum := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		indent := len(line) - len(strings.TrimLeft(line, " \t"))
		if minimum == -1 || indent < minimum {
			minimum = indent
		}
	}
	if minimum <= 0 {
		return lines
	}

	normalized := make([]string, len(lines))
	for i, line := range lines {
		if len(line) >= minimum {
			normalized[i] = line[minimum:]
			continue
		}
		normalized[i] = strings.TrimLeft(line, " \t")
	}
	return normalized
}

// receiverName returns an empty name for package functions and the documented type name for methods.
func receiverName(function *ast.FuncDecl) (string, error) {
	if function.Recv == nil || len(function.Recv.List) == 0 {
		return "", nil
	}

	name := receiverTypeName(function.Recv.List[0].Type)
	if name == "" {
		return "", errors.New("unsupported receiver type in exported method")
	}

	return name, nil
}

// receiverTypeName unwraps pointer and generic syntax because pkg.go.dev anchors use the declared receiver type name.
func receiverTypeName(expression ast.Expr) string {
	switch expression := expression.(type) {
	case *ast.Ident:
		return expression.Name
	case *ast.StarExpr:
		return receiverTypeName(expression.X)
	case *ast.IndexExpr:
		return receiverTypeName(expression.X)
	case *ast.IndexListExpr:
		return receiverTypeName(expression.X)
	default:
		return ""
	}
}

// renderAPI groups symbols into a compact index followed by their generated GoDoc examples.
func renderAPI(symbols []apiSymbol) string {
	var output strings.Builder
	output.WriteString("## API index\n\n")
	output.WriteString("The full API and these examples are also available on [pkg.go.dev](")
	output.WriteString(documentation)
	output.WriteString(").\n\n")
	output.WriteString("| Group | API |\n")
	output.WriteString("| --- | --- |\n")

	for start := 0; start < len(symbols); {
		end := start + 1
		for end < len(symbols) && symbols[end].group == symbols[start].group {
			end++
		}

		links := make([]string, 0, end-start)
		for _, symbol := range symbols[start:end] {
			links = append(links, fmt.Sprintf("[%s](#%s)", symbol.name, readmeAnchor(symbol)))
		}

		fmt.Fprintf(&output, "| %s | %s |\n", symbols[start].group, strings.Join(links, " · "))
		start = end
	}

	output.WriteString("\n## API examples\n\n")
	output.WriteString("These examples come from GoDoc and run as part of the test suite.\n\n")
	for start := 0; start < len(symbols); {
		end := start + 1
		for end < len(symbols) && symbols[end].group == symbols[start].group {
			end++
		}

		output.WriteString("### " + symbols[start].group + "\n\n")
		for _, symbol := range symbols[start:end] {
			fmt.Fprintf(&output, "#### <a id=\"%s\"></a>%s\n\n", readmeAnchor(symbol), symbol.name)
			if symbol.description != "" {
				output.WriteString(symbol.description + "\n\n")
			}
			for _, example := range symbol.examples {
				if example.label != "" && len(symbol.examples) > 1 {
					output.WriteString("_Example: " + example.label + "_\n\n")
				}
				output.WriteString("```go\n")
				output.WriteString(strings.TrimSpace(example.code))
				output.WriteString("\n```\n\n")
			}
		}
		start = end
	}

	return strings.TrimRight(output.String(), "\n")
}

// readmeAnchor returns a stable local anchor for a public API name.
func readmeAnchor(symbol apiSymbol) string {
	return strings.ToLower(symbol.name)
}

// symbolAnchor mirrors pkg.go.dev's receiver-qualified anchors for methods while leaving package functions unqualified.
func symbolAnchor(symbol apiSymbol) string {
	if symbol.receiver == "" {
		return symbol.name
	}

	return symbol.receiver + "." + symbol.name
}

// replaceMarkedSection rejects absent, repeated, or reversed markers so malformed README structure is never guessed at.
func replaceMarkedSection(document, startMarker, endMarker, replacement, section string) (string, error) {
	start, end, err := markerBounds(document, startMarker, endMarker, section)
	if err != nil {
		return "", err
	}

	return document[:start] + replacement + document[end:], nil
}

// markerBounds returns content boundaries only when a section has one correctly ordered pair of markers.
func markerBounds(document, startMarker, endMarker, section string) (int, int, error) {
	startCount := strings.Count(document, startMarker)
	if startCount == 0 {
		return 0, 0, fmt.Errorf("README %s start marker %q is missing", section, startMarker)
	}
	if startCount > 1 {
		return 0, 0, fmt.Errorf("README %s start marker %q appears %d times; expected once", section, startMarker, startCount)
	}

	endCount := strings.Count(document, endMarker)
	if endCount == 0 {
		return 0, 0, fmt.Errorf("README %s end marker %q is missing", section, endMarker)
	}
	if endCount > 1 {
		return 0, 0, fmt.Errorf("README %s end marker %q appears %d times; expected once", section, endMarker, endCount)
	}

	start := strings.Index(document, startMarker) + len(startMarker)
	end := strings.Index(document, endMarker)
	if end < start {
		return 0, 0, fmt.Errorf("README %s markers are malformed: end marker precedes start marker", section)
	}

	return start, end, nil
}

// countTests includes library and documentation tests while respecting their module boundaries.
func countTests(root string) (int, error) {
	libraryTests, err := countTestsIn(root)
	if err != nil {
		return 0, err
	}
	documentationTests, err := countTestsIn(filepath.Join(root, "docs"))
	if err != nil {
		return 0, err
	}

	return libraryTests + documentationTests, nil
}

// countTestsIn uses one module's JSON event stream so subtests and top-level tests are counted consistently.
func countTestsIn(directory string) (int, error) {
	command := exec.Command("go", "test", "./...", "-run", "Test", "-count=1", "-json")
	command.Dir = directory

	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	if err := command.Run(); err != nil {
		return 0, fmt.Errorf("go test -json in %s failed: %w\n%s", directory, err, output.String())
	}

	decoder := json.NewDecoder(bytes.NewReader(output.Bytes()))
	total := 0
	for {
		var event struct {
			Action string `json:"Action"`
			Test   string `json:"Test"`
		}
		if err := decoder.Decode(&event); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, fmt.Errorf("decode go test event: %w", err)
		}
		if event.Action == "run" && event.Test != "" {
			total++
		}
	}

	return total, nil
}

// findRoot skips nested module boundaries because README generation always targets the parent library module.
func findRoot() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	for candidate := workingDirectory; ; candidate = filepath.Dir(candidate) {
		if fileExists(filepath.Join(candidate, "go.mod")) && fileExists(filepath.Join(candidate, "string.go")) {
			return filepath.Clean(candidate), nil
		}
		parent := filepath.Dir(candidate)
		if parent == candidate {
			break
		}
	}

	return "", errors.New("could not find library module root")
}

// fileExists treats inaccessible paths as absent because callers only use it to probe root candidates.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
