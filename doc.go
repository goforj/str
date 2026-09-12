// Package str provides immutable fluent string operations with Go strings contracts.
//
// Standard-named methods replace the source string argument with a String receiver.
// Remaining arguments keep their order; a single string result becomes String.
// Multiple results, slices, and iterators keep their standard types. Join is a
// package-level constructor because it has no source string to receive.
//
// Searches return byte offsets and follow standard empty-search behavior.
// Additional application helpers such as Slice, Take, and padding use Unicode
// code points where documented. Their rune offsets are not Index byte offsets.
//
// The string-function surface includes Go 1.27's CutLast while retaining Go 1.24
// support. Stateful Builder, Reader, and Replacer types remain in package strings.
//
// String is intended for short-lived transformation chains. Call [String.String]
// to store or serialize the resulting built-in string value.
package str
