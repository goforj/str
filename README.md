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
    <img src="https://img.shields.io/badge/tests-104-brightgreen" alt="Tests">
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
| **Length** | [Len](#len) [RuneCount](#runecount) |
| **Other** | [GoString](#gostring) [Of](#of) [String](#string) |
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

_Example: chop end_

```go
val := str.Of("file.txt")
godump.Dump(val.ChopEnd(".txt", ".md").String())

// #string file
```

### <a id="chopstart"></a>ChopStart

ChopStart removes the first matching prefix if present.

_Example: chop start_

```go
val := str.Of("https://goforj.dev")
godump.Dump(val.ChopStart("https://", "http://").String())

// #string goforj.dev
```

### <a id="ensureprefix"></a>EnsurePrefix

EnsurePrefix ensures the string starts with prefix, adding it if missing.

_Example: ensure prefix_

```go
val := str.Of("path/to")
godump.Dump(val.EnsurePrefix("/").String())

// #string /path/to
```

### <a id="ensuresuffix"></a>EnsureSuffix

EnsureSuffix ensures the string ends with suffix, adding it if missing.

_Example: ensure suffix_

```go
val := str.Of("path/to")
godump.Dump(val.EnsureSuffix("/").String())

// #string path/to/
```

### <a id="unwrap"></a>Unwrap

Unwrap removes matching before and after strings if present.

_Example: unwrap string_

```go
val := str.Of("\"Laravel\"")
godump.Dump(val.Unwrap("\"", "\"").String())

// #string Laravel
```

### <a id="wrap"></a>Wrap

Wrap surrounds the string with before and after (after defaults to before).

_Example: wrap string_

```go
val := str.Of("Laravel")
godump.Dump(val.Wrap("\"", "").String())

// #string "Laravel"
```

## Case

### <a id="camel"></a>Camel

Camel converts the string to camelCase.

_Example: camel case_

```go
val := str.Of("foo_bar baz")
godump.Dump(val.Camel().String())

// #string fooBarBaz
```

### <a id="kebab"></a>Kebab

Kebab converts the string to kebab-case.

_Example: kebab case_

```go
val := str.Of("fooBar baz")
godump.Dump(val.Kebab().String())

// #string foo-bar-baz
```

### <a id="lcfirst"></a>LcFirst

LcFirst returns the string with the first rune lower-cased.

_Example: lowercase first rune_

```go
val := str.Of("Gopher")
godump.Dump(val.LcFirst().String())

// #string gopher
```

### <a id="pascal"></a>Pascal

Pascal converts the string to PascalCase.

_Example: pascal case_

```go
val := str.Of("foo_bar baz")
godump.Dump(val.Pascal().String())

// #string FooBarBaz
```

### <a id="snake"></a>Snake

Snake converts the string to snake_case using the provided separator (default "_").

_Example: snake case_

```go
val := str.Of("fooBar baz")
godump.Dump(val.Snake("_").String())

// #string foo_bar_baz
```

### <a id="title"></a>Title

Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.

_Example: title case words_

```go
val := str.Of("a nice title uses the correct case")
godump.Dump(val.Title().String())

// #string A Nice Title Uses The Correct Case
```

### <a id="tolower"></a>ToLower

ToLower returns a lowercase copy of the string using Unicode rules.

_Example: lowercase text_

```go
val := str.Of("GoLang")
godump.Dump(val.ToLower().String())

// #string golang
```

### <a id="totitle"></a>ToTitle

ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.

_Example: title map runes_

```go
val := str.Of("ß")
godump.Dump(val.ToTitle().String())

// #string SS
```

### <a id="toupper"></a>ToUpper

ToUpper returns an uppercase copy of the string using Unicode rules.

_Example: uppercase text_

```go
val := str.Of("GoLang")
godump.Dump(val.ToUpper().String())

// #string GOLANG
```

### <a id="ucfirst"></a>UcFirst

UcFirst returns the string with the first rune upper-cased.

_Example: uppercase first rune_

```go
val := str.Of("gopher")
godump.Dump(val.UcFirst().String())

// #string Gopher
```

## Checks

### <a id="isblank"></a>IsBlank

IsBlank reports whether the string contains only Unicode whitespace.

_Example: blank check_

```go
val := str.Of("  \\t\\n")
godump.Dump(val.IsBlank())

// #bool true
```

### <a id="isempty"></a>IsEmpty

IsEmpty reports whether the string has zero length.

_Example: empty check_

```go
val := str.Of("")
godump.Dump(val.IsEmpty())

// #bool true
```

## Cleanup

### <a id="deduplicate"></a>Deduplicate

Deduplicate collapses consecutive instances of char into a single instance.
If char is zero, space is used.

_Example: collapse spaces_

```go
val := str.Of("The   Laravel   Framework")
godump.Dump(val.Deduplicate(' ').String())

// #string The Laravel Framework
```

### <a id="squish"></a>Squish

Squish trims leading/trailing whitespace and collapses internal whitespace to single spaces.

_Example: squish whitespace_

```go
val := str.Of("   go   forj  ")
godump.Dump(val.Squish().String())

// #string go forj
```

### <a id="trim"></a>Trim

Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.

_Example: trim whitespace_

```go
val := str.Of("  Laravel  ")
godump.Dump(val.Trim("").String())

// #string Laravel
```

### <a id="trimleft"></a>TrimLeft

TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.

_Example: trim left_

```go
val := str.Of("  Laravel  ")
godump.Dump(val.TrimLeft("").String())

// #string Laravel
```

### <a id="trimright"></a>TrimRight

TrimRight trims trailing characters. If cutset is empty, trims Unicode whitespace.

_Example: trim right_

```go
val := str.Of("  Laravel  ")
godump.Dump(val.TrimRight("").String())

// #string   Laravel
```

## Comparison

### <a id="equals"></a>Equals

Equals reports whether the string exactly matches other (case-sensitive).

_Example: exact match_

```go
val := str.Of("gopher")
godump.Dump(val.Equals("gopher"))

// #bool true
```

### <a id="equalsfold"></a>EqualsFold

EqualsFold reports whether the string matches other using Unicode case folding.

_Example: case-insensitive match_

```go
val := str.Of("gopher")
godump.Dump(val.EqualsFold("GOPHER"))

// #bool true
```

## Length

### <a id="len"></a>Len

Len returns the number of runes in the string.

_Example: count runes instead of bytes_

```go
val := str.Of("gophers 🦫")
godump.Dump(val.Len())

// #int 9
```

### <a id="runecount"></a>RuneCount

RuneCount is an alias for Len to make intent explicit in mixed codebases.

_Example: alias for Len_

```go
val := str.Of("naïve")
godump.Dump(val.RuneCount())

// #int 5
```

## Other

### <a id="gostring"></a>GoString

GoString allows %#v formatting to print the raw string.

### <a id="of"></a>Of

Of wraps a raw string with fluent helpers.

### <a id="string"></a>String

String returns the underlying raw string value.

## Padding

### <a id="padboth"></a>PadBoth

PadBoth pads the string on both sides to reach length runes using pad (defaults to space).

_Example: pad both_

```go
val := str.Of("go")
godump.Dump(val.PadBoth(6, "-").String())

// #string --go--
```

### <a id="padleft"></a>PadLeft

PadLeft pads the string on the left to reach length runes using pad (defaults to space).

_Example: pad left_

```go
val := str.Of("go")
godump.Dump(val.PadLeft(5, " ").String())

// #string \u00a0\u00a0\u00a0go
```

### <a id="padright"></a>PadRight

PadRight pads the string on the right to reach length runes using pad (defaults to space).

_Example: pad right_

```go
val := str.Of("go")
godump.Dump(val.PadRight(5, ".").String())

// #string go...
```

## Replace

### <a id="remove"></a>Remove

Remove deletes all occurrences of provided substrings.

_Example: remove substrings_

```go
val := str.Of("The Laravel Framework")
godump.Dump(val.Remove("Laravel ").String())

// #string The Framework
```

### <a id="replacearray"></a>ReplaceArray

ReplaceArray replaces all occurrences of each old in olds with repl.

_Example: replace many_

```go
val := str.Of("The---Laravel---Framework")
godump.Dump(val.ReplaceArray([]string{"---"}, "-").String())

// #string The-Laravel-Framework
```

### <a id="replacefirst"></a>ReplaceFirst

ReplaceFirst replaces the first occurrence of old with repl.

_Example: replace first_

```go
val := str.Of("gopher gopher")
godump.Dump(val.ReplaceFirst("gopher", "go").String())

// #string go gopher
```

### <a id="replacelast"></a>ReplaceLast

ReplaceLast replaces the last occurrence of old with repl.

_Example: replace last_

```go
val := str.Of("gopher gopher")
godump.Dump(val.ReplaceLast("gopher", "go").String())

// #string gopher go
```

### <a id="replacematches"></a>ReplaceMatches

ReplaceMatches applies repl to each regex match and returns the result.

_Example: regex replace with callback_

```go
val := str.Of("Hello 123 World")
re := regexp.MustCompile(`\d+`)
godump.Dump(val.ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String())

// #string Hello [123] World
```

### <a id="swap"></a>Swap

Swap replaces multiple values using strings.Replacer built from a map.

_Example: swap map_

```go
val := str.Of("Tacos are great!")
pairs := map[string]string{"Tacos": "Burritos", "great": "fantastic"}
godump.Dump(val.Swap(pairs).String())

// #string Burritos are fantastic!
```

## Search

### <a id="contains"></a>Contains

Contains reports whether the string contains any of the provided substrings (case-sensitive).
Empty substrings return true to match strings.Contains semantics.

_Example: contains any_

```go
val := str.Of("Go means gophers")
godump.Dump(val.Contains("rust", "gopher"))

// #bool true
```

### <a id="containsall"></a>ContainsAll

ContainsAll reports whether the string contains all provided substrings (case-sensitive).
Empty substrings are ignored.

_Example: contains all_

```go
val := str.Of("Go means gophers")
godump.Dump(val.ContainsAll("Go", "gopher"))

// #bool true
```

### <a id="containsallfold"></a>ContainsAllFold

ContainsAllFold reports whether the string contains all provided substrings, using Unicode
case folding for comparisons.
Empty substrings are ignored.

_Example: contains all (case-insensitive)_

```go
val := str.Of("Go means gophers")
godump.Dump(val.ContainsAllFold("go", "GOPHER"))

// #bool true
```

### <a id="containsfold"></a>ContainsFold

ContainsFold reports whether the string contains any of the provided substrings, using Unicode
case folding for comparisons.
Empty substrings return true.

_Example: contains any (case-insensitive)_

```go
val := str.Of("Go means gophers")
godump.Dump(val.ContainsFold("GOPHER", "rust"))

// #bool true
```

### <a id="count"></a>Count

Count returns the number of non-overlapping occurrences of sub.

_Example: count substring_

```go
val := str.Of("gogophergo")
godump.Dump(val.Count("go"))

// #int 3
```

### <a id="endswith"></a>EndsWith

EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).

_Example: ends with any_

```go
val := str.Of("gopher")
godump.Dump(val.EndsWith("her", "cat"))

// #bool true
```

### <a id="endswithfold"></a>EndsWithFold

EndsWithFold reports whether the string ends with any of the provided suffixes using Unicode case folding.

_Example: ends with (case-insensitive)_

```go
val := str.Of("gopher")
godump.Dump(val.EndsWithFold("HER"))

// #bool true
```

### <a id="index"></a>Index

Index returns the rune index of the first occurrence of sub, or -1 if not found.

_Example: first rune index_

```go
val := str.Of("héllo")
godump.Dump(val.Index("llo"))

// #int 2
```

### <a id="lastindex"></a>LastIndex

LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.

_Example: last rune index_

```go
val := str.Of("go gophers go")
godump.Dump(val.LastIndex("go"))

// #int 10
```

### <a id="startswith"></a>StartsWith

StartsWith reports whether the string starts with any of the provided prefixes (case-sensitive).

_Example: starts with any_

```go
val := str.Of("gopher")
godump.Dump(val.StartsWith("go", "rust"))

// #bool true
```

### <a id="startswithfold"></a>StartsWithFold

StartsWithFold reports whether the string starts with any of the provided prefixes using Unicode case folding.

_Example: starts with (case-insensitive)_

```go
val := str.Of("gopher")
godump.Dump(val.StartsWithFold("GO"))

// #bool true
```

## Slug

### <a id="slug"></a>Slug

Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
Not locale-aware; intended for identifiers/URLs.

_Example: build slug_

```go
val := str.Of("Laravel Framework")
godump.Dump(val.Slug("-").String())

// #string laravel-framework
```

## Substrings

### <a id="after"></a>After

After returns the substring after the first occurrence of sep.
If sep is empty or not found, the original string is returned.

_Example: slice after marker_

```go
val := str.Of("gopher::go")
godump.Dump(val.After("::").String())

// #string go
```

### <a id="afterlast"></a>AfterLast

AfterLast returns the substring after the last occurrence of sep.
If sep is empty or not found, the original string is returned.

_Example: slice after last separator_

```go
val := str.Of("pkg/path/file.txt")
godump.Dump(val.AfterLast("/").String())

// #string file.txt
```

### <a id="before"></a>Before

Before returns the substring before the first occurrence of sep.
If sep is empty or not found, the original string is returned.

_Example: slice before marker_

```go
val := str.Of("gopher::go")
godump.Dump(val.Before("::").String())

// #string gopher
```

### <a id="beforelast"></a>BeforeLast

BeforeLast returns the substring before the last occurrence of sep.
If sep is empty or not found, the original string is returned.

_Example: slice before last separator_

```go
val := str.Of("pkg/path/file.txt")
godump.Dump(val.BeforeLast("/").String())

// #string pkg/path
```

### <a id="between"></a>Between

Between returns the substring between the first occurrence of start and the last occurrence of end.
Returns an empty string if either marker is missing or overlapping.

_Example: between first and last_

```go
val := str.Of("This is my name")
godump.Dump(val.Between("This", "name").String())

// #string  is my
```

### <a id="betweenfirst"></a>BetweenFirst

BetweenFirst returns the substring between the first occurrence of start and the first occurrence of end after it.
Returns an empty string if markers are missing.

_Example: minimal span between markers_

```go
val := str.Of("[a] bc [d]")
godump.Dump(val.BetweenFirst("[", "]").String())

// #string a
```

### <a id="charat"></a>CharAt

CharAt returns the rune at the given index and true if within bounds.

_Example: rune lookup_

```go
val := str.Of("gopher")
r, ok := val.CharAt(2)
godump.Dump(string(r), ok)

// #string p
// #bool true
```

### <a id="limit"></a>Limit

Limit truncates the string to length runes, appending suffix if truncation occurs.

_Example: limit with suffix_

```go
val := str.Of("Perfectly balanced, as all things should be.")
godump.Dump(val.Limit(10, "...").String())

// #string Perfectly...
```

### <a id="slice"></a>Slice

Slice returns the substring between rune offsets [start:end).
Indices are clamped; if start >= end the result is empty.

_Example: rune-safe slice_

```go
val := str.Of("naïve café")
godump.Dump(val.Slice(3, 7).String())

// #string e ca
```

### <a id="take"></a>Take

Take returns the first length runes of the string (clamped).

_Example: take head_

```go
val := str.Of("gophers")
godump.Dump(val.Take(3).String())

// #string gop
```

### <a id="takelast"></a>TakeLast

TakeLast returns the last length runes of the string (clamped).

_Example: take tail_

```go
val := str.Of("gophers")
godump.Dump(val.TakeLast(4).String())

// #string hers
```

## Transform

### <a id="repeat"></a>Repeat

Repeat repeats the string count times (non-negative).

_Example: repeat string_

```go
val := str.Of("go")
godump.Dump(val.Repeat(3).String())

// #string gogogo
```

### <a id="reverse"></a>Reverse

Reverse returns a rune-safe reversed string.

_Example: reverse_

```go
val := str.Of("naïve")
godump.Dump(val.Reverse().String())

// #string evïan
```

## Words

### <a id="firstword"></a>FirstWord

FirstWord returns the first word token or empty string.

_Example: first word_

```go
val := str.Of("Hello world")
godump.Dump(val.FirstWord().String())

// #string Hello
```

### <a id="join"></a>Join

Join joins the provided words with sep.

_Example: join words_

```go
val := str.Of("unused")
godump.Dump(val.Join([]string{"foo", "bar"}, "-").String())

// #string foo-bar
```

### <a id="lastword"></a>LastWord

LastWord returns the last word token or empty string.

_Example: last word_

```go
val := str.Of("Hello world")
godump.Dump(val.LastWord().String())

// #string world
```

### <a id="splitwords"></a>SplitWords

SplitWords splits the string into words (Unicode letters/digits runs).

_Example: split words_

```go
val := str.Of("one, two, three")
godump.Dump(val.SplitWords())

// #[]string [one two three]
```

### <a id="wordcount"></a>WordCount

WordCount returns the number of word tokens (letters/digits runs).

_Example: count words_

```go
val := str.Of("Hello, world!")
godump.Dump(val.WordCount())

// #int 2
```

### <a id="words"></a>Words

Words limits the string to count words, appending suffix if truncated.

_Example: limit words_

```go
val := str.Of("Perfectly balanced, as all things should be.")
godump.Dump(val.Words(3, " >>>").String())

// #string Perfectly balanced as >>>
```

### <a id="wrapwords"></a>WrapWords

WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.

_Example: wrap words_

```go
val := str.Of("The quick brown fox jumped over the lazy dog.")
godump.Dump(val.WrapWords(20, "\n").String())

// #string The quick brown fox\njumped over the lazy\ndog.
```
<!-- api:embed:end -->