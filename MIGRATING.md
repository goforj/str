# Migrating from v1 to v2

Version 2 removes aliases, makes defaults explicit, and adopts Go standard-library vocabulary where a familiar name exists. Update the module path first:

```go
import "github.com/goforj/str/v2"
```

## Canonical replacements

| v1 | v2 |
| --- | --- |
| `Len()` or `RuneCount()` | `RuneCount()` |
| `Trim("")` or `TrimSpace()` | `Trim()` |
| `Trim(chars)` | `TrimChars(chars)` |
| `TrimLeft("")` / `TrimRight("")` | `TrimLeft()` / `TrimRight()` |
| `TrimLeft(chars)` / `TrimRight(chars)` | Go's `strings.TrimLeft` / `strings.TrimRight` for custom one-sided cutsets |
| `value.Join(elements, sep)` | Unchanged; `Join` remains fluent |
| `StartsWith(prefix)` | `HasPrefix(prefix)` |
| `StartsWithFold(prefix)` | `HasPrefixFold(prefix)` |
| `EndsWith(suffix)` | `HasSuffix(suffix)` |
| `EndsWithFold(suffix)` | `HasSuffixFold(suffix)` |
| `EqualsFold(other)` | `EqualFold(other)` |
| `ChopStart(prefix)` / `ChopEnd(suffix)` | `TrimPrefix(prefix)` / `TrimSuffix(suffix)` |
| `ReplaceStart(old, repl)` / `ReplaceEnd(old, repl)` | `ReplacePrefix(old, repl)` / `ReplaceSuffix(old, repl)` |
| `Snake("_")` | `Snake()` |
| `Snake("-")` | `Kebab()` |
| `Slug("-")` | `Slug()` |
| `Squish()` | `NormalizeSpace()` |
| `BetweenFirst(start, end)` | `Between(start, end)` |
| `Is(pattern)` | `Match(pattern)`, which also returns malformed-pattern errors |
| `UcSplit()` | `SplitWords()` |

`Contains`, `ContainsFold`, `HasPrefix`, `HasPrefixFold`, `HasSuffix`, and `HasSuffixFold` accept one search term. Loop at the call site when application logic requires any/all matching; the library no longer maintains parallel variadic and `All` APIs.

## Removed wrappers

Use the corresponding Go operation directly for functionality that does not benefit from another fluent spelling:

| Removed v1 API | Replacement |
| --- | --- |
| `Equals(other)` | `value.String() == other` |
| `NewLine()` | `Append("\n")` |
| `HasSurrounding(before, after)` | `HasPrefix(before) && HasSuffix(after)` |
| `ContainsAll` / `ContainsAllFold` | A call-site loop over `Contains` / `ContainsFold` |
| `ToTitle()` | `strings.ToTitle` when its specialized Unicode mapping is required |
| `IsMatch`, regexp `Match`/`MatchAll`, `ReplaceMatches` | Go's `regexp` package |
| Fold-specific before/after/count/index and first/last replacement variants | Application logic using `EqualFold`, `ContainsFold`, or `ReplaceFold` |
| `UcWords()` | `Title()` when lowercasing the remaining letters is appropriate, otherwise application-specific casing |
| `Transliterate()` | A dedicated transliteration package with an explicit language policy |

## Behavior changes

- Empty search terms are no matches. Replacement methods are no-ops when their search term is empty.
- `Between` uses the first closing marker after the first opening marker.
- `NormalizeSpace` trims surrounding whitespace and collapses internal Unicode whitespace.
- `Words` and `WrapWords` preserve punctuation instead of rebuilding text from stripped tokens.
- `Slug` preserves Unicode letters and digits, lowercases them, and always uses hyphens. It no longer claims partial ASCII transliteration.
- `PadLeft`, `PadRight`, and `PadBoth` never truncate or erase a value when the requested width is already satisfied.
- `Wrap` and `Unwrap` treat `before` and `after` literally. An empty `after` no longer reuses `before`.
- Case-insensitive operations consistently use Unicode simple folding.
- Identifier casing consistently recognizes acronym boundaries such as `HTTPRequestID` → `http_request_id`.
- `Plural` and `Singular` apply expanded English rules to the final word in an identifier.
