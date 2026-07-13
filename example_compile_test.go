package str

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestGeneratedExamplesMatchDocumentedOutput guards its covered contract against regressions.
func TestGeneratedExamplesMatchDocumentedOutput(t *testing.T) {
	t.Parallel()

	root, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot determine module root: %v", err)
	}

	modulePath, err := readModulePath(root)
	if err != nil {
		t.Fatal(err)
	}

	examplesDir := "examples"
	entries, err := os.ReadDir(filepath.Join(root, examplesDir))
	if err != nil {
		t.Fatalf("cannot read examples directory: %v", err)
	}

	exampleCount := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		exampleCount++
		name := entry.Name()
		path := filepath.Join(examplesDir, name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			sourcePath := filepath.Join(root, path, "main.go")
			source, err := os.ReadFile(sourcePath)
			if err != nil {
				t.Fatalf("read generated example: %v", err)
			}

			if err := requireLocalModuleImport(sourcePath, source, modulePath); err != nil {
				t.Fatal(err)
			}

			expected, err := documentedOutput(source)
			if err != nil {
				t.Fatalf("parse documented output: %v", err)
			}

			actual, err := runExampleWithoutBuildTags(root, path, source)
			if err != nil {
				t.Fatalf("example %q failed to execute:\n%v", name, err)
			}

			if !bytes.Equal(actual, expected) {
				t.Errorf(
					"example %q output mismatch:\nexpected: %q\nactual:   %q",
					name,
					expected,
					actual,
				)
			}
		})
	}

	if exampleCount == 0 {
		t.Fatal("no generated example directories found")
	}
}

// TestDocumentedOutput guards its covered contract against regressions.
func TestDocumentedOutput(t *testing.T) {
	t.Parallel()

	source := []byte(`package main

func main() {
	println(42, true)
	// #int 42
	// #bool true

	println("first\nsecond")
	// #string first\nsecond

	println("  padded ")
` + "\t// #string   padded \n}\n")

	got, err := documentedOutput(source)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte("42 true\nfirst\nsecond\n  padded \n")
	if !bytes.Equal(got, want) {
		t.Errorf("documented output mismatch:\nexpected: %q\nactual:   %q", want, got)
	}
}

// TestStripBuildTagsPreservesFileHeader guards its covered contract against regressions.
func TestStripBuildTagsPreservesFileHeader(t *testing.T) {
	t.Parallel()

	source := []byte(`// Copyright Example
//go:build ignore
// +build ignore

// Package documentation remains intact.
package main

func main() {}
`)

	got := stripBuildTags(source)
	if bytes.Contains(got, []byte("//go:build")) || bytes.Contains(got, []byte("// +build")) {
		t.Errorf("build constraints remain in stripped source:\n%s", got)
	}
	if !bytes.Contains(got, []byte("// Copyright Example")) {
		t.Errorf("leading file header was removed:\n%s", got)
	}
	if !bytes.Contains(got, []byte("// Package documentation remains intact.")) {
		t.Errorf("package documentation was removed:\n%s", got)
	}
}

// readModulePath returns the module path declared by the root go.mod file.
func readModulePath(root string) (string, error) {
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", fmt.Errorf("read go.mod: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 || fields[0] != "module" {
			continue
		}

		if strings.HasPrefix(fields[1], `"`) {
			path, err := strconv.Unquote(fields[1])
			if err != nil {
				return "", fmt.Errorf("parse module path: %w", err)
			}
			return path, nil
		}

		return fields[1], nil
	}

	return "", fmt.Errorf("module path not found in go.mod")
}

// requireLocalModuleImport verifies that the example imports the module under test.
func requireLocalModuleImport(filename string, src []byte, modulePath string) error {
	file, err := parser.ParseFile(token.NewFileSet(), filename, src, parser.ImportsOnly)
	if err != nil {
		return fmt.Errorf("parse generated example imports: %w", err)
	}

	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return fmt.Errorf("parse generated example import %s: %w", spec.Path.Value, err)
		}
		if path == modulePath {
			return nil
		}
	}

	return fmt.Errorf("generated example does not import local module %q", modulePath)
}

// documentedOutput converts contiguous // #type value comments into output lines.
func documentedOutput(src []byte) ([]byte, error) {
	normalized := strings.ReplaceAll(string(src), "\r\n", "\n")
	var outputLines []string
	var values []string

	flush := func() {
		if len(values) == 0 {
			return
		}
		outputLines = append(outputLines, strings.Join(values, " "))
		values = nil
	}

	for lineNumber, line := range strings.Split(normalized, "\n") {
		comment := strings.TrimLeft(line, " \t")
		if !strings.HasPrefix(comment, "//") {
			flush()
			continue
		}

		comment = strings.TrimLeft(strings.TrimPrefix(comment, "//"), " \t")
		if !strings.HasPrefix(comment, "#") {
			flush()
			continue
		}

		separator := strings.IndexAny(comment, " \t")
		if separator < 2 {
			return nil, fmt.Errorf("line %d has malformed expected output %q", lineNumber+1, comment)
		}

		value, err := unescapeExpectedValue(comment[separator+1:])
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNumber+1, err)
		}
		values = append(values, value)
	}
	flush()

	if len(outputLines) == 0 {
		return nil, fmt.Errorf("no // #type expected output lines found")
	}

	return []byte(strings.Join(outputLines, "\n") + "\n"), nil
}

// unescapeExpectedValue decodes Go escape sequences used to document multiline and Unicode output.
func unescapeExpectedValue(value string) (string, error) {
	var decoded strings.Builder
	for value != "" {
		if value[0] == '"' {
			decoded.WriteByte(value[0])
			value = value[1:]
			continue
		}

		r, _, tail, err := strconv.UnquoteChar(value, '"')
		if err != nil {
			return "", fmt.Errorf("decode expected output %q: %w", value, err)
		}
		decoded.WriteRune(r)
		value = tail
	}

	return decoded.String(), nil
}

// runExampleWithoutBuildTags executes an ignored example through a temporary overlay.
func runExampleWithoutBuildTags(root, exampleDir string, src []byte) ([]byte, error) {
	orig := filepath.Join(root, exampleDir, "main.go")
	clean := stripBuildTags(src)

	tmpDir, err := os.MkdirTemp("", "example-overlay-*")
	if err != nil {
		return nil, fmt.Errorf("create example overlay directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(tmpFile, clean, 0o644); err != nil {
		return nil, fmt.Errorf("write example overlay source: %w", err)
	}

	overlay := map[string]any{
		"Replace": map[string]string{
			orig: tmpFile,
		},
	}

	overlayJSON, err := json.Marshal(overlay)
	if err != nil {
		return nil, fmt.Errorf("encode example overlay: %w", err)
	}

	overlayPath := filepath.Join(tmpDir, "overlay.json")
	if err := os.WriteFile(overlayPath, overlayJSON, 0o644); err != nil {
		return nil, fmt.Errorf("write example overlay: %w", err)
	}

	cmd := exec.Command(
		"go", "run",
		"-overlay", overlayPath,
		"./"+filepath.ToSlash(exampleDir),
	)
	cmd.Dir = root

	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("go run: %w\n%s", err, output)
	}

	return output, nil
}

// stripBuildTags removes modern and legacy build constraints from the file header.
func stripBuildTags(src []byte) []byte {
	lines := strings.Split(string(src), "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "package ") {
			break
		}
		if isBuildConstraintLine(trimmed) {
			lines[i] = ""
		}
	}

	return []byte(strings.Join(lines, "\n"))
}

// isBuildConstraintLine reports whether a header line contains a modern or legacy build constraint.
func isBuildConstraintLine(line string) bool {
	for _, prefix := range []string{"//go:build", "// +build"} {
		if line == prefix {
			return true
		}
		if len(line) > len(prefix) && strings.HasPrefix(line, prefix) {
			next := line[len(prefix)]
			if next == ' ' || next == '\t' {
				return true
			}
		}
	}

	return false
}
