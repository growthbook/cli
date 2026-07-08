//go:build live

// Live CLI e2e harness (NOT built or run by the default offline `go test ./...`
// — the `live` build tag gates the whole package). It drives the compiled
// ./growthbook binary as a subprocess against a real GrowthBook server, so it
// validates the actual shipped artifact: flag parsing, body serialization, auth,
// output formatting, and exit codes. Stand up the server + creds with
// scripts/e2e-setup.sh first; see test/live/README.md.
package live

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync/atomic"
	"testing"
)

// Set by TestMain in live_test.go from the environment.
var (
	gbBinary    string // path to the compiled growthbook binary
	gbServerURL string // GB_SERVER_URL, e.g. http://localhost:3100/api
	gbBearer    string // GBCLI_BEARER_AUTH secret key
	cov         = newCoverage()
)

var nameCounter atomic.Int64

// uniqueName returns a collision-free, run-scoped name so reruns and parallel
// resources never clash (and are easy to spot/clean up on the server).
func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-e2e-%d-%d", prefix, os.Getpid(), nameCounter.Add(1))
}

// gbResult captures everything a live command produced.
type gbResult struct {
	Args     []string
	Stdout   string
	Stderr   string
	ExitCode int
	JSON     any // parsed stdout when it's valid JSON, else nil
}

// runGB drives the CLI end to end: it appends --server-url/--no-interactive,
// injects the bearer token via env (as scripts/gb does), runs the binary, and
// returns stdout/stderr, the exit code, and parsed JSON. It also records the
// command path for the coverage report.
func runGB(t *testing.T, args ...string) gbResult {
	return runGBAs(t, gbBearer, args...)
}

// runGBAs is runGB with an explicit bearer token, for auth error-path tests.
func runGBAs(t *testing.T, bearer string, args ...string) gbResult {
	t.Helper()
	cov.record(args)

	full := append([]string{}, args...)
	full = append(full, "--server-url", gbServerURL, "--no-interactive")
	cmd := exec.Command(gbBinary, full...)
	cmd.Env = append(os.Environ(),
		"GBCLI_BEARER_AUTH="+bearer,
		"GBCLI_NO_UPDATE_CHECK=1",
		"GBCLI_PROFILE=",
	)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	err := cmd.Run()

	res := gbResult{Args: args, Stdout: out.String(), Stderr: errb.String()}
	if err != nil {
		var ee *exec.ExitError
		if !isExitError(err, &ee) {
			t.Fatalf("could not execute %s %v: %v", gbBinary, args, err)
		}
		res.ExitCode = ee.ExitCode()
	}
	_ = json.Unmarshal(out.Bytes(), &res.JSON)
	return res
}

func isExitError(err error, target **exec.ExitError) bool {
	if ee, ok := err.(*exec.ExitError); ok {
		*target = ee
		return true
	}
	return false
}

// ok asserts the command exited 0, failing with the captured stderr otherwise.
// It returns the result so calls can chain.
func (r gbResult) ok(t *testing.T) gbResult {
	t.Helper()
	if r.ExitCode != 0 {
		t.Fatalf("`growthbook %s` exited %d\nstderr: %s\nstdout: %s",
			strings.Join(r.Args, " "), r.ExitCode, r.Stderr, r.Stdout)
	}
	return r
}

// failed asserts the command exited non-zero (used for error-path tests).
func (r gbResult) failed(t *testing.T) gbResult {
	t.Helper()
	if r.ExitCode == 0 {
		t.Fatalf("`growthbook %s` unexpectedly succeeded\nstdout: %s",
			strings.Join(r.Args, " "), r.Stdout)
	}
	return r
}

// premiumGated reports whether the command failed because the resource needs a
// premium/enterprise license — which a free self-hosted test instance lacks. The
// suite skips these rather than failing (they surface as uncovered).
func (r gbResult) premiumGated() bool {
	if r.ExitCode == 0 {
		return false
	}
	blob := r.Stdout + r.Stderr
	return strings.Contains(blob, "premium plan") || strings.Contains(blob, "requires an enterprise")
}

// get navigates the parsed JSON body by successive object keys, returning nil if
// any key is missing or the shape doesn't match.
func (r gbResult) get(keys ...string) any {
	cur := r.JSON
	for _, k := range keys {
		m, ok := cur.(map[string]any)
		if !ok {
			return nil
		}
		cur = m[k]
	}
	return cur
}

// str is a nil-safe string coercion for JSON values.
func str(v any) string {
	s, _ := v.(string)
	return s
}

// listContainsID reports whether the array at r.get(arrayKey) has an element
// whose "id" equals id. Used for "list includes the thing we just created".
func (r gbResult) listContainsID(arrayKey, id string) bool {
	arr, ok := r.get(arrayKey).([]any)
	if !ok {
		return false
	}
	for _, el := range arr {
		if m, ok := el.(map[string]any); ok && str(m["id"]) == id {
			return true
		}
	}
	return false
}

// crud describes the standard create→list→get→update→delete lifecycle for a
// resource whose create/get/update responses wrap the object under itemKey and
// whose list wraps the array under listKey (they often differ, e.g. SDK
// connections use "sdkConnection" but list under "connections").
type crud struct {
	group     string
	itemKey   string
	listKey   string
	create    []string // flags after "<group> create"
	update    []string // flags after "<group> update --id <id>" (omit to skip update)
	hasGet    bool     // resource exposes a get-by-id command
	paginated bool     // list supports --limit (pass a high limit so the new item is on the page)

	// Some resources refuse a REST delete until the object is archived. Set one
	// of these to archive first (and exercise the archive path as a bonus).
	archiveVerb      string // e.g. "archive" — runs "<group> archive --id <id>"
	archiveViaUpdate bool   // runs "<group> update --id <id> --body {archived:true}"
}

// run drives the full lifecycle, failing t on any deviation. It exercises the
// CLI's flag→body serialization, path-param handling, and output parsing against
// the live server, and records each command for the coverage report.
func (c crud) run(t *testing.T) {
	t.Helper()

	createArgs := append([]string{c.group, "create"}, c.create...)
	createArgs = append(createArgs, "-o", "json")
	res := runGB(t, createArgs...)
	if res.premiumGated() {
		t.Skipf("%s requires a premium plan — not available on a free test instance", c.group)
	}
	res.ok(t)
	id := str(res.get(c.itemKey, "id"))
	if id == "" {
		t.Fatalf("%s create: no %s.id in response:\n%s", c.group, c.itemKey, res.Stdout)
	}

	listArgs := []string{c.group, "list", "-o", "json"}
	if c.paginated {
		listArgs = []string{c.group, "list", "--limit", "100", "-o", "json"}
	}
	if list := runGB(t, listArgs...).ok(t); !list.listContainsID(c.listKey, id) {
		t.Errorf("%s list does not include created id %s", c.group, id)
	}

	if c.hasGet {
		got := runGB(t, c.group, "get", "--id", id, "-o", "json").ok(t)
		if str(got.get(c.itemKey, "id")) != id {
			t.Errorf("%s get --id %s returned wrong id: %s", c.group, id, got.Stdout)
		}
	}

	if len(c.update) > 0 {
		updateArgs := append([]string{c.group, "update", "--id", id}, c.update...)
		updateArgs = append(updateArgs, "-o", "json")
		runGB(t, updateArgs...).ok(t)
	}

	if c.archiveVerb != "" {
		runGB(t, c.group, c.archiveVerb, "--id", id, "-o", "json").ok(t)
	}
	if c.archiveViaUpdate {
		runGB(t, c.group, "update", "--id", id, "--body", `{"archived":true}`, "-o", "json").ok(t)
	}

	runGB(t, c.group, "delete", "--id", id, "-o", "json").ok(t)
}
