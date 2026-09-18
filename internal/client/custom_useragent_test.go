package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/growthbook/cli/v2/internal/customcfg"
)

// sends reports the User-Agent the wrapped client actually puts on the wire,
// given the one the request already carries ("" for none).
func sends(t *testing.T, userAgent string) string {
	t.Helper()
	var got string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.UserAgent()
	}))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	if userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}

	res, err := WrapClientForUserAgent(&http.Client{}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	return got
}

func TestUserAgentReplacesSDKDefault(t *testing.T) {
	// What the generated SDK sets before handing the request to the client.
	got := sends(t, "speakeasy-sdk/go 0.0.1 2.930.0 5.0.1 github.com/growthbook/cli/v2/internal/sdk")
	if want := customcfg.UserAgent(); got != want {
		t.Errorf("sent User-Agent = %q, want %q", got, want)
	}
}

func TestUserAgentKeepsCallerToken(t *testing.T) {
	// A User-Agent from -H/--header follows the CLI's own token.
	got := sends(t, "mycorp-wrapper/1.2")
	if want := customcfg.UserAgent() + " mycorp-wrapper/1.2"; got != want {
		t.Errorf("sent User-Agent = %q, want %q", got, want)
	}
}

func TestUserAgentWhenRequestHasNone(t *testing.T) {
	got := sends(t, "")
	if want := customcfg.UserAgent(); got != want {
		t.Errorf("sent User-Agent = %q, want %q", got, want)
	}
}
