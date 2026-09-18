package client

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestUserAgentFormat(t *testing.T) {
	ua := UserAgent()
	want := regexp.MustCompile(`^growthbook-cli/\S+ \(go\S+; \w+/\w+\)$`)
	if !want.MatchString(ua) {
		t.Errorf("User-Agent %q does not match %v", ua, want)
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
