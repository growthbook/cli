package client

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/growthbook/cli/v2/internal/customcfg"
)

// version reports the CLI release. goreleaser stamps customcfg.Version via
// ldflags; a `go install module@vX` build has no ldflags but does record the
// module version in its build info, so read that before giving up on "dev".
// `internal/cli`.Version isn't reachable here — that package imports this one.
func version() string {
	if customcfg.Version != "" && customcfg.Version != "dev" {
		return customcfg.Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return strings.TrimPrefix(v, "v")
		}
	}
	return "dev"
}

// UserAgent identifies CLI traffic to the GrowthBook API. The generated SDK
// otherwise sends Speakeasy's default, which carries the generator's version
// rather than the CLI's and reads as a library instead of a release.
func UserAgent() string {
	return fmt.Sprintf(
		"growthbook-cli/%s (%s; %s/%s)",
		version(), runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)
}

type userAgentClient struct {
	inner HTTPClient
}

func (c *userAgentClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", UserAgent())
	return c.inner.Do(req)
}

// WrapClientForUserAgent overwrites the SDK's User-Agent on the way out. Wrap it
// outside the diagnostics client so `--debug` shows what actually gets sent.
func WrapClientForUserAgent(inner HTTPClient) HTTPClient {
	return &userAgentClient{inner: inner}
}
