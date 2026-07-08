//go:build live

package live

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMain wires up the live suite: it loads credentials (from the env or the
// repo-root .env.local written by scripts/e2e-setup.sh), locates the compiled
// binary, loads the command-surface golden for coverage, runs the tests, then
// writes the coverage report. It hard-fails with setup instructions if the
// server/creds/binary aren't available, so a misconfigured run is obvious rather
// than silently green.
func TestMain(m *testing.M) {
	// repoRoot = two levels up from test/live.
	repoRoot, err := filepath.Abs("../..")
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot resolve repo root:", err)
		os.Exit(1)
	}
	loadDotEnv(filepath.Join(repoRoot, ".env.local"))

	gbServerURL = os.Getenv("GB_SERVER_URL")
	gbBearer = os.Getenv("GBCLI_BEARER_AUTH")

	gbBinary = os.Getenv("GB_CLI_BINARY")
	if gbBinary == "" {
		gbBinary = filepath.Join(repoRoot, "growthbook")
	}
	if abs, err := filepath.Abs(gbBinary); err == nil {
		gbBinary = abs
	}

	if gbServerURL == "" || gbBearer == "" {
		fmt.Fprintln(os.Stderr, setupHint("GB_SERVER_URL and GBCLI_BEARER_AUTH must be set"))
		os.Exit(1)
	}
	if _, err := os.Stat(gbBinary); err != nil {
		fmt.Fprintln(os.Stderr, setupHint(fmt.Sprintf("CLI binary not found at %s — run: go build -o growthbook ./cmd/growthbook", gbBinary)))
		os.Exit(1)
	}
	if err := cov.loadSurface(filepath.Join(repoRoot, "internal/cli/testdata/command-surface.txt")); err != nil {
		fmt.Fprintln(os.Stderr, "cannot load command-surface golden:", err)
		os.Exit(1)
	}

	code := m.Run()

	// Emit the coverage report to stdout and to test/live/coverage.txt.
	covered, total := cov.report(os.Stdout)
	if f, err := os.Create("coverage.txt"); err == nil {
		cov.report(f)
		f.Close()
		fmt.Printf("coverage report written to test/live/coverage.txt (%d/%d)\n", covered, total)
	}
	os.Exit(code)
}

func setupHint(msg string) string {
	return "live suite not runnable: " + msg + "\n" +
		"Stand up a server and bootstrap credentials first:\n" +
		"  ./scripts/e2e-setup.sh\n" +
		"  go build -o growthbook ./cmd/growthbook\n" +
		"  go test -tags live ./test/live/ -v"
}

// loadDotEnv sets KEY=VALUE pairs from a .env file into the process env, without
// overriding anything already set. Best-effort: a missing file is fine.
func loadDotEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if _, exists := os.LookupEnv(k); !exists {
			os.Setenv(k, v)
		}
	}
}
