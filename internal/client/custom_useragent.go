package client

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/growthbook/cli/v2/internal/customcfg"
)

// UserAgent identifies CLI traffic to the GrowthBook API. The generated SDK
// otherwise sends Speakeasy's default, which carries the generator's version
// rather than the CLI's and reads as a library instead of a release.
func UserAgent() string {
	return fmt.Sprintf(
		"growthbook-cli/%s (%s; %s/%s)",
		customcfg.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH,
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
