// User-Agent for CLI traffic (hand-written, NOT generated).
package customcfg

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// version reports the CLI release. goreleaser stamps Version via ldflags; a
// `go install module@vX` build has no ldflags but does record the module
// version in its build info, so read that before giving up on "dev".
func version() string {
	if Version != "" && Version != "dev" {
		return Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return strings.TrimPrefix(v, "v")
		}
	}
	return "dev"
}

// UserAgent identifies CLI traffic; the API classifies first-party clients on
// the leading `growthbook-<product>/<version>` token.
func UserAgent() string {
	return fmt.Sprintf(
		"growthbook-cli/%s (%s; %s/%s)",
		version(), runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)
}
