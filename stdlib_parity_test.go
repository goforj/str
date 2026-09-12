package str

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"iter"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// standardOperation records the standard signature, including the source argument position.
type standardOperation struct {
	name     string
	function any
	source   int
}

// standardOperations is the reviewed Go 1.27 string-function contract, excluding stateful constructors.
var standardOperations = []standardOperation{
	{"Clone", strings.Clone, 0},
	{"Compare", strings.Compare, 0},
	{"Contains", strings.Contains, 0},
	{"ContainsAny", strings.ContainsAny, 0},
	{"ContainsFunc", strings.ContainsFunc, 0},
	{"ContainsRune", strings.ContainsRune, 0},
	{"Count", strings.Count, 0},
	{"Cut", strings.Cut, 0},
	{"CutLast", cutLastStandard, 0},
	{"CutPrefix", strings.CutPrefix, 0},
	{"CutSuffix", strings.CutSuffix, 0},
	{"EqualFold", strings.EqualFold, 0},
	{"Fields", strings.Fields, 0},
	{"FieldsFunc", strings.FieldsFunc, 0},
	{"FieldsFuncSeq", strings.FieldsFuncSeq, 0},
	{"FieldsSeq", strings.FieldsSeq, 0},
	{"HasPrefix", strings.HasPrefix, 0},
	{"HasSuffix", strings.HasSuffix, 0},
	{"Index", strings.Index, 0},
	{"IndexAny", strings.IndexAny, 0},
	{"IndexByte", strings.IndexByte, 0},
	{"IndexFunc", strings.IndexFunc, 0},
	{"IndexRune", strings.IndexRune, 0},
	{"Join", strings.Join, -1},
	{"LastIndex", strings.LastIndex, 0},
	{"LastIndexAny", strings.LastIndexAny, 0},
	{"LastIndexByte", strings.LastIndexByte, 0},
	{"LastIndexFunc", strings.LastIndexFunc, 0},
	{"Lines", strings.Lines, 0},
	{"Map", strings.Map, 1},
	{"Repeat", strings.Repeat, 0},
	{"Replace", strings.Replace, 0},
	{"ReplaceAll", strings.ReplaceAll, 0},
	{"Split", strings.Split, 0},
	{"SplitAfter", strings.SplitAfter, 0},
	{"SplitAfterN", strings.SplitAfterN, 0},
	{"SplitAfterSeq", strings.SplitAfterSeq, 0},
	{"SplitN", strings.SplitN, 0},
	{"SplitSeq", strings.SplitSeq, 0},
	{"Title", strings.Title, 0},
	{"ToLower", strings.ToLower, 0},
	{"ToLowerSpecial", strings.ToLowerSpecial, 1},
	{"ToTitle", strings.ToTitle, 0},
	{"ToTitleSpecial", strings.ToTitleSpecial, 1},
	{"ToUpper", strings.ToUpper, 0},
	{"ToUpperSpecial", strings.ToUpperSpecial, 1},
	{"ToValidUTF8", strings.ToValidUTF8, 0},
	{"Trim", strings.Trim, 0},
	{"TrimFunc", strings.TrimFunc, 0},
	{"TrimLeft", strings.TrimLeft, 0},
	{"TrimLeftFunc", strings.TrimLeftFunc, 0},
	{"TrimPrefix", strings.TrimPrefix, 0},
	{"TrimRight", strings.TrimRight, 0},
	{"TrimRightFunc", strings.TrimRightFunc, 0},
	{"TrimSpace", strings.TrimSpace, 0},
	{"TrimSuffix", strings.TrimSuffix, 0},
}

// referenceCutLast permits the parity suite to compile on the Go 1.24 minimum.
// Go 1.27 and newer select strings.CutLast itself for the full differential suite.
func referenceCutLast(value, sep string) (string, string, bool) {
	index := strings.LastIndex(value, sep)
	if index < 0 {
		return value, "", false
	}
	return value[:index], value[index+len(sep):], true
}

// TestStandardAPIInventory detects missing operations and signature drift against the active toolchain.
func TestStandardAPIInventory(t *testing.T) {
	t.Parallel()
	registered := make(map[string]bool)
	wrapper := reflect.TypeOf(String{})
	raw := reflect.TypeOf("")
	for _, operation := range standardOperations {
		registered[operation.name] = true
		standard := reflect.TypeOf(operation.function)
		var actual reflect.Type
		if operation.source < 0 {
			actual = reflect.TypeOf(Join)
		} else {
			method, ok := wrapper.MethodByName(operation.name)
			if !ok {
				t.Errorf("missing standard method %s", operation.name)
				continue
			}
			actual = method.Type
		}
		if actual.NumIn() != standard.NumIn() || actual.NumOut() != standard.NumOut() || actual.IsVariadic() != standard.IsVariadic() {
			t.Errorf("%s shape: %v versus %v", operation.name, actual, standard)
			continue
		}
		var expectedInputs []reflect.Type
		if operation.source >= 0 {
			expectedInputs = append(expectedInputs, wrapper)
		}
		for i := 0; i < standard.NumIn(); i++ {
			if i != operation.source {
				expectedInputs = append(expectedInputs, standard.In(i))
			}
		}
		for i, expected := range expectedInputs {
			if actual.In(i) != expected {
				t.Errorf("%s argument %d: %v, want %v", operation.name, i, actual.In(i), expected)
			}
		}
		for i := 0; i < standard.NumOut(); i++ {
			expected := standard.Out(i)
			if standard.NumOut() == 1 && expected == raw {
				expected = wrapper
			}
			if actual.Out(i) != expected {
				t.Errorf("%s result %d: %v, want %v", operation.name, i, actual.Out(i), expected)
			}
		}
	}
	// Inspect only files selected for this toolchain, so Go's build constraints remain authoritative.
	pkg, err := build.Default.Import("strings", "", 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range pkg.GoFiles {
		file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(pkg.Dir, name), nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range file.Decls {
			fn, ok := declaration.(*ast.FuncDecl)
			if !ok || fn.Recv != nil || !fn.Name.IsExported() {
				continue
			}
			if fn.Name.Name == "NewReader" || fn.Name.Name == "NewReplacer" {
				continue
			}
			if !registered[fn.Name.Name] {
				t.Errorf("standard function %s needs an explicit parity entry", fn.Name.Name)
			}
		}
	}
	if _, ok := wrapper.MethodByName("TrimChars"); ok {
		t.Error("removed TrimChars alias remains")
	}
	if _, ok := wrapper.MethodByName("Join"); ok {
		t.Error("Join must not discard a fluent receiver")
	}
}

// parityCase varies delimiters, counts, rune predicates, and malformed UTF-8 independently of the wrapper.
type parityCase struct {
	needle      string
	replacement string
	count       int
	character   rune
	predicate   func(rune) bool
	mapping     func(rune) rune
	special     unicode.SpecialCase
	elements    []string
}

// dropDigits exercises Map's negative-rune removal contract.
func dropDigits(r rune) rune {
	if unicode.IsDigit(r) {
		return -1
	}
	return r
}

// parityInputs includes empty, multibyte, decomposed, newline, NUL, and invalid-byte sources.
var parityInputs = []string{"", "abc", "aaaa", "héllo", "é", "e\u0301clair", "hELLO wORLD", "foo_bar", "Σσς", "İIıi", "  go\t2\u2003", "a\r\nb\rc\n", "a\u0085b\u2028c\u2029d", "\n", "a\x00b", "\xffa\xfe", "a\xff\xfeb", "🦫go🦫"}

// parityCases exercises both successful operations and documented failure modes.
var parityCases = []parityCase{
	{"", "-", -1, utf8.RuneError, unicode.IsSpace, unicode.ToUpper, unicode.TurkishCase, nil},
	{"a", "", 0, 'a', unicode.IsLetter, unicode.ToLower, unicode.AzeriCase, []string{}},
	{"a", "aa", 1, 'é', unicode.IsDigit, dropDigits, nil, []string{"", "a", ""}},
	{"é", "X", 2, '🦫', unicode.IsSpace, unicode.ToTitle, unicode.TurkishCase, []string{"go", "forj"}},
	{"\n", "\r\n", 3, '\n', unicode.IsControl, dropDigits, nil, []string{"a\n", "b"}},
	{"\xff", "?", -2, -1, unicode.IsLetter, unicode.ToUpper, nil, []string{"\xff", "é"}},
	{"\xa9", "\xff", 4, unicode.MaxRune + 1, unicode.IsDigit, dropDigits, unicode.TurkishCase, []string{"é"}},
	{"missing", "new", 5, 'Σ', unicode.IsSpace, unicode.ToLower, nil, []string{""}},
	{" \t", "!", 1, 'I', nil, nil, nil, nil},
}

// operationArguments preserves source positions for Map and special-case conversions.
func operationArguments(operation standardOperation, input string, scenario parityCase) []reflect.Value {
	signature := reflect.TypeOf(operation.function)
	arguments := make([]reflect.Value, signature.NumIn())
	stringIndex := 0
	for i := range arguments {
		argumentType := signature.In(i)
		var value any
		if i == operation.source {
			value = input
		} else {
			switch argumentType {
			case reflect.TypeOf(""):
				value = scenario.needle
				if stringIndex > 0 {
					value = scenario.replacement
				}
				stringIndex++
			case reflect.TypeOf(int(0)):
				value = scenario.count
			case reflect.TypeOf(rune(0)):
				value = scenario.character
			case reflect.TypeOf(byte(0)):
				value = byte(scenario.character)
			case reflect.TypeOf([]string{}):
				value = scenario.elements
			case reflect.TypeOf(unicode.SpecialCase{}):
				value = scenario.special
			case reflect.TypeOf((func(rune) bool)(nil)):
				value = scenario.predicate
			case reflect.TypeOf((func(rune) rune)(nil)):
				value = scenario.mapping
			default:
				panic(fmt.Sprintf("unsupported parity argument: %v", argumentType))
			}
		}
		arguments[i] = reflect.ValueOf(value)
	}
	return arguments
}

// operationResult retains panic behavior as well as terminal values.
type operationResult struct {
	values     []any
	panicValue string
}

// invokeOperation materializes iterators inside recovery so lazy failures are compared too.
func invokeOperation(function reflect.Value, arguments []reflect.Value) (result operationResult) {
	defer func() {
		if value := recover(); value != nil {
			result.panicValue = fmt.Sprintf("%T: %v", value, value)
		}
	}()
	for _, value := range function.Call(arguments) {
		output := value.Interface()
		switch typed := output.(type) {
		case String:
			output = typed.String()
		case iter.Seq[string]:
			output = slices.Collect(typed)
		}
		result.values = append(result.values, output)
	}
	return result
}

// checkStandardOperation compares independently invoked fluent and standard operations.
func checkStandardOperation(t *testing.T, operation standardOperation, input string, scenario parityCase) {
	t.Helper()
	arguments := operationArguments(operation, input, scenario)
	expected := invokeOperation(reflect.ValueOf(operation.function), arguments)
	function := reflect.ValueOf(Join)
	if operation.source >= 0 {
		function = reflect.ValueOf(Of(input)).MethodByName(operation.name)
		arguments = append(arguments[:operation.source:operation.source], arguments[operation.source+1:]...)
	}
	actual := invokeOperation(function, arguments)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%s(%q; needle=%q count=%d rune=%U): got %#v, want %#v", operation.name, input, scenario.needle, scenario.count, scenario.character, actual, expected)
	}
}

// TestStandardLibraryParity checks every mirrored operation against the standard library.
func TestStandardLibraryParity(t *testing.T) {
	for _, operation := range standardOperations {
		t.Run(operation.name, func(t *testing.T) {
			t.Parallel()
			for _, input := range parityInputs {
				for _, scenario := range parityCases {
					checkStandardOperation(t, operation, input, scenario)
				}
			}
		})
	}
}

// TestStandardIterators preserves laziness, early-stop consumption, and operation-specific reuse behavior.
func TestStandardIterators(t *testing.T) {
	for _, name := range []string{"Lines", "FieldsSeq", "FieldsFuncSeq", "SplitSeq", "SplitAfterSeq"} {
		t.Run(name, func(t *testing.T) {
			var operation standardOperation
			for _, candidate := range standardOperations {
				if candidate.name == name {
					operation = candidate
				}
			}
			arguments := operationArguments(operation, "a b c\n", parityCase{needle: " ", predicate: unicode.IsSpace})
			standard := reflect.ValueOf(operation.function).Call(arguments)[0].Interface().(iter.Seq[string])
			fluent := reflect.ValueOf(Of("a b c\n")).MethodByName(name).Call(arguments[1:])[0].Interface().(iter.Seq[string])
			consume := func(sequence iter.Seq[string]) []string {
				var first []string
				sequence(func(value string) bool { first = append(first, value); return false })
				return first
			}
			if got, want := consume(fluent), consume(standard); !reflect.DeepEqual(got, want) {
				t.Fatalf("first: %q versus %q", got, want)
			}
			if got, want := slices.Collect(fluent), slices.Collect(standard); !reflect.DeepEqual(got, want) {
				t.Fatalf("resume: %q versus %q", got, want)
			}
			if got, want := slices.Collect(fluent), slices.Collect(standard); !reflect.DeepEqual(got, want) {
				t.Fatalf("reused iterator: %q versus %q", got, want)
			}
		})
	}
	calls := 0
	iterator := Of("a b").FieldsFuncSeq(func(r rune) bool { calls++; return r == ' ' })
	if calls != 0 {
		t.Fatal("FieldsFuncSeq evaluated its predicate eagerly")
	}
	_ = slices.Collect(iterator)
	if calls == 0 {
		t.Fatal("FieldsFuncSeq never evaluated its predicate")
	}
}

// TestCloneOwnership verifies the allocation guarantee that value equality alone cannot prove.
func TestCloneOwnership(t *testing.T) {
	value := strings.Repeat("gopher", 1000)[2:6]
	cloned := Of(value).Clone().String()
	if cloned != value || unsafe.StringData(cloned) == unsafe.StringData(value) {
		t.Fatal("Clone did not allocate independent backing storage")
	}
}

// TestRepeatOverflow preserves panic behavior without attempting a huge allocation.
func TestRepeatOverflow(t *testing.T) {
	for _, input := range []string{"", "ab"} {
		scenario := parityCase{count: int(^uint(0) >> 1)}
		checkStandardOperation(t, standardOperation{"Repeat", strings.Repeat, 0}, input, scenario)
	}
}

// FuzzStandardLibraryParity keeps malformed UTF-8 and empty searches in the contract.
func FuzzStandardLibraryParity(f *testing.F) {
	for _, input := range parityInputs {
		f.Add(input, "", "-", 2)
	}
	f.Fuzz(func(t *testing.T, input, needle, replacement string, count int) {
		if len(input)+len(needle)+len(replacement) > 2048 {
			t.Skip()
		}
		count %= 8
		scenario := parityCase{needle, replacement, count, utf8.RuneError, unicode.IsSpace, dropDigits, unicode.TurkishCase, []string{input, needle}}
		for _, operation := range standardOperations {
			checkStandardOperation(t, operation, input, scenario)
		}
	})
}
