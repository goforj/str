package str

import (
	"fmt"
	"iter"
	"reflect"
	"slices"
	"testing"
	"unicode"
)

// TestFieldIteratorReuse distinguishes reusable field sequences from single-use split sequences.
func TestFieldIteratorReuse(t *testing.T) {
	for name, sequence := range map[string]iter.Seq[string]{
		"FieldsSeq":     Of(" a\tb\u2003c ").FieldsSeq(),
		"FieldsFuncSeq": Of(" a\tb\u2003c ").FieldsFuncSeq(unicode.IsSpace),
	} {
		t.Run(name, func(t *testing.T) {
			sequence(func(field string) bool {
				if field != "a" {
					t.Fatalf("first field = %q, want a", field)
				}
				return false
			})
			for pass := range 2 {
				if got := slices.Collect(sequence); !slices.Equal(got, []string{"a", "b", "c"}) {
					t.Fatalf("pass %d = %q, want [a b c]", pass, got)
				}
			}
		})
	}
}

// TestStandardCallbackBehavior compares callback invocation and panic propagation with fresh callback state.
func TestStandardCallbackBehavior(t *testing.T) {
	for _, operation := range standardOperations {
		signature := reflect.TypeOf(operation.function)
		callback := false
		for i := range signature.NumIn() {
			callback = callback || signature.In(i).Kind() == reflect.Func
		}
		if !callback {
			continue
		}
		t.Run(operation.name, func(t *testing.T) {
			for _, input := range []string{"", " aéb\xffc ", "bbb"} {
				for _, fail := range []bool{false, true} {
					run := func(fluent bool) (operationResult, []rune) {
						var calls []rune
						record := func(r rune) {
							calls = append(calls, r)
							if fail && r == 'b' {
								panic("callback failure")
							}
						}
						scenario := parityCase{
							predicate: func(r rune) bool { record(r); return unicode.IsSpace(r) },
							mapping:   func(r rune) rune { record(r); return unicode.ToUpper(r) },
						}
						arguments := operationArguments(operation, input, scenario)
						function := reflect.ValueOf(operation.function)
						if fluent {
							function = reflect.ValueOf(Of(input)).MethodByName(operation.name)
							arguments = append(arguments[:operation.source:operation.source], arguments[operation.source+1:]...)
						}
						return invokeOperation(function, arguments), calls
					}
					want, wantCalls := run(false)
					got, gotCalls := run(true)
					if !reflect.DeepEqual(got, want) || !slices.Equal(gotCalls, wantCalls) {
						t.Fatalf("input=%q panic=%t: result=%#v calls=%U, want result=%#v calls=%U", input, fail, got, gotCalls, want, wantCalls)
					}
				}
			}
		})
	}
}

// TestEmptyReplacementEncoding records the different boundaries of ReplaceAll and Replacer explicitly.
func TestEmptyReplacementEncoding(t *testing.T) {
	value := Of("é")
	if got := value.ReplaceAll("", "-").String(); got != "-é-" {
		t.Fatalf("ReplaceAll = %q, want -é-", got)
	}
	if got := fmt.Sprintf("% x", value.Swap(map[string]string{"": "-"}).String()); got != "2d c3 2d a9 2d" {
		t.Fatalf("Swap bytes = %s, want 2d c3 2d a9 2d", got)
	}
}
