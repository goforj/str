# Standard string API contract

The v3 contract mirrors the non-constructor functions in Go 1.27's `strings` package while preserving Go 1.24 as the minimum. This is a fluent adaptation: the source string becomes the receiver, remaining arguments keep their order, and a single string result becomes `str.String`. Booleans, indexes, slices, iterators, and multiple results keep their standard types.

`Join(elements, sep) String` is a package-level constructor because there is no source string. `strings.Builder`, `strings.Reader`, and `strings.Replacer`, including `NewReader` and `NewReplacer`, remain in the standard library. The wrapper is immutable and does not reproduce their mutable or I/O interfaces.

## Complete function inventory

| Area | Operations |
| --- | --- |
| Comparison and containment | `Compare`, `Contains`, `ContainsAny`, `ContainsFunc`, `ContainsRune`, `EqualFold`, `HasPrefix`, `HasSuffix` |
| Counting and indexing | `Count`, `Index`, `IndexAny`, `IndexByte`, `IndexFunc`, `IndexRune`, `LastIndex`, `LastIndexAny`, `LastIndexByte`, `LastIndexFunc` |
| Cutting | `Cut`, `CutLast`, `CutPrefix`, `CutSuffix` |
| Fields and lines | `Fields`, `FieldsFunc`, `FieldsSeq`, `FieldsFuncSeq`, `Lines` |
| Splitting and joining | `Split`, `SplitN`, `SplitAfter`, `SplitAfterN`, `SplitSeq`, `SplitAfterSeq`, `Join` |
| Transformation | `Clone`, `Map`, `Repeat`, `Replace`, `ReplaceAll`, `ToValidUTF8` |
| Casing | `Title`, `ToLower`, `ToLowerSpecial`, `ToTitle`, `ToTitleSpecial`, `ToUpper`, `ToUpperSpecial` |
| Trimming | `Trim`, `TrimFunc`, `TrimLeft`, `TrimLeftFunc`, `TrimPrefix`, `TrimRight`, `TrimRightFunc`, `TrimSpace`, `TrimSuffix` |

There are 56 operations. `Title` mirrors the deprecated standard function and carries its deprecation notice. `CutLast` is backported using byte search because delegating to the Go 1.27 function would raise the minimum Go version.

## Boundaries that must remain stable

- Standard indexes are byte offsets, including `IndexRune` and predicate-based searches.
- Empty searches retain Go's operation-specific behavior. An empty cutset is not shorthand for whitespace.
- `Replace` supports zero and negative counts exactly as Go does. Negative `Repeat` counts and overflow panic.
- Split results preserve nil versus non-nil empty slices where Go distinguishes them.
- Sequence methods retain the standard lazy, single-use iterator contract, including newline bytes in `Lines`.
- Multi-result cut operations return built-in strings and `found`, not fluent wrappers.
- `Clone` must allocate independent backing storage for nonempty input.
- Malformed UTF-8 is accepted and handled by the corresponding standard function; the wrapper adds no validation or normalization layer.
- Predicate and mapper arguments preserve the standard callback and panic behavior.

Application helpers have their own documented contracts. `Slice` and width-based helpers operate on runes; `SplitWords` recognizes punctuation and identifier boundaries rather than acting as an alias for `Fields`. Folded predicates match empty strings, and replacement helpers use the corresponding standard insertion boundaries. No package-wide rule overrides the contracts of individual operations.

## Regression checks

`stdlib_parity_test.go` records each standard signature and verifies the receiver adaptation, argument ordering, and result types. It also inventories exported functions from the running toolchain's selected `strings` source files, so a new standard function prompts an explicit maintenance decision instead of silently escaping coverage.

The differential suite invokes every operation with Unicode, combining marks, malformed UTF-8, NULs, empty strings, absent delimiters, replacement limits, special casing, and nil callbacks. It compares values, slice shape, and panics directly. Separate checks cover iterator consumption and laziness, clone allocation, and repetition overflow. On Go 1.27 and newer, the complete differential suite uses the actual `strings.CutLast` as its reference. The fuzz target also compares all 56 operations against the standard baseline.

Run the normal test matrix on Go 1.24 and the current supported Go release before releasing a major. Keep the nested docs and examples modules in the matrix, regenerate the examples and README, and verify generation is idempotent. A new major requires reviewing the complete API and migration guide, not merely changing the module suffix.
