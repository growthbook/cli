package client

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/growthbook/cli/v2/internal/customcfg"
)

func TestUserAgentFormat(t *testing.T) {
	ua := UserAgent()
	want := regexp.MustCompile(`^growthbook-cli/\S+ \(go\S+; \w+/\w+\)$`)
	if !want.MatchString(ua) {
		t.Errorf("User-Agent %q does not match %v", ua, want)
	}
}

func TestVersionPrefersLdflagsOverBuildInfo(t *testing.T) {
	original := customcfg.Version
	t.Cleanup(func() { customcfg.Version = original })

	customcfg.Version = "2.6.0"
	if got := version(); got != "2.6.0" {
		t.Errorf("version() = %q, want the ldflags value 2.6.0", got)
	}

	// Unstamped build: falls through to build info, and under `go test` that
	// has no module version either, so it lands on the "dev" sentinel.
	customcfg.Version = "dev"
	if got := version(); got == "" {
		t.Error("version() returned empty for an unstamped build")
	}
}

func TestWrapClientForUserAgentOverwritesSDKDefault(t *testing.T) {
	var got string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.UserAgent()
	}))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	// What the generated SDK sets before handing the request to the client.
	req.Header.Set("User-Agent", "speakeasy-sdk/go 0.0.1")

	res, err := WrapClientForUserAgent(&http.Client{}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if got != UserAgent() {
		t.Errorf("sent User-Agent = %q, want %q", got, UserAgent())
	}
}
