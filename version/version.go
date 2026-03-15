package version

import (
	"fmt"
	"runtime/debug"
)

// Set via -ldflags at build time.
var (
	Version = "dev"
	Commit  = "unknown"
)

func String() string {
	s := fmt.Sprintf("kne version %s (commit %s)", Version, Commit)
	if info, ok := debug.ReadBuildInfo(); ok {
		s += fmt.Sprintf(" go %s", info.GoVersion)
	}
	return s
}
