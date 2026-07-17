<p align="center">
  <img src="https://raw.githubusercontent.com/goforj/str/main/docs/images/logo.png" width="300" alt="str logo">
</p>

<p align="center">
  Fluent string helpers for Go.
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/str/v2"><img src="https://pkg.go.dev/badge/github.com/goforj/str/v2.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
    <a href="https://github.com/goforj/str/actions"><img src="https://github.com/goforj/str/actions/workflows/test.yml/badge.svg" alt="Go Test"></a>
    <a href="https://go.dev"><img src="https://img.shields.io/badge/go-1.24%2B-blue?logo=go" alt="Go 1.24 or newer"></a>
    <img src="https://img.shields.io/github/v/tag/goforj/str?label=version&sort=semver" alt="Latest tag">
    <a href="https://codecov.io/gh/goforj/str"><img src="https://codecov.io/github/goforj/str/graph/badge.svg?token=9KT46ZORP3" alt="Coverage"></a>
<!-- test-count:embed:start -->
    <img src="https://img.shields.io/badge/tests-281-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
</p>

`str` wraps a Go string so cleanup and transformation steps can be chained from left to right. Method names follow the standard library where possible, and operations that count, slice, or pad text work in runes rather than bytes.

## Installation

Requires Go 1.24 or newer.

```sh
go get github.com/goforj/str/v2
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

func main() {
	result := str.Of("  welcome_to_go  ").Trim().Headline().String()
	fmt.Println(result) // Welcome to Go
}
```

## API principles

`str` keeps the API deliberately small. These rules decide what belongs:

- **Chains come first.** Start with `str.Of`. Methods that change text return a new `str.String`, so the chain can continue. Checks, counts, parsers, and splits return ordinary Go values.
- **One job, one name.** There are no aliases or compatibility shims. If two names mean the same thing, keep the clearer one.
- **Use Go's words.** When the standard library already names an operation, use the same name and argument order.
- **Work in runes.** Character positions, counts, slices, padding, and case changes handle Unicode text instead of UTF-8 bytes unless the method says otherwise.
- **Do not hide edge cases.** Empty searches do not match, replacing an empty search does nothing, and parsing or pattern errors are returned to the caller.
- **Every method earns its place.** It must solve a common application problem or make a chain meaningfully clearer. Narrow, project-specific rules belong elsewhere.
- **Examples must keep working.** Every public operation has a generated example, and the test suite runs each one and checks its output.

## Why not just the standard library?

Often, you should. Go's `strings`, `unicode`, `strconv`, and `regexp` packages are the right choice when you only need one or two operations. This is already clear:

```go
username := strings.ToLower(strings.TrimSpace("  GoForj_Admin  "))
// goforj_admin
```

The same cleanup with `str` reads from left to right:

```go
username := str.Of("  GoForj_Admin  ").Trim().ToLower().String()
// goforj_admin
```

Either version is reasonable. The difference is easier to see when more rules belong together.

Using the standard library:

```go
func configKey(name string) string {
	key := strings.TrimSpace(name)
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, "-", "_")
	key = strings.Trim(key, "_")
	if !strings.HasPrefix(key, "APP_") {
		key = "APP_" + key
	}
	return key
}

// configKey("  --billing-worker--  ") == "APP_BILLING_WORKER"
```

Using `str`:

```go
func configKey(name string) string {
	return str.Of(name).
		Trim().
		ToUpper().
		ReplaceAll("-", "_").
		TrimChars("_").
		EnsurePrefix("APP_").
		String()
}

// configKey("  --billing-worker--  ") == "APP_BILLING_WORKER"
```

Some jobs do not have a single standard library call. For example, `Slug` handles case, punctuation, repeated separators, and Unicode letters. It can be one step in a longer chain that turns a report title into a CSV filename with a 64-rune base name:

```go
func exportFilename(reportTitle string) string {
	return str.Of(reportTitle).
		ReplaceAll("&", "and").
		Slug().
		Take(64).
		TrimChars("-").
		EnsurePrefix("report-").
		EnsureSuffix(".csv").
		String()
}

filename := exportFilename("Q3 Sales & Returns — North America")
// report-q3-sales-and-returns-north-america.csv
```

`str` uses the standard library underneath and has no dependencies. Use whichever version makes the rules easiest to see.

<!-- performance:embed:start -->

## Performance

These comparisons measure equivalent standard-library and `str` operations. Each cell reports the median of 10 samples as `ns/op · B/op · allocs/op`.

Recorded with `go1.26.1` on `linux/arm64` using `-cpu=1` (`GOMAXPROCS=1`).

| Workload | Standard library | `str` chain |
| --- | ---: | ---: |
| Trim | 3.2 ns/op · 0 B/op · 0 allocs/op | 3.2 ns/op · 0 B/op · 0 allocs/op |
| ToLower | 61.2 ns/op · 32 B/op · 1 allocs/op | 63.0 ns/op · 32 B/op · 1 allocs/op |
| NormalizeSpace (Fields + Join) | 191.6 ns/op · 208 B/op · 2 allocs/op | 185.2 ns/op · 80 B/op · 1 allocs/op |
| Trim → ToLower | 71.8 ns/op · 32 B/op · 1 allocs/op | 69.3 ns/op · 32 B/op · 1 allocs/op |
| ReplaceAll × 3 | 136.6 ns/op · 96 B/op · 3 allocs/op | 133.1 ns/op · 96 B/op · 3 allocs/op |

Timing is machine-specific; use it to understand the scale of these operations, not as a universal speed claim. Treat small timing differences within the raw sample spread as noise. Allocation counts are less sensitive to machine speed and show how much heap work each composition performs. In these workloads, wrapping and unwrapping added no heap allocations; allocations came from transformations that produced new text. `NormalizeSpace` is algorithmically different: the standard-library composition builds a field slice before joining it, while `str` uses one builder pass.

The [benchmark source](string_benchmark_test.go) and [committed raw output](docs/readme/benchmarks.txt) record exactly what ran, including the Go version and command. Refresh the measurements explicitly with `go -C docs run ./readme -record-benchmarks`; ordinary README generation only renders that frozen snapshot.
<!-- performance:embed:end -->

<!-- api:embed:start -->

## API index

The full API and these examples are also available on [pkg.go.dev](https://pkg.go.dev/github.com/goforj/str/v2).

| Group | API |
| --- | --- |
| Affixes | [EnsurePrefix](#ensureprefix) · [EnsureSuffix](#ensuresuffix) · [TrimPrefix](#trimprefix) · [TrimSuffix](#trimsuffix) · [Unwrap](#unwrap) · [Wrap](#wrap) |
| Case | [Camel](#camel) · [Headline](#headline) · [Kebab](#kebab) · [LcFirst](#lcfirst) · [Pascal](#pascal) · [Snake](#snake) · [Title](#title) · [ToLower](#tolower) · [ToUpper](#toupper) · [UcFirst](#ucfirst) |
| Checks | [IsASCII](#isascii) · [IsAlnum](#isalnum) · [IsAlpha](#isalpha) · [IsBlank](#isblank) · [IsEmpty](#isempty) · [IsNumeric](#isnumeric) |
| Cleanup | [Deduplicate](#deduplicate) · [NormalizeNewlines](#normalizenewlines) · [NormalizeSpace](#normalizespace) · [Trim](#trim) · [TrimChars](#trimchars) · [TrimLeft](#trimleft) · [TrimRight](#trimright) |
| Comparison | [EqualFold](#equalfold) |
| Compose | [Append](#append) · [Prepend](#prepend) |
| Constructor | [Of](#of) |
| Conversion | [Bool](#bool) · [Float64](#float64) · [Int](#int) |
| Encoding | [FromBase64](#frombase64) · [ToBase64](#tobase64) |
| Fluent | [GoString](#gostring) · [String](#string) |
| Length | [RuneCount](#runecount) |
| Masking | [Mask](#mask) |
| Match | [Match](#match) |
| Padding | [PadBoth](#padboth) · [PadLeft](#padleft) · [PadRight](#padright) |
| Pluralize | [Plural](#plural) · [Singular](#singular) |
| Replace | [Remove](#remove) · [ReplaceAll](#replaceall) · [ReplaceArray](#replacearray) · [ReplaceFirst](#replacefirst) · [ReplaceFold](#replacefold) · [ReplaceLast](#replacelast) · [ReplacePrefix](#replaceprefix) · [ReplaceSuffix](#replacesuffix) · [Swap](#swap) |
| Search | [Contains](#contains) · [ContainsFold](#containsfold) · [Count](#count) · [HasPrefix](#hasprefix) · [HasPrefixFold](#hasprefixfold) · [HasSuffix](#hassuffix) · [HasSuffixFold](#hassuffixfold) · [Index](#index) · [LastIndex](#lastindex) |
| Slug | [Slug](#slug) |
| Snippet | [Excerpt](#excerpt) |
| Split | [Lines](#lines) · [Split](#split) |
| Substrings | [After](#after) · [AfterLast](#afterlast) · [Before](#before) · [BeforeLast](#beforelast) · [Between](#between) · [CharAt](#charat) · [CommonPrefix](#commonprefix) · [CommonSuffix](#commonsuffix) · [Limit](#limit) · [Slice](#slice) · [SubstrReplace](#substrreplace) · [Take](#take) · [TakeLast](#takelast) |
| Transform | [Repeat](#repeat) · [Reverse](#reverse) |
| Words | [FirstWord](#firstword) · [Initials](#initials) · [Join](#join) · [LastWord](#lastword) · [SplitWords](#splitwords) · [WordCount](#wordcount) · [Words](#words) · [WrapWords](#wrapwords) |

## API examples

These examples come from GoDoc and run as part of the test suite.

### Affixes

#### <a id="ensureprefix"></a>EnsurePrefix

EnsurePrefix ensures the string starts with prefix, adding it if missing.
Similar: EnsureSuffix and TrimPrefix.

```go
v := str.Of("path/to").EnsurePrefix("/").String()
println(v)
// #string /path/to
```

#### <a id="ensuresuffix"></a>EnsureSuffix

EnsureSuffix ensures the string ends with suffix, adding it if missing.
Similar: EnsurePrefix and TrimSuffix.

```go
v := str.Of("path/to").EnsureSuffix("/").String()
println(v)
// #string path/to/
```

#### <a id="trimprefix"></a>TrimPrefix

TrimPrefix removes prefix when it appears at the start of the string.
Similar: TrimSuffix and EnsurePrefix.

```go
v := str.Of("https://goforj.dev").TrimPrefix("https://").String()
println(v)
// #string goforj.dev
```

#### <a id="trimsuffix"></a>TrimSuffix

TrimSuffix removes suffix when it appears at the end of the string.
Similar: TrimPrefix and EnsureSuffix.

```go
v := str.Of("file.txt").TrimSuffix(".txt").String()
println(v)
// #string file
```

#### <a id="unwrap"></a>Unwrap

Unwrap removes matching before and after strings if present.
Similar: Wrap.

```go
v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
println(v)
// #string GoForj
```

#### <a id="wrap"></a>Wrap

Wrap surrounds the string with before and after.
Similar: Unwrap.

```go
v := str.Of("GoForj").Wrap(`"`, `"`).String()
println(v)
// #string "GoForj"
```

### Case

#### <a id="camel"></a>Camel

Camel converts the string to camelCase.
Similar: Pascal.

```go
v := str.Of("foo_bar baz").Camel().String()
println(v)
// #string fooBarBaz
```

#### <a id="headline"></a>Headline

Headline converts the string into a human-friendly headline:
splits on case/underscores/dashes/whitespace, title-cases words, and lowercases small words (except the first).
Similar: Title.

```go
v := str.Of("emailNotification_sent").Headline().String()
println(v)
// #string Email Notification Sent
```

#### <a id="kebab"></a>Kebab

Kebab converts the string to kebab-case.
Similar: Snake.

```go
v := str.Of("fooBar baz").Kebab().String()
println(v)
// #string foo-bar-baz
```

#### <a id="lcfirst"></a>LcFirst

LcFirst returns the string with the first rune lower-cased.
Similar: UcFirst and ToLower.

```go
v := str.Of("Gopher").LcFirst().String()
fmt.Println(v)
// #string gopher
```

#### <a id="pascal"></a>Pascal

Pascal converts the string to PascalCase.
Similar: Camel.

```go
v := str.Of("foo_bar baz").Pascal().String()
fmt.Println(v)
// #string FooBarBaz
```

#### <a id="snake"></a>Snake

Snake converts the string to snake_case.
Similar: Kebab.

```go
v := str.Of("fooBar baz").Snake().String()
println(v)
// #string foo_bar_baz
```

#### <a id="title"></a>Title

Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.
Similar: Headline.

```go
v := str.Of("a nice title uses the correct case").Title().String()
println(v)
// #string A Nice Title Uses The Correct Case
```

#### <a id="tolower"></a>ToLower

ToLower returns a lowercase copy of the string using Unicode rules.
Similar: ToUpper and LcFirst.

```go
v := str.Of("GoLang").ToLower().String()
println(v)
// #string golang
```

#### <a id="toupper"></a>ToUpper

ToUpper returns an uppercase copy of the string using Unicode rules.
Similar: ToLower and UcFirst.

```go
v := str.Of("GoLang").ToUpper().String()
println(v)
// #string GOLANG
```

#### <a id="ucfirst"></a>UcFirst

UcFirst returns the string with the first rune upper-cased.
Similar: LcFirst and ToUpper.

```go
v := str.Of("gopher").UcFirst().String()
println(v)
// #string Gopher
```

### Checks

#### <a id="isascii"></a>IsASCII

IsASCII reports whether the string consists solely of 7-bit ASCII runes.

```go
v := str.Of("gopher").IsASCII()
println(v)
// #bool true
```

#### <a id="isalnum"></a>IsAlnum

IsAlnum reports whether the string contains at least one rune and every rune is a Unicode letter or number.

```go
v := str.Of("Gopher2025").IsAlnum()
println(v)
// #bool true
```

#### <a id="isalpha"></a>IsAlpha

IsAlpha reports whether the string contains at least one rune and every rune is a Unicode letter.

```go
v := str.Of("Gopher").IsAlpha()
println(v)
// #bool true
```

#### <a id="isblank"></a>IsBlank

IsBlank reports whether the string contains only Unicode whitespace.
Similar: IsEmpty.

```go
v := str.Of("  \t\n")
println(v.IsBlank())
// #bool true
```

#### <a id="isempty"></a>IsEmpty

IsEmpty reports whether the string has zero length.
Similar: IsBlank.

```go
v := str.Of("").IsEmpty()
println(v)
// #bool true
```

#### <a id="isnumeric"></a>IsNumeric

IsNumeric reports whether the string contains at least one rune and every rune is a Unicode number.

```go
v := str.Of("12345").IsNumeric()
println(v)
// #bool true
```

### Cleanup

#### <a id="deduplicate"></a>Deduplicate

Deduplicate collapses consecutive instances of char into a single instance.
If char is zero, space is used.
Similar: NormalizeSpace.

```go
v := str.Of("The   Go   Playground").Deduplicate(' ').String()
println(v)
// #string The Go Playground
```

#### <a id="normalizenewlines"></a>NormalizeNewlines

NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.
Similar: Lines.

```go
v := str.Of("a\r\nb\u2028c").NormalizeNewlines().String()
println(v)
// #string a\nb\nc
```

#### <a id="normalizespace"></a>NormalizeSpace

NormalizeSpace removes surrounding whitespace and collapses internal whitespace to single spaces.
Similar: Trim.

```go
v := str.Of("  go   forj  ").NormalizeSpace().String()
println(v)
// #string go forj
```

#### <a id="trim"></a>Trim

Trim removes leading and trailing Unicode whitespace.
Similar: TrimLeft, TrimRight, and TrimChars.

```go
v := str.Of("  GoForj  ").Trim().String()
println(v)
// #string GoForj
```

#### <a id="trimchars"></a>TrimChars

TrimChars removes leading and trailing runes contained in cutset.
Similar: Trim.

```go
v := str.Of("..GoForj!!").TrimChars(".!").String()
println(v)
// #string GoForj
```

#### <a id="trimleft"></a>TrimLeft

TrimLeft removes leading Unicode whitespace.
Similar: Trim and TrimRight.

```go
v := str.Of("  GoForj  ").TrimLeft().String()
println(v)
// #string GoForj\u0020\u0020
```

#### <a id="trimright"></a>TrimRight

TrimRight removes trailing Unicode whitespace.
Similar: Trim and TrimLeft.

```go
v := str.Of("  GoForj  ").TrimRight().String()
println(v)
// #string \u0020\u0020GoForj
```

### Comparison

#### <a id="equalfold"></a>EqualFold

EqualFold reports whether the string matches other using Unicode simple case folding.

```go
v := str.Of("gopher").EqualFold("GOPHER")
println(v)
// #bool true
```

### Compose

#### <a id="append"></a>Append

Append concatenates the provided parts to the end of the string.
Similar: Prepend.

```go
v := str.Of("Go").Append("Forj", "!").String()
println(v)
// #string GoForj!
```

#### <a id="prepend"></a>Prepend

Prepend concatenates the provided parts to the beginning of the string.
Similar: Append.

```go
v := str.Of("World").Prepend("Hello ", "Go ").String()
println(v)
// #string Hello Go World
```

### Constructor

#### <a id="of"></a>Of

Of wraps a raw string with fluent helpers.

```go
v := str.Of("gopher")
println(v.String())
// #string gopher
```

### Conversion

#### <a id="bool"></a>Bool

Bool parses the string as a bool using strconv.ParseBool semantics.
Similar: Int and Float64.

```go
v, err := str.Of("true").Bool()
println(v, err == nil)
// #bool true
// #bool true
```

#### <a id="float64"></a>Float64

Float64 parses the string as a float64 using strconv.ParseFloat semantics.
Similar: Bool and Int.

```go
v, err := str.Of("3.14").Float64()
fmt.Println(v, err == nil)
// #float64 3.14
// #bool true
```

#### <a id="int"></a>Int

Int parses the string as a base-10 int using strconv.Atoi semantics.
Similar: Bool and Float64.

```go
v, err := str.Of("42").Int()
println(v, err == nil)
// #int 42
// #bool true
```

### Encoding

#### <a id="frombase64"></a>FromBase64

FromBase64 decodes a standard Base64 string.
Similar: ToBase64.

```go
v, err := str.Of("Z29waGVy").FromBase64()
println(v.String(), err == nil)
// #string gopher
// #bool true
```

#### <a id="tobase64"></a>ToBase64

ToBase64 encodes the string using standard Base64.
Similar: FromBase64.

```go
v := str.Of("gopher").ToBase64().String()
println(v)
// #string Z29waGVy
```

### Fluent

#### <a id="gostring"></a>GoString

GoString allows %#v formatting to print the raw string.

```go
v := str.Of("go")
println(fmt.Sprintf("%#v", v))
// #string go
```

#### <a id="string"></a>String

String returns the underlying raw string value.

```go
v := str.Of("go").String()
println(v)
// #string go
```

### Length

#### <a id="runecount"></a>RuneCount

RuneCount returns the number of Unicode code points in the string.

```go
v := str.Of("gophers 🦫").RuneCount()
println(v)
// #int 9
```

### Masking

#### <a id="mask"></a>Mask

Mask replaces the middle of the string with the given rune, revealing revealLeft runes
at the start and revealRight runes at the end. Negative reveal values count from the end.
If the reveal counts cover the whole string, the original string is returned.

```go
v := str.Of("gopher@example.com").Mask('*', 3, 4).String()
println(v)
// #string gop***********.com
```

### Match

#### <a id="match"></a>Match

Match reports whether the entire string matches pattern using [path.Match] syntax.
A malformed pattern returns an error, and wildcards do not match a slash.

```go
matched, err := str.Of("billing:reports").Match("billing:*")
println(matched, err == nil)
// #bool true
// #bool true
```

### Padding

#### <a id="padboth"></a>PadBoth

PadBoth pads the string on both sides to reach length runes using pad (defaults to space).
Widths at or below the current rune width leave the string unchanged.
Similar: PadLeft and PadRight.

```go
v := str.Of("go").PadBoth(6, "-").String()
println(v)
// #string --go--
```

#### <a id="padleft"></a>PadLeft

PadLeft pads the string on the left to reach length runes using pad (defaults to space).
Widths at or below the current rune width leave the string unchanged.
Similar: PadRight and PadBoth.

```go
v := str.Of("go").PadLeft(5, " ").String()
println(v)
// #string \u0020\u0020\u0020go
```

#### <a id="padright"></a>PadRight

PadRight pads the string on the right to reach length runes using pad (defaults to space).
Widths at or below the current rune width leave the string unchanged.
Similar: PadLeft and PadBoth.

```go
v := str.Of("go").PadRight(5, ".").String()
println(v)
// #string go...
```

### Pluralize

#### <a id="plural"></a>Plural

Plural returns a best-effort English plural form of the final identifier word.
It handles common English forms and identifier boundaries without claiming to
resolve every irregular or ambiguous noun.
Similar: Singular.

```go
v := str.Of("city").Plural().String()
println(v)
// #string cities
```

#### <a id="singular"></a>Singular

Singular returns a best-effort English singular form of the final identifier word.
It handles common English forms and identifier boundaries without claiming to
resolve every irregular or ambiguous noun.
Similar: Plural.

```go
v := str.Of("people").Singular().String()
println(v)
// #string person
```

### Replace

#### <a id="remove"></a>Remove

Remove deletes all occurrences of provided substrings.

```go
v := str.Of("The Go Toolkit").Remove("Go ").String()
println(v)
// #string The Toolkit
```

#### <a id="replaceall"></a>ReplaceAll

ReplaceAll replaces all occurrences of old with new in the string.
If old is empty, the original string is returned unchanged.

```go
v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
println(v)
// #string Go Gopher Go
```

#### <a id="replacearray"></a>ReplaceArray

ReplaceArray replaces all occurrences of each old in olds with repl.
Similar: ReplaceAll and Swap.

```go
v := str.Of("The---Go---Toolkit")
println(v.ReplaceArray([]string{"---"}, "-").String())
// #string The-Go-Toolkit
```

#### <a id="replacefirst"></a>ReplaceFirst

ReplaceFirst replaces the first occurrence of old with repl.
Similar: ReplaceLast and ReplaceAll.

```go
v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
println(v)
// #string go gopher
```

#### <a id="replacefold"></a>ReplaceFold

ReplaceFold replaces all non-overlapping occurrences of old with repl using Unicode simple case folding.
An empty old string leaves the receiver unchanged.
Similar: ReplaceAll.

```go
v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
println(v)
// #string Go Gopher Go
```

#### <a id="replacelast"></a>ReplaceLast

ReplaceLast replaces the last occurrence of old with repl.
Similar: ReplaceFirst and ReplaceAll.

```go
v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
println(v)
// #string gopher go
```

#### <a id="replaceprefix"></a>ReplacePrefix

ReplacePrefix replaces old with repl when old is a prefix of the string.
Similar: ReplaceSuffix and TrimPrefix.

```go
v := str.Of("prefix-value").ReplacePrefix("prefix-", "new-").String()
println(v)
// #string new-value
```

#### <a id="replacesuffix"></a>ReplaceSuffix

ReplaceSuffix replaces old with repl when old is a suffix of the string.
Similar: ReplacePrefix and TrimSuffix.

```go
v := str.Of("file.old").ReplaceSuffix(".old", ".new").String()
println(v)
// #string file.new
```

#### <a id="swap"></a>Swap

Swap replaces multiple values using strings.Replacer built from a map.
Similar: ReplaceArray.

```go
pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
v := str.Of("Gophers are great!").Swap(pairs).String()
println(v)
// #string GoForj is fantastic!
```

### Search

#### <a id="contains"></a>Contains

Contains reports whether the string contains sub using a case-sensitive comparison.
An empty substring is not a match.
Similar: ContainsFold.

```go
v := str.Of("Go means gophers").Contains("gopher")
println(v)
// #bool true
```

#### <a id="containsfold"></a>ContainsFold

ContainsFold reports whether the string contains sub using Unicode simple case folding.
An empty substring is not a match.
Similar: Contains.

```go
v := str.Of("Go means gophers").ContainsFold("GOPHER")
println(v)
// #bool true
```

#### <a id="count"></a>Count

Count returns the number of non-overlapping occurrences of sub.

```go
v := str.Of("gogophergo").Count("go")
println(v)
// #int 3
```

#### <a id="hasprefix"></a>HasPrefix

HasPrefix reports whether the string starts with prefix using a case-sensitive comparison.
An empty prefix is not a match.
Similar: HasPrefixFold and HasSuffix.

```go
v := str.Of("gopher").HasPrefix("go")
println(v)
// #bool true
```

#### <a id="hasprefixfold"></a>HasPrefixFold

HasPrefixFold reports whether the string starts with prefix using Unicode simple case folding.
An empty prefix is not a match.
Similar: HasPrefix and HasSuffixFold.

```go
v := str.Of("gopher").HasPrefixFold("GO")
println(v)
// #bool true
```

#### <a id="hassuffix"></a>HasSuffix

HasSuffix reports whether the string ends with suffix using a case-sensitive comparison.
An empty suffix is not a match.
Similar: HasSuffixFold and HasPrefix.

```go
v := str.Of("gopher").HasSuffix("her")
println(v)
// #bool true
```

#### <a id="hassuffixfold"></a>HasSuffixFold

HasSuffixFold reports whether the string ends with suffix using Unicode simple case folding.
An empty suffix is not a match.
Similar: HasSuffix and HasPrefixFold.

```go
v := str.Of("gopher").HasSuffixFold("HER")
println(v)
// #bool true
```

#### <a id="index"></a>Index

Index returns the rune index of the first occurrence of sub, or -1 if not found.
Similar: LastIndex.

```go
v := str.Of("héllo").Index("llo")
println(v)
// #int 2
```

#### <a id="lastindex"></a>LastIndex

LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.
Similar: Index.

```go
v := str.Of("go gophers go").LastIndex("go")
println(v)
// #int 11
```

### Slug

#### <a id="slug"></a>Slug

Slug returns a lowercase Unicode slug separated by hyphens.
Unicode letters and digits are preserved, while every other run is collapsed
to one hyphen.
Similar: Kebab.

```go
v := str.Of("Go Forj Toolkit").Slug().String()
println(v)
// #string go-forj-toolkit
```

### Snippet

#### <a id="excerpt"></a>Excerpt

Excerpt returns a snippet around the first occurrence of needle with the given radius.
If needle is not found, an empty string is returned. If radius <= 0, a default of 100 is used.
Omission is used at the start/end when text is trimmed (default "...").

```go
v := str.Of("This is my name").Excerpt("my", 3, "...")
println(v.String())
// #string ...is my na...
```

### Split

#### <a id="lines"></a>Lines

Lines splits the string into lines after normalizing newline variants.
Similar: NormalizeNewlines.

```go
v := str.Of("a\r\nb\nc").Lines()
fmt.Println(v)
// #[]string [a b c]
```

#### <a id="split"></a>Split

Split splits the string by the given separator.

```go
v := str.Of("a,b,c").Split(",")
fmt.Println(v)
// #[]string [a b c]
```

### Substrings

#### <a id="after"></a>After

After returns the substring after the first occurrence of sep.
If sep is empty or not found, the original string is returned.
Similar: AfterLast and Before.

```go
v := str.Of("gopher::go").After("::").String()
println(v)
// #string go
```

#### <a id="afterlast"></a>AfterLast

AfterLast returns the substring after the last occurrence of sep.
If sep is empty or not found, the original string is returned.
Similar: After and BeforeLast.

```go
v := str.Of("pkg/path/file.txt").AfterLast("/").String()
println(v)
// #string file.txt
```

#### <a id="before"></a>Before

Before returns the substring before the first occurrence of sep.
If sep is empty or not found, the original string is returned.
Similar: BeforeLast and After.

```go
v := str.Of("gopher::go").Before("::").String()
println(v)
// #string gopher
```

#### <a id="beforelast"></a>BeforeLast

BeforeLast returns the substring before the last occurrence of sep.
If sep is empty or not found, the original string is returned.
Similar: Before and AfterLast.

```go
v := str.Of("pkg/path/file.txt").BeforeLast("/").String()
println(v)
// #string pkg/path
```

#### <a id="between"></a>Between

Between returns the substring between the first start marker and the first end marker after it.
It returns an empty string when either marker is empty or missing.

```go
v := str.Of("[first] and [second]").Between("[", "]").String()
println(v)
// #string first
```

#### <a id="charat"></a>CharAt

CharAt returns the rune at the given index and true if within bounds.
Similar: Slice and RuneCount.

```go
v, ok := str.Of("gopher").CharAt(2)
println(string(v), ok)
// #string p
// #bool true
```

#### <a id="commonprefix"></a>CommonPrefix

CommonPrefix returns the longest shared prefix between the string and all provided others.
Comparison is rune-safe. If no others are provided, the original string is returned.
Similar: CommonSuffix.

```go
v := str.Of("gopher").CommonPrefix("go", "gold").String()
println(v)
// #string go
```

#### <a id="commonsuffix"></a>CommonSuffix

CommonSuffix returns the longest shared suffix between the string and all provided others.
Comparison is rune-safe. If no others are provided, the original string is returned.
Similar: CommonPrefix.

```go
v := str.Of("main_test.go").CommonSuffix("user_test.go", "api_test.go").String()
println(v)
// #string _test.go
```

#### <a id="limit"></a>Limit

Limit truncates the string to length runes, appending suffix if truncation occurs.

```go
v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
println(v)
// #string Perfectly\u0020...
```

#### <a id="slice"></a>Slice

Slice returns the substring between rune offsets [start:end).
Indices are clamped; if start >= end the result is empty.

```go
v := str.Of("naïve café").Slice(3, 7).String()
println(v)
// #string ve c
```

#### <a id="substrreplace"></a>SubstrReplace

SubstrReplace replaces the rune slice in [start:end) with repl.

```go
v := str.Of("naïve café").SubstrReplace("i", 2, 3).String()
println(v)
// #string naive café
```

#### <a id="take"></a>Take

Take returns the first length runes of the string (clamped).
Similar: TakeLast and Limit.

```go
v := str.Of("gophers").Take(3).String()
println(v)
// #string gop
```

#### <a id="takelast"></a>TakeLast

TakeLast returns the last length runes of the string (clamped).
Similar: Take.

```go
v := str.Of("gophers").TakeLast(4).String()
println(v)
// #string hers
```

### Transform

#### <a id="repeat"></a>Repeat

Repeat repeats the string count times (non-negative).

```go
v := str.Of("go").Repeat(3).String()
println(v)
// #string gogogo
```

#### <a id="reverse"></a>Reverse

Reverse returns a rune-safe reversed string.

```go
v := str.Of("naïve").Reverse().String()
println(v)
// #string evïan
```

### Words

#### <a id="firstword"></a>FirstWord

FirstWord returns the first detected word or an empty string.
Similar: LastWord and SplitWords.

```go
v := str.Of("Hello world")
println(v.FirstWord().String())
// #string Hello
```

#### <a id="initials"></a>Initials

Initials returns the uppercase first rune of each detected word.
Words are split the same way as SplitWords, including camel case and acronym boundaries.
Similar: SplitWords.

```go
v := str.Of("portableNetwork graphics").Initials().String()
println(v)
// #string PNG
```

#### <a id="join"></a>Join

Join concatenates elements with sep and returns the result to the fluent chain.
The receiver provides fluent access and is not included in elements.
Similar: Split.

```go
v := str.Of("").Join([]string{"foo", "bar"}, "-").String()
println(v)
// #string foo-bar
```

#### <a id="lastword"></a>LastWord

LastWord returns the last detected word or an empty string.
Similar: FirstWord and SplitWords.

```go
v := str.Of("Hello world").LastWord().String()
println(v)
// #string world
```

#### <a id="splitwords"></a>SplitWords

SplitWords splits the string into Unicode words, including camel case and acronym boundaries.
Similar: FirstWord, LastWord, WordCount, and Words.

```go
v := str.Of("one, two, three").SplitWords()
fmt.Println(v)
// #[]string [one two three]
```

#### <a id="wordcount"></a>WordCount

WordCount returns the number of detected words.
Similar: SplitWords.

```go
v := str.Of("Hello, world!").WordCount()
println(v)
// #int 2
```

#### <a id="words"></a>Words

Words limits the string to count words, preserving the source through the
selected word boundary and appending suffix if truncated.
Similar: SplitWords and WrapWords.

```go
v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
println(v)
// #string Perfectly balanced, as >>>
```

#### <a id="wrapwords"></a>WrapWords

WrapWords wraps the string to the given rune width on whitespace boundaries,
using breakStr between lines without discarding punctuation.
Similar: Words.

```go
v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
println(v)
// #string The quick brown fox\njumped over the lazy\ndog.
```
<!-- api:embed:end -->

## Documentation

- [API documentation and examples](https://pkg.go.dev/github.com/goforj/str/v2)
- [v1 to v2 migration guide](./MIGRATING.md)
- [Report a bug or request a feature](https://github.com/goforj/str/issues)

## Development

`docs` and `examples` are separate Go modules, keeping their tooling and generated programs out of the library module download.

Run the tests and rebuild the generated examples and README with:

```sh
go test ./...
go -C docs test ./...
go -C examples test ./...
go -C docs run ./examplegen
go -C docs run ./readme
```

Licensed under the [MIT License](./LICENSE).
