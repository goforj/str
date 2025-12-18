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
    <img src="https://img.shields.io/badge/tests-116-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
    <a href="https://goreportcard.com/report/github.com/goforj/str"><img src="https://goreportcard.com/badge/github.com/goforj/str" alt="Go Report Card"></a>
</p>


<!-- api:embed:start -->

## API Index

| Group | Functions |
|------:|-----------|
| **Affixes** | [ChopEnd](#chopend) [ChopStart](#chopstart) [EnsurePrefix](#ensureprefix) [EnsureSuffix](#ensuresuffix) [Unwrap](#unwrap) [Wrap](#wrap) |
| **Case** | [Camel](#camel) [Kebab](#kebab) [LcFirst](#lcfirst) [Pascal](#pascal) [Snake](#snake) [Title](#title) [ToLower](#tolower) [ToTitle](#totitle) [ToUpper](#toupper) [UcFirst](#ucfirst) |
| **Checks** | [IsBlank](#isblank) [IsEmpty](#isempty) |
| **Cleanup** | [Deduplicate](#deduplicate) [Squish](#squish) [Trim](#trim) [TrimLeft](#trimleft) [TrimRight](#trimright) |
| **Comparison** | [Equals](#equals) [EqualsFold](#equalsfold) |
| **Fluent** | [GoString](#gostring) [Of](#of) [String](#string) |
| **Length** | [Len](#len) [RuneCount](#runecount) |
| **Padding** | [PadBoth](#padboth) [PadLeft](#padleft) [PadRight](#padright) |
| **Replace** | [Remove](#remove) [ReplaceArray](#replacearray) [ReplaceFirst](#replacefirst) [ReplaceLast](#replacelast) [ReplaceMatches](#replacematches) [Swap](#swap) |
| **Search** | [Contains](#contains) [ContainsAll](#containsall) [ContainsAllFold](#containsallfold) [ContainsFold](#containsfold) [Count](#count) [EndsWith](#endswith) [EndsWithFold](#endswithfold) [Index](#index) [LastIndex](#lastindex) [StartsWith](#startswith) [StartsWithFold](#startswithfold) |
| **Slug** | [Slug](#slug) |
| **Substrings** | [After](#after) [AfterLast](#afterlast) [Before](#before) [BeforeLast](#beforelast) [Between](#between) [BetweenFirst](#betweenfirst) [CharAt](#charat) [Limit](#limit) [Slice](#slice) [Take](#take) [TakeLast](#takelast) |
| **Transform** | [Repeat](#repeat) [Reverse](#reverse) |
| **Words** | [FirstWord](#firstword) [Join](#join) [LastWord](#lastword) [SplitWords](#splitwords) [WordCount](#wordcount) [Words](#words) [WrapWords](#wrapwords) |


## Affixes

### <a id="chopend"></a>ChopEnd

ChopEnd removes the first matching suffix if present.

```go
v := str.Of("file.txt").ChopEnd(".txt", ".md").String()
godump.Dump(v)
// #string file
```

### <a id="chopstart"></a>ChopStart

ChopStart removes the first matching prefix if present.

```go
v := str.Of("https://goforj.dev").ChopStart("https://", "http://").String()
godump.Dump(v)
// #string goforj.dev
```

### <a id="ensureprefix"></a>EnsurePrefix

EnsurePrefix ensures the string starts with prefix, adding it if missing.

```go
v := str.Of("path/to").EnsurePrefix("/").String()
godump.Dump(v)
// #string /path/to
```

### <a id="ensuresuffix"></a>EnsureSuffix

EnsureSuffix ensures the string ends with suffix, adding it if missing.

```go
v := str.Of("path/to").EnsureSuffix("/").String()
godump.Dump(v)
// #string path/to/
```

### <a id="unwrap"></a>Unwrap

Unwrap removes matching before and after strings if present.

```go
v := str.Of("\"GoForj\"").Unwrap("\"", "\"").String()
godump.Dump(v)
// #string GoForj
```

### <a id="wrap"></a>Wrap

Wrap surrounds the string with before and after (after defaults to before).

```go
v := str.Of("GoForj").Wrap("\"", "").String()
godump.Dump(v)
// #string "GoForj"
```

## Case

### <a id="camel"></a>Camel

Camel converts the string to camelCase.

```go
v := str.Of("foo_bar baz").Camel().String()
godump.Dump(v)
// #string fooBarBaz
```

### <a id="kebab"></a>Kebab

Kebab converts the string to kebab-case.

```go
v := str.Of("fooBar baz").Kebab().String()
godump.Dump(v)
// #string foo-bar-baz
```

### <a id="lcfirst"></a>LcFirst

LcFirst returns the string with the first rune lower-cased.

```go
v := str.Of("Gopher")
godump.Dump(v)
// #string gopher
```

### <a id="pascal"></a>Pascal

Pascal converts the string to PascalCase.

```go
v := str.Of("foo_bar baz")
godump.Dump(v)
// #string FooBarBaz
```

### <a id="snake"></a>Snake

Snake converts the string to snake_case using the provided separator (default "_").

```go
v := str.Of("fooBar baz").Snake("_").String()
godump.Dump(v)
// #string foo_bar_baz
```

### <a id="title"></a>Title

Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.

```go
v := str.Of("a nice title uses the correct case").Title().String()
godump.Dump(v)
// #string A Nice Title Uses The Correct Case
```

### <a id="tolower"></a>ToLower

ToLower returns a lowercase copy of the string using Unicode rules.

```go
v := str.Of("GoLang").ToLower().String()
godump.Dump(v)
// #string golang
```

### <a id="totitle"></a>ToTitle

ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.

```go
v := str.Of("ß").ToTitle().String()
godump.Dump(v)
// #string SS
```

### <a id="toupper"></a>ToUpper

ToUpper returns an uppercase copy of the string using Unicode rules.

```go
v := str.Of("GoLang").ToUpper().String()
godump.Dump(v)
// #string GOLANG
```

### <a id="ucfirst"></a>UcFirst

UcFirst returns the string with the first rune upper-cased.

```go
v := str.Of("gopher").UcFirst().String()
godump.Dump(v)
// #string Gopher
```

## Checks

### <a id="isblank"></a>IsBlank

IsBlank reports whether the string contains only Unicode whitespace.

```go
v := str.Of("  \\t\\n")
godump.Dump(v.IsBlank())
// #bool true
```

### <a id="isempty"></a>IsEmpty

IsEmpty reports whether the string has zero length.

```go
v := str.Of("").IsEmpty()
godump.Dump(v)
// #bool true
```

## Cleanup

### <a id="deduplicate"></a>Deduplicate

Deduplicate collapses consecutive instances of char into a single instance.
If char is zero, space is used.

```go
v := str.Of("The   Go   Playground").Deduplicate(' ').String()
godump.Dump(v)
// #string The Go Playground
```

### <a id="squish"></a>Squish

Squish trims leading/trailing whitespace and collapses internal whitespace to single spaces.

```go
v := str.Of("   go   forj  ").Squish().String()
godump.Dump(v)
// #string go forj
```

### <a id="trim"></a>Trim

Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").Trim("").String()
godump.Dump(v)
// #string GoForj
```

### <a id="trimleft"></a>TrimLeft

TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").TrimLeft("").String()
godump.Dump(v)
// #string GoForj
```

### <a id="trimright"></a>TrimRight

TrimRight trims trailing characters. If cutset is empty, trims Unicode whitespace.

```go
v := str.Of("  GoForj  ").TrimRight("").String()
godump.Dump(v)
// #string   GoForj
```

## Comparison

### <a id="equals"></a>Equals

Equals reports whether the string exactly matches other (case-sensitive).

```go
v := str.Of("gopher").Equals("gopher")
godump.Dump(v)
// #bool true
```

### <a id="equalsfold"></a>EqualsFold

EqualsFold reports whether the string matches other using Unicode case folding.

```go
v := str.Of("gopher").EqualsFold("GOPHER")
godump.Dump(v)
// #bool true
```

## Fluent

### <a id="gostring"></a>GoString

GoString allows %#v formatting to print the raw string.

```go
v := str.Of("go")
godump.Dump(fmt.Sprintf("%#v", v))
// #string go
```

### <a id="of"></a>Of

Of wraps a raw string with fluent helpers.

```go
v := str.Of("gopher")
godump.Dump(v.String())
// #string gopher
```

### <a id="string"></a>String

String returns the underlying raw string value.

```go
v := str.Of("go").String()
godump.Dump(v)
// #string go
```

## Length

### <a id="len"></a>Len

Len returns the number of runes in the string.

```go
v := str.Of("gophers 🦫").Len()
godump.Dump(v)
// #int 9
```

### <a id="runecount"></a>RuneCount

RuneCount is an alias for Len to make intent explicit in mixed codebases.

```go
v := str.Of("naïve").RuneCount()
godump.Dump(v)
// #int 5
```

## Padding

### <a id="padboth"></a>PadBoth

PadBoth pads the string on both sides to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadBoth(6, "-").String()
godump.Dump(v)
// #string --go--
```

### <a id="padleft"></a>PadLeft

PadLeft pads the string on the left to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadLeft(5, " ").String()
godump.Dump(v)
// #string \u00a0\u00a0\u00a0go
```

### <a id="padright"></a>PadRight

PadRight pads the string on the right to reach length runes using pad (defaults to space).

```go
v := str.Of("go").PadRight(5, ".").String()
godump.Dump(v)
// #string go...
```

## Replace

### <a id="remove"></a>Remove

Remove deletes all occurrences of provided substrings.

```go
v := str.Of("The Go Toolkit").Remove("Go ").String()
godump.Dump(v)
// #string The Toolkit
```

### <a id="replacearray"></a>ReplaceArray

ReplaceArray replaces all occurrences of each old in olds with repl.

```go
v := str.Of("The---Go---Toolkit")
godump.Dump(v.ReplaceArray([]string{"---"}, "-").String())
// #string The-Go-Toolkit
```

### <a id="replacefirst"></a>ReplaceFirst

ReplaceFirst replaces the first occurrence of old with repl.

```go
v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
godump.Dump(v)
// #string go gopher
```

### <a id="replacelast"></a>ReplaceLast

ReplaceLast replaces the last occurrence of old with repl.

```go
v := str.Of("gopher gopher").ReplaceLast("gopher", "go").String()
godump.Dump(v)
// #string gopher go
```

### <a id="replacematches"></a>ReplaceMatches

ReplaceMatches applies repl to each regex match and returns the result.

```go
re := regexp.MustCompile(`\d+`)
v := str.Of("Hello 123 World").ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String()
godump.Dump(v)
// #string Hello [123] World
```

### <a id="swap"></a>Swap

Swap replaces multiple values using strings.Replacer built from a map.

```go
pairs := map[string]string{"Gophers": "GoForj", "great": "fantastic"}
v := str.Of("Gophers are great!").Swap(pairs).String()
godump.Dump(v)
// #string GoForj are fantastic!
```

## Search

### <a id="contains"></a>Contains

Contains reports whether the string contains any of the provided substrings (case-sensitive).
Empty substrings return true to match strings.Contains semantics.

```go
v := str.Of("Go means gophers").Contains("rust", "gopher")
godump.Dump(v)
// #bool true
```

### <a id="containsall"></a>ContainsAll

ContainsAll reports whether the string contains all provided substrings (case-sensitive).
Empty substrings are ignored.

```go
v := str.Of("Go means gophers").ContainsAll("Go", "gopher")
godump.Dump(v)
// #bool true
```

### <a id="containsallfold"></a>ContainsAllFold

ContainsAllFold reports whether the string contains all provided substrings, using Unicode
case folding for comparisons.
Empty substrings are ignored.

```go
v := str.Of("Go means gophers").ContainsAllFold("go", "GOPHER")
godump.Dump(v)
// #bool true
```

### <a id="containsfold"></a>ContainsFold

ContainsFold reports whether the string contains any of the provided substrings, using Unicode
case folding for comparisons.
Empty substrings return true.

```go
v := str.Of("Go means gophers").ContainsFold("GOPHER", "rust")
godump.Dump(v)
// #bool true
```

### <a id="count"></a>Count

Count returns the number of non-overlapping occurrences of sub.

```go
v := str.Of("gogophergo").Count("go")
godump.Dump(v)
// #int 3
```

### <a id="endswith"></a>EndsWith

EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).

```go
v := str.Of("gopher").EndsWith("her", "cat")
godump.Dump(v)
// #bool true
```

### <a id="endswithfold"></a>EndsWithFold

EndsWithFold reports whether the string ends with any of the provided suffixes using Unicode case folding.

```go
v := str.Of("gopher").EndsWithFold("HER")
godump.Dump(v)
// #bool true
```

### <a id="index"></a>Index

Index returns the rune index of the first occurrence of sub, or -1 if not found.

```go
v := str.Of("héllo").Index("llo")
godump.Dump(v)
// #int 2
```

### <a id="lastindex"></a>LastIndex

LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.

```go
v := str.Of("go gophers go").LastIndex("go")
godump.Dump(v)
// #int 10
```

### <a id="startswith"></a>StartsWith

StartsWith reports whether the string starts with any of the provided prefixes (case-sensitive).

```go
v := str.Of("gopher").StartsWith("go", "rust")
godump.Dump(v)
// #bool true
```

### <a id="startswithfold"></a>StartsWithFold

StartsWithFold reports whether the string starts with any of the provided prefixes using Unicode case folding.

```go
v := str.Of("gopher").StartsWithFold("GO")
godump.Dump(v)
// #bool true
```

## Slug

### <a id="slug"></a>Slug

Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
Not locale-aware; intended for identifiers/URLs.

```go
v := str.Of("Go Forj Toolkit").Slug("-").String()
godump.Dump(v)
// #string go-forj-toolkit
```

## Substrings

### <a id="after"></a>After

After returns the substring after the first occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("gopher::go").After("::").String()
godump.Dump(v)
// #string go
```

### <a id="afterlast"></a>AfterLast

AfterLast returns the substring after the last occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("pkg/path/file.txt").AfterLast("/").String()
godump.Dump(v)
// #string file.txt
```

### <a id="before"></a>Before

Before returns the substring before the first occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("gopher::go").Before("::").String()
godump.Dump(v)
// #string gopher
```

### <a id="beforelast"></a>BeforeLast

BeforeLast returns the substring before the last occurrence of sep.
If sep is empty or not found, the original string is returned.

```go
v := str.Of("pkg/path/file.txt").BeforeLast("/").String()
godump.Dump(v)
// #string pkg/path
```

### <a id="between"></a>Between

Between returns the substring between the first occurrence of start and the last occurrence of end.
Returns an empty string if either marker is missing or overlapping.

```go
v := str.Of("This is my name").Between("This", "name").String()
godump.Dump(v)
// #string  is my
```

### <a id="betweenfirst"></a>BetweenFirst

BetweenFirst returns the substring between the first occurrence of start and the first occurrence of end after it.
Returns an empty string if markers are missing.

```go
v := str.Of("[a] bc [d]").BetweenFirst("[", "]").String()
godump.Dump(v)
// #string a
```

### <a id="charat"></a>CharAt

CharAt returns the rune at the given index and true if within bounds.

```go
v, ok := str.Of("gopher").CharAt(2)
godump.Dump(string(v), ok)
// #string p
// #bool true
```

### <a id="limit"></a>Limit

Limit truncates the string to length runes, appending suffix if truncation occurs.

```go
v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
godump.Dump(v)
// #string Perfectly...
```

### <a id="slice"></a>Slice

Slice returns the substring between rune offsets [start:end).
Indices are clamped; if start >= end the result is empty.

```go
v := str.Of("naïve café").Slice(3, 7).String()
godump.Dump(v)
// #string e ca
```

### <a id="take"></a>Take

Take returns the first length runes of the string (clamped).

```go
v := str.Of("gophers").Take(3).String()
godump.Dump(v)
// #string gop
```

### <a id="takelast"></a>TakeLast

TakeLast returns the last length runes of the string (clamped).

```go
v := str.Of("gophers").TakeLast(4).String()
godump.Dump(v)
// #string hers
```

## Transform

### <a id="repeat"></a>Repeat

Repeat repeats the string count times (non-negative).

```go
v := str.Of("go").Repeat(3).String()
godump.Dump(v)
// #string gogogo
```

### <a id="reverse"></a>Reverse

Reverse returns a rune-safe reversed string.

```go
v := str.Of("naïve").Reverse().String()
godump.Dump(v)
// #string evïan
```

## Words

### <a id="firstword"></a>FirstWord

FirstWord returns the first word token or empty string.

```go
v := str.Of("Hello world")
godump.Dump(v.FirstWord().String())
// #string Hello
```

### <a id="join"></a>Join

Join joins the provided words with sep.

```go
v := str.Of("unused").Join([]string{"foo", "bar"}, "-").String()
godump.Dump(v)
// #string foo-bar
```

### <a id="lastword"></a>LastWord

LastWord returns the last word token or empty string.

```go
v := str.Of("Hello world").LastWord().String()
godump.Dump(v)
// #string world
```

### <a id="splitwords"></a>SplitWords

SplitWords splits the string into words (Unicode letters/digits runs).

```go
v := str.Of("one, two, three").SplitWords()
godump.Dump(v)
// #[]string [one two three]
```

### <a id="wordcount"></a>WordCount

WordCount returns the number of word tokens (letters/digits runs).

```go
v := str.Of("Hello, world!").WordCount()
godump.Dump(v)
// #int 2
```

### <a id="words"></a>Words

Words limits the string to count words, appending suffix if truncated.

```go
v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
godump.Dump(v)
// #string Perfectly balanced as >>>
```

### <a id="wrapwords"></a>WrapWords

WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.

```go
v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
godump.Dump(v)
// #string The quick brown fox\njumped over the lazy\ndog.
```
<!-- api:embed:end -->