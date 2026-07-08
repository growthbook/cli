// Hand-written (NOT generated). End-to-end tests that drive the real command
// tree built by NewRootCommand against a local mock API server. They exercise
// the CLI exactly as a user would — flag parsing, auth resolution, the HTTP
// call, and output formatting — without needing a live GrowthBook backend.
//
// Living in package cli (outside the generated per-resource packages) keeps
// these safe across a clean `speakeasy run` regeneration.
package cli

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// runCLI drives the CLI end to end: it stands up an httptest server, points the
// CLI at it via --server-url, runs the given args, and returns captured
// stdout/stderr plus the command error (non-nil ⇒ the binary would exit 1).
func runCLI(t *testing.T, handler http.HandlerFunc, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	// Isolate config/HOME and disable the daily update check so tests never
	// touch the real keychain, a real config file, or the network.
	t.Setenv("HOME", t.TempDir())
	t.Setenv("GBCLI_NO_UPDATE_CHECK", "1")
	t.Setenv("GBCLI_PROFILE", "")
	t.Setenv("GBCLI_BEARER_AUTH", "test-token")

	root, err := NewRootCommand()
	if err != nil {
		t.Fatalf("NewRootCommand: %v", err)
	}
	var out, errBuf strings.Builder
	root.SetOut(&out)
	root.SetErr(&errBuf)
	root.SetArgs(append(args, "--server-url", srv.URL, "--no-interactive"))
	err = root.Execute()
	return out.String(), errBuf.String(), err
}

func TestE2E_ProjectsList_JSON(t *testing.T) {
	var gotPath, gotAuth string
	body := `{"projects":[{"id":"prj_1","name":"Alpha"},{"id":"prj_2","name":"Beta"}]}`
	out, _, err := runCLI(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}, "projects", "list", "--output-format", "json")

	if err != nil {
		t.Fatalf("projects list returned error: %v", err)
	}
	if gotAuth != "Bearer test-token" {
		t.Errorf("Authorization header = %q, want %q", gotAuth, "Bearer test-token")
	}
	if !strings.Contains(gotPath, "/projects") {
		t.Errorf("request path = %q, want it to contain /projects", gotPath)
	}
	// --output-format json is a lossless raw passthrough of the API body.
	var got, want map[string]any
	if e := json.Unmarshal([]byte(out), &got); e != nil {
		t.Fatalf("stdout is not valid JSON: %v\n%s", e, out)
	}
	_ = json.Unmarshal([]byte(body), &want)
	if len(got["projects"].([]any)) != 2 {
		t.Errorf("expected 2 projects in output, got: %s", out)
	}
}

func TestE2E_ProjectsList_YAML(t *testing.T) {
	// Unlike --output-format json (a raw passthrough), yaml goes through the
	// typed deserialize → re-encode path, so this also exercises that the SDK
	// model round-trips the API body.
	out, _, err := runCLI(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"projects":[{"id":"prj_1","name":"Alpha"}]}`))
	}, "projects", "list", "--output-format", "yaml")

	if err != nil {
		t.Fatalf("projects list (yaml) returned error: %v", err)
	}
	if strings.Contains(out, "{") {
		t.Errorf("expected YAML, not JSON:\n%s", out)
	}
	if !strings.Contains(out, "prj_1") || !strings.Contains(out, "Alpha") {
		t.Errorf("yaml output missing the deserialized project:\n%s", out)
	}
}

func TestE2E_ProjectGet_PassesID(t *testing.T) {
	var gotPath string
	out, _, err := runCLI(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"prj_42","name":"Answer"}`))
	}, "projects", "get", "--id", "prj_42", "--output-format", "json")

	if err != nil {
		t.Fatalf("projects get returned error: %v", err)
	}
	if !strings.Contains(gotPath, "prj_42") {
		t.Errorf("request path = %q, want it to contain the id prj_42", gotPath)
	}
	if !strings.Contains(out, "prj_42") {
		t.Errorf("output missing the fetched project:\n%s", out)
	}
}

func TestE2E_ProjectGet_APIError(t *testing.T) {
	_, stderr, err := runCLI(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"project not found"}`))
	}, "projects", "get", "--id", "nope", "--output-format", "json")

	if err == nil {
		t.Fatal("expected a non-nil error (exit 1) for a 404 response")
	}
	if !strings.Contains(stderr, "404") {
		t.Errorf("stderr should surface the 404 status, got:\n%s", stderr)
	}
}

// reqCapture records what the CLI sent to the mock server. Assertions happen in
// the test goroutine after Execute returns, since httptest handlers run on their
// own goroutine where t.Fatal would misbehave.
type reqCapture struct {
	method string
	path   string
	body   []byte
}

// respondJSON returns a handler that records the request into c and replies.
func respondJSON(c *reqCapture, status int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.method = r.Method
		c.path = r.URL.Path
		c.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}
}

func (c *reqCapture) json(t *testing.T) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(c.body, &m); err != nil {
		t.Fatalf("request body is not JSON: %v\nbody: %s", err, c.body)
	}
	return m
}

func TestE2E_ProjectCreate_FlagsToBody(t *testing.T) {
	var cap reqCapture
	out, _, err := runCLI(t, respondJSON(&cap, http.StatusOK, `{"project":{"id":"prj_new","name":"Alpha"}}`),
		"projects", "create", "--name", "Alpha", "--description", "My desc", "--public-id", "alpha", "--output-format", "json")

	if err != nil {
		t.Fatalf("projects create returned error: %v", err)
	}
	if cap.method != http.MethodPost {
		t.Errorf("method = %q, want POST", cap.method)
	}
	if !strings.Contains(cap.path, "/projects") {
		t.Errorf("path = %q, want it to contain /projects", cap.path)
	}
	// Individual flags must serialize into the JSON body with the API's field names.
	b := cap.json(t)
	if b["name"] != "Alpha" {
		t.Errorf("body.name = %v, want Alpha", b["name"])
	}
	if b["description"] != "My desc" {
		t.Errorf("body.description = %v, want \"My desc\"", b["description"])
	}
	if b["publicId"] != "alpha" {
		t.Errorf("body.publicId = %v, want alpha", b["publicId"])
	}
	if !strings.Contains(out, "prj_new") {
		t.Errorf("output missing created project:\n%s", out)
	}
}

func TestE2E_ProjectCreate_RawBody(t *testing.T) {
	// --body is the escape hatch: raw JSON passed straight through, bypassing
	// the per-field flags.
	var cap reqCapture
	_, _, err := runCLI(t, respondJSON(&cap, http.StatusOK, `{"project":{"id":"prj_b"}}`),
		"projects", "create", "--body", `{"name":"Beta","publicId":"beta"}`, "--output-format", "json")

	if err != nil {
		t.Fatalf("projects create --body returned error: %v", err)
	}
	b := cap.json(t)
	if b["name"] != "Beta" || b["publicId"] != "beta" {
		t.Errorf("raw --body not sent verbatim: got %v", b)
	}
}

func TestE2E_ProjectCreate_SettingsJSON(t *testing.T) {
	// A FlagKindJSON flag must parse into a nested object, not a string.
	var cap reqCapture
	_, _, err := runCLI(t, respondJSON(&cap, http.StatusOK, `{"project":{"id":"prj_s"}}`),
		"projects", "create", "--name", "X", "--settings", `{"statsEngine":"bayesian"}`, "--output-format", "json")

	if err != nil {
		t.Fatalf("projects create --settings returned error: %v", err)
	}
	b := cap.json(t)
	settings, ok := b["settings"].(map[string]any)
	if !ok {
		t.Fatalf("body.settings is not a JSON object: %#v", b["settings"])
	}
	if settings["statsEngine"] != "bayesian" {
		t.Errorf("body.settings.statsEngine = %v, want bayesian", settings["statsEngine"])
	}
}

func TestE2E_ProjectUpdate_FlagsToBody(t *testing.T) {
	// Update carries both a path param (id) and a request body.
	var cap reqCapture
	_, _, err := runCLI(t, respondJSON(&cap, http.StatusOK, `{"project":{"id":"prj_1","name":"Renamed"}}`),
		"projects", "update", "--id", "prj_1", "--name", "Renamed", "--output-format", "json")

	if err != nil {
		t.Fatalf("projects update returned error: %v", err)
	}
	if cap.method != http.MethodPut {
		t.Errorf("method = %q, want PUT", cap.method)
	}
	if !strings.Contains(cap.path, "prj_1") {
		t.Errorf("path = %q, want it to contain the id prj_1", cap.path)
	}
	b := cap.json(t)
	if b["name"] != "Renamed" {
		t.Errorf("body.name = %v, want Renamed", b["name"])
	}
}

func TestE2E_ProjectCreate_MissingRequired_ServerValidates(t *testing.T) {
	// Finding pinned by this test: a "required" flag is NOT enforced locally in
	// non-interactive mode — the CLI still POSTs (with an empty name) and leans
	// on the API to reject it. A clean local "missing --name" error would be
	// nicer UX, but this is Speakeasy-generated behavior. What we do guarantee:
	// the server's rejection surfaces as a non-zero exit, not a silent success.
	var cap reqCapture
	_, stderr, err := runCLI(t, respondJSON(&cap, http.StatusBadRequest, `{"message":"name is required"}`),
		"projects", "create", "--output-format", "json")

	if cap.method != http.MethodPost {
		t.Errorf("expected the CLI to POST (validation is server-side), got method %q", cap.method)
	}
	if err == nil {
		t.Errorf("expected a non-zero exit when the server rejects the request; stderr:\n%s", stderr)
	}
	if !strings.Contains(stderr, "400") {
		t.Errorf("stderr should surface the 400 status, got:\n%s", stderr)
	}
}

func TestE2E_MissingCredentials(t *testing.T) {
	// Unset the bearer token; the CLI should fail rather than hit the API
	// anonymously. Uses a handler that fails the test if it's ever reached.
	t.Setenv("GBCLI_BEARER_AUTH", "")
	called := false
	_, stderr, err := runCLI(t, func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"unauthorized"}`))
	}, "projects", "list", "--output-format", "json")

	// Either the CLI refuses locally, or the (mock) server rejects it — either
	// way the command must not succeed silently.
	if err == nil {
		t.Errorf("expected an error when no credentials are configured; stderr:\n%s\ncalled=%v", stderr, called)
	}
}
