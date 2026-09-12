# Migrating from v2 to v3

Version 3 aligns standard-named operations with Go's `strings` package. It changes both signatures and runtime behavior, so updating the import path alone is not sufficient. The minimum remains Go 1.24, and the library remains dependency-free.

```go
import "github.com/goforj/str/v3"
```

Update your module requirement to the v3 release when available. The repository's `docs` and `examples` modules are development tooling, not additional runtime modules to import or release. Their local replacement intentionally exercises the unpublished library during development.

## Source changes

| v2 | v3 |
| --- | --- |
| `value.Trim()` | `value.TrimSpace()` |
| `value.TrimChars(chars)` | `value.Trim(chars)` |
| `value.TrimLeft()` | `value.TrimLeftFunc(unicode.IsSpace)` |
| `value.TrimRight()` | `value.TrimRightFunc(unicode.IsSpace)` |
| Custom one-sided cutsets through `strings` | `value.TrimLeft(chars)` or `value.TrimRight(chars)` |
| `value.Join(elements, sep)` | `str.Join(elements, sep)` |
| `value.Lines()` returning normalized `[]string` | `value.NormalizeNewlines().Split("\n")` to retain that behavior |
| `value.Lines()` when iterating actual lines | `for line := range value.Lines()` |

`TrimChars` and the receiver method `Join` are removed. There are no compatibility aliases or optional cutset arguments. `Trim("")`, `TrimLeft("")`, and `TrimRight("")` mean an empty cutset and leave the input unchanged; they do not mean whitespace trimming.

```go
name := str.Of("  GoForj  ").TrimSpace().ToLower().String()
fmt.Println(name)
// goforj

path := str.Of("///var/log/").TrimLeft("/").String()
fmt.Println(path)
// var/log/

key := str.Join([]string{"billing", "reports"}, ":").ToUpper().String()
fmt.Println(key)
// BILLING:REPORTS
```

## Byte indexes and rune helpers

`Index` and `LastIndex` now return byte offsets, as do `IndexAny`, `IndexByte`, `IndexFunc`, `IndexRune`, and the available last-index variants. `IndexRune` searches for a rune but still returns a byte offset.

```go
value := str.Of("héllo")
index := value.Index("llo")
fmt.Println(index)
// 3
fmt.Println(value.String()[index:])
// llo
```

`Slice`, `Take`, `TakeLast`, `CharAt`, `SubstrReplace`, `Mask`, padding, and rune-count/width helpers keep their documented rune-based contracts. An index returned by a standard search must not be passed directly into those operations. For a valid UTF-8 substring match, convert a byte offset explicitly when a rune position is required:

```go
raw := "héllo"
index := str.Of(raw).Index("llo")
if index >= 0 {
    runeIndex := utf8.RuneCountInString(raw[:index])
    fmt.Println(str.Of(raw).Slice(runeIndex, utf8.RuneCountInString(raw)).String())
    // llo
}
```

The conversion example assumes a rune-aligned match. Standard byte searches can match malformed UTF-8 or a byte within an encoded rune, so use ordinary string slicing for byte-oriented data.

## Empty searches and replacement boundaries

| Operation on `"abc"` | v2 | v3 |
| --- | --- | --- |
| `Contains("")`, `HasPrefix("")`, `HasSuffix("")` | false | true |
| `Count("")` | 0 | 4 |
| `Index("")` | -1 | 0 |
| `LastIndex("")` | -1 | 3 |
| `ReplaceAll("", "-")` | `"abc"` | `"-a-b-c-"` |
| `ReplaceFirst("", "-")` | `"abc"` | `"-abc"` |
| `ReplaceLast("", "-")` | `"abc"` | `"abc-"` |
| `ReplacePrefix("", "-")` | `"abc"` | `"-abc"` |
| `ReplaceSuffix("", "-")` | `"abc"` | `"abc-"` |

`ContainsFold`, `HasPrefixFold`, and `HasSuffixFold` also match empty searches. `ReplaceFold` and each entry in `ReplaceArray` use the same empty-search insertion rule as `ReplaceAll`. `Swap` continues to follow `strings.Replacer`: replacements happen in one pass, longer keys win at the same position, and empty keys retain the standard replacer's behavior. Replacement values are not rescanned by `Swap`; `ReplaceArray` remains sequential.

For non-ASCII input, empty replacement insertion occurs at UTF-8 sequence boundaries, not between the bytes of a valid rune. Applications that require the v2 no-op behavior should reject or skip empty input explicitly before calling a search or replacement operation.

`Before`, `After`, `BeforeLast`, and `AfterLast` retain their original-string result for an empty separator. `Between` and `Excerpt` retain their documented no-match results for empty markers. These application helpers do not replace `Cut`, `CutPrefix`, or `CutSuffix`, whose result includes a `found` flag and follows the standard empty-boundary rules.

## Line iteration

`Lines()` returns `iter.Seq[string]`. It preserves newline bytes and does not normalize CRLF, lone CR, or Unicode separators. Empty input yields no lines. A trailing newline belongs to the preceding line and does not create an extra empty line.

```go
for line := range str.Of("a\r\nb\n").Lines() {
    fmt.Printf("%q\n", line)
}
// "a\r\n"
// "b\n"
```

Use `slices.Collect(value.Lines())` when a slice of these exact lines is needed. A standard line iterator yields values with `for line := range ...`; keeping `for line := range ...` from a v2 slice loop changes `line` from an integer index to a string. Review range loops as well as compile errors.

All sequence methods follow the standard iterator contract. Consume each iterator once; obtain a fresh iterator for a new pass. `Split`, `SplitN`, `SplitAfter`, `SplitAfterN`, `Fields`, and `FieldsFunc` return slices.

## Casing and invalid counts

`Title` now mirrors `strings.Title`: it preserves letters after the word-initial letter. `str.Of("hELLO wORLD").Title().String()` produces `"HELLO WORLD"`, not `"Hello World"`. Combining marks no longer trigger the previous custom word-boundary algorithm. The method carries the same deprecation notice as `strings.Title` because its word boundaries do not handle Unicode punctuation properly. For linguistic title casing, use `golang.org/x/text/cases`; `ToLower().Title()` can approximate the old normalization for simple words but is not identical for every boundary. `Headline` remains the separate application-oriented headline helper. `ToTitle` maps every rune to Unicode titlecase and is not word title casing.

`Repeat(-1)` now panics, including on an empty source. Overflow also panics. Validate negative counts explicitly if they can come from user input. `Repeat(0)` still returns an empty string.

## Additional standard operations

The [standard API contract](docs/standard-api.md) lists the complete surface and test guarantees. New operations include `Replace(old, new, n)`, the `Cut` family, `Fields`, predicate-based trimming/searching, split limits and iterators, special-case conversion, `Map`, `Clone`, and `ToValidUTF8`. Multiple-result operations such as `Cut` return built-in strings and a boolean, so wrap a returned string with `str.Of` to start another chain.

`CutLast` provides Go 1.27 semantics through a local implementation that also compiles on Go 1.24. No new runtime dependency, configuration migration, or framework release is required to adopt v3. Existing v2 consumers can continue using the v2 module path.

For applications still using v1, consult the [historical v1 to v2 guide](docs/migrating-v2.md), then apply this guide's corrections before adopting v3.
