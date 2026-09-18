package customcfg

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestUserAgentFormat(t *testing.T) {
	ua := UserAgent()
	want := regexp.MustCompile(`^growthbook-cli/\S+ \(.+; \w+/\w+\)$`)
	if !want.MatchString(ua) {
		t.Errorf("User-Agent %q does not match %v", ua, want)
	}
}

// The compat check bypasses the SDK client, so it has to set the header itself.
func TestFetchServerBuildSendsUserAgent(t *testing.T) {
	var got string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.UserAgent()
		_, _ = w.Write([]byte(`{"version":"3.0.0"}`))
	}))
	defer server.Close()

	if _, err := fetchServerBuild(context.Background(), server.URL, ""); err != nil {
		t.Fatal(err)
	}
	if want := UserAgent(); got != want {
		t.Errorf("sent User-Agent = %q, want %q", got, want)
	}
}

func TestVersionPrefersLdflagsOverBuildInfo(t *testing.T) {
	original := Version
	t.Cleanup(func() { Version = original })

	Version = "2.6.0"
	if got := version(); got != "2.6.0" {
		t.Errorf("version() = %q, want the ldflags value 2.6.0", got)
	}

	// Unstamped build: falls through to build info, and under `go test` that
	// has no module version either, so it lands on the "dev" sentinel.
	Version = "dev"
	if got := version(); got == "" {
		t.Error("version() returned empty for an unstamped build")
	}
}
