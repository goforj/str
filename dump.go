package str

import "github.com/goforj/godump"

// dumpExample is a no-op wrapper to keep the godump import live for doc examples.
func dumpExample(values ...interface{}) {
	godump.Dump(values...)
}
