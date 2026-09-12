//go:build !go1.27

package str

// cutLastStandard keeps the same test inventory runnable on the Go 1.24 minimum.
var cutLastStandard = referenceCutLast
