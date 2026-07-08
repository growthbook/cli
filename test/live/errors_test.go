//go:build live

// Error-path behavior: the CLI must surface server failures as a non-zero exit
// with the status in stderr, not swallow them into a silent success. Mirrors the
// offline TestE2E_ProjectGet_APIError / TestE2E_MissingCredentials, against the
// live server.
package live

import (
	"strings"
	"testing"
)

func TestErrorNotFound(t *testing.T) {
	// GrowthBook returns 400 ("Could not find ... with that id") for a missing
	// resource; what matters is the CLI surfaces the failure rather than exiting 0.
	res := runGB(t, "projects", "get", "--id", "prj_does_not_exist_e2e", "-o", "json").failed(t)
	if !strings.Contains(res.Stderr, "400") && !strings.Contains(res.Stderr, "404") {
		t.Errorf("expected a 4xx status in stderr, got:\n%s", res.Stderr)
	}
}

func TestErrorBadAuth(t *testing.T) {
	res := runGBAs(t, "secret_invalid_token", "projects", "list", "-o", "json").failed(t)
	if !strings.Contains(res.Stderr, "401") && !strings.Contains(res.Stderr, "403") {
		t.Errorf("expected 401/403 in stderr for an invalid token, got:\n%s", res.Stderr)
	}
}
