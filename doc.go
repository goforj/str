// Package str provides an immutable fluent wrapper for practical Unicode-aware
// string operations that are awkward or repetitive with Go's standard library.
//
// Names follow the Go standard library where an equivalent concept exists.
// Empty search terms are treated as no match, replacement operations are no-ops
// for empty search terms, and positions, lengths, and widths are measured in
// Unicode code points unless a method explicitly documents otherwise.
//
// String is intended for short-lived transformation chains. Call [String.String]
// to store or serialize the resulting built-in string value.
package str
