<p align="center">
  <img src="./docs/images/logo.png?v=2" width="400" alt="str logo">
</p>

<p align="center">
    A fluent, Laravel-inspired string toolkit for Go, focused on rune-safe helpers,
    expressive transformations, and predictable behavior beyond the standard library.
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/str"><img src="https://pkg.go.dev/badge/github.com/goforj/str.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
    <a href="https://github.com/goforj/str/actions"><img src="https://github.com/goforj/str/actions/workflows/test.yml/badge.svg" alt="Go Test"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/go-1.18+-blue?logo=go" alt="Go version"></a>
    <img src="https://img.shields.io/github/v/tag/goforj/str?label=version&sort=semver" alt="Latest tag">
    <a href="https://codecov.io/gh/goforj/str" ><img src="https://codecov.io/github/goforj/str/graph/badge.svg?token=9KT46ZORP3"/></a>
<!-- test-count:embed:start -->
    <img src="https://img.shields.io/badge/tests-196-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
    <a href="https://goreportcard.com/report/github.com/goforj/str"><img src="https://goreportcard.com/badge/github.com/goforj/str" alt="Go Report Card"></a>
</p>

## Installation

```sh
go get github.com/goforj/str
```

## Runnable examples

Every function has a corresponding runnable example under [`./examples`](./examples).

These examples are **generated directly from the documentation blocks** of each function, ensuring the docs and code never drift. These are the same examples you see here in the README and GoDoc.

An automated test executes **every example** to verify it builds and runs successfully.

This guarantees all examples are valid, up-to-date, and remain functional as the API evolves.

<!-- api:embed:start -->

## API Index

| Group | Functions |
|------:|-----------|
| **Affixes** | [ChopEnd](#chopend) [ChopStart](#chopstart) [EnsurePrefix](#ensureprefix) [EnsureSuffix](#ensuresuffix) [Unwrap](#unwrap) [Wrap](#wrap) |
| **Case** | [Camel](#camel) [Headline](#headline) [Kebab](#kebab) [LcFirst](#lcfirst) [Pascal](#pascal) [Snake](#snake) [Title](#title) [ToLower](#tolower) [ToTitle](#totitle) [ToUpper](#toupper) [UcFirst](#ucfirst) [UcWords](#ucwords) |
| **Checks** | [IsASCII](#isascii) [IsBlank](#isblank) [IsEmpty](#isempty) |
| **Cleanup** | [Deduplicate](#deduplicate) [NormalizeNewlines](#normalizenewlines) [NormalizeSpace](#normalizespace) [Squish](#squish) [Trim](#trim) [TrimLeft](#trimleft) [TrimRight](#trimright) [TrimSpace](#trimspace) |
| **Comparison** | [Equals](#equals) [EqualsFold](#equalsfold) |
| **Compose** | [Append](#append) [NewLine](#newline) [Prepend](#prepend) |
| **Constructor** | [Of](#of) |
| **Encoding** | [FromBase64](#frombase64) [ToBase64](#tobase64) |
| **Fluent** | [GoString](#gostring) [String](#string) |
| **Length** | [Len](#len) [RuneCount](#runecount) |
| **Masking** | [Mask](#mask) |
| **Match** | [Is](#is) [IsMatch](#ismatch) [Match](#match) [MatchAll](#matchall) |
| **Padding** | [PadBoth](#padboth) [PadLeft](#padleft) [PadRight](#padright) |
| **Pluralize** | [Plural](#plural) [Singular](#singular) |
| **Replace** | [Remove](#remove) [ReplaceAll](#replaceall) [ReplaceArray](#replacearray) [ReplaceEnd](#replaceend) [ReplaceFirst](#replacefirst) [ReplaceFirstFold](#replacefirstfold) [ReplaceFold](#replacefold) [ReplaceLast](#replacelast) [ReplaceLastFold](#replacelastfold) [ReplaceMatches](#replacematches) [ReplaceStart](#replacestart) [Swap](#swap) |
| **Search** | [Contains](#contains) [ContainsAll](#containsall) [ContainsAllFold](#containsallfold) [ContainsFold](#containsfold) [Count](#count) [EndsWith](#endswith) [EndsWithFold](#endswithfold) [Index](#index) [LastIndex](#lastindex) [StartsWith](#startswith) [StartsWithFold](#startswithfold) |
| **Slug** | [Slug](#slug) |
| **Snippet** | [Excerpt](#excerpt) |
| **Split** | [Lines](#lines) [Split](#split) [UcSplit](#ucsplit) |
| **Substrings** | [After](#after) [AfterLast](#afterlast) [Before](#before) [BeforeLast](#beforelast) [Between](#between) [BetweenFirst](#betweenfirst) [CharAt](#charat) [Limit](#limit) [Slice](#slice) [SubstrReplace](#substrreplace) [Take](#take) [TakeLast](#takelast) |
| **Transform** | [Repeat](#repeat) [Reverse](#reverse) [Transliterate](#transliterate) |
| **Words** | [FirstWord](#firstword) [Join](#join) [LastWord](#lastword) [SplitWords](#splitwords) [WordCount](#wordcount) [Words](#words) [WrapWords](#wrapwords) |


## Affixes

### <a id="chopend"></a>ChopEnd

ChopEnd removes the first matching suffix if present.

```go
v := str.Of("file.txt").ChopEnd(".txt", ".md").String()
println(v)
// #string file
```

### <a id="chopstart"></a>ChopStart

ChopStart removes the first matching prefix if present.

```go
v := str.Of("https://goforj.dev").ChopStart("https://", "http://").String()
println(v)
// #string goforj.dev
```

### <a id="ensureprefix"></a>EnsurePrefix

EnsurePrefix ensures the string starts with prefix, adding it if missing.

```go
v := str.Of("path/to").EnsurePrefix("/").String()
println(v)
// #string /path/to
```

### <a id="ensuresuffix"></a>EnsureSuffix

EnsureSuffix ensures the string ends with suffix, adding it if missing.

```go
v := str.Of("path/to").EnsureSuffix("/").String()
println(v)
// #string path/to/
```

### <a id="unwrap"></a>Unwrap

Unwrap removes matching before and after strings if present.

```go
v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
println(v)
// #string GoForj
```

### <a id="wrap"></a>Wrap

Wrap surrounds the string with before and after (after defaults to before).

```go
v := str.Of("GoForj").Wrap(`"`, "").String()
println(v)
// #string "GoForj"
```

## Case

### <a id="camel"></a>Camel

Camel converts the string to camelCase.

```go
v := str.Of("foo_bar baz").Camel().String()
println(v)
// #string fooBarBaz
```

### <a id="headline"></a>Headline

Headline converts the string into a human-friendly headline:
splits on case/underscores/dashes/whitespace, title-cases words, and lowercases small words (except the first).

```go
v := str.Of("emailNotification_sent").Headline().String()
println(v)
// #string Email Notification Sent
```

### <a id="kebab"></a>Kebab

Kebab converts the string to kebab-case.

```go
v := str.Of("fooBar baz").Kebab().String()
println(v)
// #string foo-bar-baz
```

### <a id="lcfirst"></a>LcFirst

LcFirst returns the string with the first rune lower-cased.

```go
v := str.Of("Gopher").LcFirst().String()
fmt.Println(v)
// #string gopher
```

### <a id="pascal"></a>Pascal

Pascal converts the string to PascalCase.

```go
v := str.Of("foo_bar baz").Pascal().String()
fmt.Println(v)
// #string FooBarBaz
```

### <a id="snake"></a>Snake

Snake converts the string to snake_case using the provided separator (default "_").

```go
v := str.Of("fooBar baz").Snake("_").String()
println(v)
// #string foo_bar_baz
```

### <a id="title"></a>Title

Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.

```go
v := str.Of("a nice title uses the correct case").Title().String()
println(v)
// #string A Nice Title Uses The Correct Case
```

### <a id="tolower"></a>ToLower

ToLower returns a lowercase copy of the string using Unicode rules.

```go
v := str.Of("GoLang").ToLower().String()
println(v)
// #string golang
```

### <a id="totitle"></a>ToTitle

ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.

```go
v := str.Of("ß").ToTitle().String()
println(v)
// #string SS
```

### <a id="toupper"></a>ToUpper

ToUpper returns an uppercase copy of the string using Unicode rules.

```go
v := str.Of("GoLang").ToUpper().String()
println(v)
// #string GOLANG
```

### <a id="ucfirst"></a>UcFirst

UcFirst returns the string with the first rune upper-cased.

```go
v := str.Of("gopher").UcFirst().String()
println(v)
// #string Gopher
```

### <a id="ucwords"></a>UcWords

UcWords uppercases the first rune of each word, leaving the rest unchanged.
Words are sequences of letters/digits.

```go
v := str.Of("hello WORLD").UcWords().String()
println(v)
// #string Hello WORLD
```

## Checks

### <a id="isascii"></a>IsASCII

IsASCII reports whether the string consists solely of 7-bit ASCII runes.

```go
v := str.Of("gopher").IsASCII()
println(v)
// #bool true
```

### <a id="isblank"></a>IsBlank

IsBlank reports whether the string contains only Unicode whitespace.

```go
v := str.Of("  \\t\\n")
println(v.IsBlank())
// #bool true
```

### <a id="isempty"></a>IsEmpty

IsEmpty reports whether the string has zero length.

```go
v := str.Of("").IsEmpty()
println(v)
// #bool true
```

## Cleanup

### <a id="deduplicate"></a>Deduplicate

Deduplicate collapses consecutive instances of char into a single instance.
If char is zero, space is used.

```go
v := str.Of("The   Go   Playground").Deduplicate(' ').String()
println(v)
// #string The Go Playground
```

### <a id="normalizenewlines"></a>NormalizeNewlines

NormalizeNewlines replaces CRLF, CR, and Unicode separators with \n.

```go
v := str.Of("a\\r\\nb\\u2028c").NormalizeNewlines().String()
println(v)
// #string a\nb\nc
```

### <a id="normalizespace"></a>NormalizeSpace

NormalizeSpace collapses whitespace runs to single spaces without trimming ends.

```go
v := str.Of("  go   forj  ").NormalizeSpace().String()
println(v)
// #string  go forj
```

### <a id="squish"></a>Squish

Squish trims leading/trailing whitespace and collapses internal whitespace to single spaces.

```go
v := str.Of("   go   forj  ").Squish().String()
println(v)
// #string go forj
```

### <a id="trim"></a>Trim

Trim trims leading and trailing characters. If cutset is the zero value (empty string), trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").Trim("").String()
println(v)
// #string GoForj
```

### <a id="trimleft"></a>TrimLeft

TrimLeft trims leading characters. If cutset is the zero value (empty string), trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").TrimLeft("").String()
println(v)
// #string GoForj
```

### <a id="trimright"></a>TrimRight

TrimRight trims trailing characters. If cutset is the zero value (empty string), trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").TrimRight("").String()
println(v)
// #string   GoForj
```

### <a id="trimspace"></a>TrimSpace

TrimSpace trims leading and trailing Unicode whitespace.

```go
v := str.Of("  GoForj  ").TrimSpace().String()
println(v)
// #string GoForj
```

## Comparison

### <a id="equals"></a>Equals

Equals reports whether the string exactly matches other (case-sensitive).

```go
v := str.Of("gopher").Equals("gopher")
println(v)
// #bool true
```

### <a id="equalsfold"></a>EqualsFold

EqualsFold reports whether the string matches other using Unicode case folding.

```go
v := str.Of("gopher").EqualsFold("GOPHER")
println(v)
// #bool true
```

## Compose

### <a id="append"></a>Append

Append concatenates the provided parts to the end of the string.

```go
v := str.Of("Go").Append("Forj", "!").String()
println(v)
// #string GoForj!
```

### <a id="newline"></a>NewLine

NewLine appends a newline character to the string.

```go
v := str.Of("hello").NewLine().Append("world").String()
println(v)
// #string hello\nworld
```

### <a id="prepend"></a>Prepend

Prepend concatenates the provided parts to the beginning of the string.

```go
v := str.Of("World").Prepend("Hello ", "Go ").String()
println(v)
// #string Hello Go World
```

## Constructor

### <a id="of"></a>Of

Of wraps a raw string with fluent helpers.

```go
v := str.Of("gopher")
println(v.String())
// #string gopher
```

## Encoding

### <a id="frombase64"></a>FromBase64

FromBase64 decodes a standard Base64 string.

```go
v, err := str.Of("Z29waGVy").FromBase64()
println(v.String(), err)
// #string gopher
// #error <nil>
```

### <a id="tobase64"></a>ToBase64

ToBase64 encodes the string using standard Base64.

```go
v := str.Of("gopher").ToBase64().String()
println(v)
// #string Z29waGVy
```

## Fluent

### <a id="gostring"></a>GoString

GoString allows %#v formatting to print the raw string.

```go
v := str.Of("go")
println(fmt.Sprintf("%#v", v))
// #string go
```

### <a id="string"></a>String

String returns the underlying raw string value.

```go
v := str.Of("go").String()
println(v)
// #string go
```

## Length

### <a id="len"></a>Len

Len returns the number of runes in the string.

```go
v := str.Of("gophers 🦫").Len()
println(v)
// #int 9
```

### <a id="runecount"></a>RuneCount

RuneCount is an alias for Len to make intent explicit in mixed codebases.

```go
v := str.Of("naïve").RuneCount()
println(v)
// #int 5
```

## Masking

### <a id="mask"></a>Mask

Mask replaces the middle of the string with the given rune, revealing revealLeft runes
at the start and revealRight runes at the end. Negative reveal values count from the end.
If the reveal counts cover the whole string, the original string is returned.

```go
v := str.Of("gopher@example.com").Mask('*', 3, 4).String()
println(v)
// #string gop***********.com
```

## Match

### <a id="is"></a>Is

Is reports whether the string matches any of the provided wildcard patterns.
Patterns use '*' as a wildcard. Case-sensitive.

```go
v := str.Of("foo/bar").Is("foo/*")
println(v)
// #bool true
```

### <a id="ismatch"></a>IsMatch

IsMatch reports whether the string matches the provided regular expression.

```go
v := str.Of("abc123").IsMatch(regexp.MustCompile(`\d+`))
println(v)
// #bool true
```

### <a id="match"></a>Match

Match returns the first match and submatches for the pattern.

```go
re := regexp.MustCompile(`g(o+)`)
v := str.Of("gooo").Match(re)
println(v)
// #[]string [gooo ooo]
```

### <a id="matchall"></a>MatchAll

MatchAll returns all matches and submatches for the pattern.

```go
re := regexp.MustCompile(`go+`)
v := str.Of("go gopher gooo").MatchAll(re)
println(v)
// #[][]string [[go] [gooo]]
```

## Padding

### <a id="padboth"></a>PadBoth

PadBoth pads the string on both sides to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadBoth(6, "-").String()
println(v)
// #string --go--
```

### <a id="padleft"></a>PadLeft

PadLeft pads the string on the left to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadLeft(5, " ").String()
println(v)
// #string \u00a0\u00a0\u00a0go
```

### <a id="padright"></a>PadRight

PadRight pads the string on the right to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadRight(5, ".").String()
println(v)
// #string go...
```

## Pluralize

### <a id="plural"></a>Plural

Plural returns a best-effort English plural form of the last word.

```go
v := str.Of("city").Plural().String()
println(v)
// #string cities
```

### <a id="singular"></a>Singular

Singular returns a best-effort English singular form of the last word.

```go
v := str.Of("people").Singular().String()
println(v)
// #string person
```

## Replace

### <a id="remove"></a>Remove

Remove deletes all occurrences of provided substrings.

```go
v := str.Of("The Go Toolkit").Remove("Go ").String()
println(v)
// #string The Toolkit
```

### <a id="replaceall"></a>ReplaceAll

ReplaceAll replaces all occurrences of old with new in the string.
If old is empty, the original string is returned unchanged.

```go
v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
println(v)
// #string Go Gopher Go
```

### <a id="replacearray"></a>ReplaceArray

ReplaceArray replaces all occurrences of each old in olds with repl.

```go
v := str.Of("The---Go---Toolkit")
println(v.ReplaceArray([]string{"---"}, "-").String())
// #string The-Go-Toolkit
```

### <a id="replaceend"></a>ReplaceEnd

ReplaceEnd replaces old with repl at the end of the string, if present.

```go
v := str.Of("file.old").ReplaceEnd(".old", ".new").String()
println(v)
// #string file.new
```

### <a id="replacefirst"></a>ReplaceFirst

ReplaceFirst replaces the first occurrence of old with repl.

```go
v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
println(v)
// #string go gopher
```

### <a id="replacefirstfold"></a>ReplaceFirstFold

ReplaceFirstFold replaces the first occurrence of old with repl using Unicode case folding.

```go
v := str.Of("go gopher GO").ReplaceFirstFold("GO", "Go").String()
println(v)
// #string Go gopher GO
```

### <a id="replacefold"></a>ReplaceFold

ReplaceFold replaces all occurrences of old with repl using Unicode case folding.

```go
v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
println(v)
// #string Go Gopher Go
```

### <a id="replacelast"></a>ReplaceLast

ReplaceLast replaces the last occurrence of old with repl.

```go
v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
println(v)
// #string gopher go
```

### <a id="replacelastfold"></a>ReplaceLastFold

ReplaceLastFold replaces the last occurrence of old with repl using Unicode case folding.

```go
v := str.Of("go gopher GO").ReplaceLastFold("GO", "Go").String()
println(v)
// #string go gopher Go
```

### <a id="replacematches"></a>ReplaceMatches

ReplaceMatches applies repl to each regex match and returns the result.

```go
re := regexp.MustCompile(`\d+`)
v := str.Of("Hello 123 World").ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String()
println(v)
// #string Hello [123] World
```

### <a id="replacestart"></a>ReplaceStart

ReplaceStart replaces old with repl at the start of the string, if present.

```go
v := str.Of("prefix-value").ReplaceStart("prefix-", "new-").String()
println(v)
// #string new-value
```

### <a id="swap"></a>Swap

Swap replaces multiple values using strings.Replacer built from a map.

```go
pairs := map[string]string{"Gophers": "GoForj", "are": "is", "great": "fantastic"}
v := str.Of("Gophers are great!").Swap(pairs).String()
println(v)
// #string GoForj is fantastic!
```

## Search

### <a id="contains"></a>Contains

Contains reports whether the string contains any of the provided substrings (case-sensitive).
Empty substrings return true to match strings.Contains semantics.

```go
v := str.Of("Go means gophers").Contains("rust", "gopher")
println(v)
// #bool true
```

### <a id="containsall"></a>ContainsAll

ContainsAll reports whether the string contains all provided substrings (case-sensitive).
Empty substrings are ignored.

```go
v := str.Of("Go means gophers").ContainsAll("Go", "gopher")
println(v)
// #bool true
```

### <a id="containsallfold"></a>ContainsAllFold

ContainsAllFold reports whether the string contains all provided substrings, using Unicode
case folding for comparisons.
Empty substrings are ignored.

```go
v := str.Of("Go means gophers").ContainsAllFold("go", "GOPHER")
println(v)
// #bool true
```

### <a id="containsfold"></a>ContainsFold

ContainsFold reports whether the string contains any of the provided substrings, using Unicode
case folding for comparisons.
Empty substrings return true.

```go
v := str.Of("Go means gophers").ContainsFold("GOPHER", "rust")
println(v)
// #bool true
```

### <a id="count"></a>Count

Count returns the number of non-overlapping occurrences of sub.

```go
v := str.Of("gogophergo").Count("go")
println(v)
// #int 3
```

### <a id="endswith"></a>EndsWith

EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).

```go
v := str.Of("gopher").EndsWith("her", "cat")
println(v)
// #bool true
```

### <a id="endswithfold"></a>EndsWithFold

EndsWithFold reports whether the string ends with any of the provided suffixes using Unicode case folding.

```go
v := str.Of("gopher").EndsWithFold("HER")
println(v)
// #bool true
```

### <a id="index"></a>Index

Index returns the rune index of the first occurrence of sub, or -1 if not found.

```go
v := str.Of("héllo").Index("llo")
println(v)
// #int 2
```

### <a id="lastindex"></a>LastIndex

LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.

```go
v := str.Of("go gophers go").LastIndex("go")
println(v)
// #int 10
```

### <a id="startswith"></a>StartsWith

StartsWith reports whether the string starts with any of the provided prefixes (case-sensitive).

```go
v := str.Of("gopher").StartsWith("go", "rust")
println(v)
// #bool true
```

### <a id="startswithfold"></a>StartsWithFold

StartsWithFold reports whether the string starts with any of the provided prefixes using Unicode case folding.

```go
v := str.Of("gopher").StartsWithFold("GO")
println(v)
// #bool true
```

## Slug

### <a id="slug"></a>Slug

Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
Not locale-aware; intended for identifiers/URLs.

```go
v := str.Of("Go Forj Toolkit").Slug("-").String()
println(v)
// #string go-forj-toolkit
```

## Snippet

### <a id="excerpt"></a>Excerpt

Excerpt returns a snippet around the first occurrence of needle with the given radius.
If needle is not found, an empty string is returned. If radius <= 0, a default of 100 is used.
Omission is used at the start/end when text is trimmed (default "...").

```go
v := str.Of("This is my name").Excerpt("my", 3, "...")
println(v.String())
// #string ...is my na...
```

## Split

### <a id="lines"></a>Lines

Lines splits the string into lines after normalizing newline variants.

```go
v := str.Of("a\\r\\nb\\nc").Lines()
println(v)
// #[]string [a b c]
```

### <a id="split"></a>Split

Split splits the string by the given separator.

```go
v := str.Of("a,b,c").Split(",")
println(v)
// #[]string [a b c]
```

### <a id="ucsplit"></a>UcSplit

UcSplit splits the string on uppercase boundaries and delimiters, returning segments.

```go
v := str.Of("HTTPRequestID").UcSplit()
println(v)
// #[]string [HTTP Request ID]
```

## Substrings

### <a id="after"></a>After

After returns the substring after the first occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("gopher::go").After("::").String()
println(v)
// #string go
```

### <a id="afterlast"></a>AfterLast

AfterLast returns the substring after the last occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("pkg/path/file.txt").AfterLast("/").String()
println(v)
// #string file.txt
```

### <a id="before"></a>Before

Before returns the substring before the first occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("gopher::go").Before("::").String()
println(v)
// #string gopher
```

### <a id="beforelast"></a>BeforeLast

BeforeLast returns the substring before the last occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("pkg/path/file.txt").BeforeLast("/").String()
println(v)
// #string pkg/path
```

### <a id="between"></a>Between

Between returns the substring between the first occurrence of start and the last occurrence of end.
Returns an empty string if either marker is missing or overlapping.

```go
v := str.Of("This is my name").Between("This", "name").String()
println(v)
// #string  is my
```

### <a id="betweenfirst"></a>BetweenFirst

BetweenFirst returns the substring between the first occurrence of start and the first occurrence of end after it.
Returns an empty string if markers are missing.

```go
v := str.Of("[a] bc [d]").BetweenFirst("[", "]").String()
println(v)
// #string a
```

### <a id="charat"></a>CharAt

CharAt returns the rune at the given index and true if within bounds.

```go
v, ok := str.Of("gopher").CharAt(2)
println(string(v), ok)
// #string p
// #bool true
```

### <a id="limit"></a>Limit

Limit truncates the string to length runes, appending suffix if truncation occurs.

```go
v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
println(v)
// #string Perfectly...
```

### <a id="slice"></a>Slice

Slice returns the substring between rune offsets [start:end).
Indices are clamped; if start >= end the result is empty.

```go
v := str.Of("naïve café").Slice(3, 7).String()
println(v)
// #string e ca
```

### <a id="substrreplace"></a>SubstrReplace

SubstrReplace replaces the rune slice in [start:end) with repl.

```go
v := str.Of("naïve café").SubstrReplace("i", 2, 3).String()
println(v)
// #string naive café
```

### <a id="take"></a>Take

Take returns the first length runes of the string (clamped).

```go
v := str.Of("gophers").Take(3).String()
println(v)
// #string gop
```

### <a id="takelast"></a>TakeLast

TakeLast returns the last length runes of the string (clamped).

```go
v := str.Of("gophers").TakeLast(4).String()
println(v)
// #string hers
```

## Transform

### <a id="repeat"></a>Repeat

Repeat repeats the string count times (non-negative).

```go
v := str.Of("go").Repeat(3).String()
println(v)
// #string gogogo
```

### <a id="reverse"></a>Reverse

Reverse returns a rune-safe reversed string.

```go
v := str.Of("naïve").Reverse().String()
println(v)
// #string evïan
```

### <a id="transliterate"></a>Transliterate

Transliterate replaces a small set of accented runes with ASCII equivalents.

```go
v := str.Of("café déjà vu").Transliterate().String()
println(v)
// #string cafe deja vu
```

## Words

### <a id="firstword"></a>FirstWord

FirstWord returns the first word token or empty string.

```go
v := str.Of("Hello world")
println(v.FirstWord().String())
// #string Hello
```

### <a id="join"></a>Join

Join joins the provided words with sep.

```go
v := str.Of("unused").Join([]string{"foo", "bar"}, "-").String()
println(v)
// #string foo-bar
```

### <a id="lastword"></a>LastWord

LastWord returns the last word token or empty string.

```go
v := str.Of("Hello world").LastWord().String()
println(v)
// #string world
```

### <a id="splitwords"></a>SplitWords

SplitWords splits the string into words (Unicode letters/digits runs).

```go
v := str.Of("one, two, three").SplitWords()
println(v)
// #[]string [one two three]
```

### <a id="wordcount"></a>WordCount

WordCount returns the number of word tokens (letters/digits runs).

```go
v := str.Of("Hello, world!").WordCount()
println(v)
// #int 2
```

### <a id="words"></a>Words

Words limits the string to count words, appending suffix if truncated.

```go
v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
println(v)
// #string Perfectly balanced as >>>
```

### <a id="wrapwords"></a>WrapWords

WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.

```go
v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
println(v)
// #string The quick brown fox\njumped over the lazy\ndog.
```
<!-- api:embed:end -->
