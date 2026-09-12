//go:build go1.27

package str

import "strings"

// cutLastStandard selects the real standard contract when the toolchain supports it.
var cutLastStandard = strings.CutLast
