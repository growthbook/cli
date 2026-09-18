package client

import (
	"net/http"
	"strings"

	"github.com/growthbook/cli/v2/internal/customcfg"
)

// sdkDefaultUserAgentPrefix marks the generated SDK's own User-Agent, the one value replaced outright.
const sdkDefaultUserAgentPrefix = "speakeasy-sdk/"

type userAgentClient struct {
	inner HTTPClient
}

func (c *userAgentClient) Do(req *http.Request) (*http.Response, error) {
	ua := customcfg.UserAgent()
	// The API classifies callers on the leading token, so a `-H User-Agent` goes behind ours.
	if existing := req.Header.Get("User-Agent"); existing != "" && !strings.HasPrefix(existing, sdkDefaultUserAgentPrefix) {
		ua += " " + existing
	}
	req.Header.Set("User-Agent", ua)
	return c.inner.Do(req)
}

// WrapClientForUserAgent puts the CLI's product token at the front of every
// request's User-Agent. Wrap it outside the diagnostics client so `--debug`
// shows what actually gets sent.
func WrapClientForUserAgent(inner HTTPClient) HTTPClient {
	return &userAgentClient{inner: inner}
}
