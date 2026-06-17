// Update / compatibility advisory (hand-written, NOT generated). Lives in the
// custom layer so it survives regeneration.
//
// Principle (see the migration plan §F): the nudge must be *correct* — never push
// a user toward a CLI that targets endpoints their backend lacks. The orderable
// compatibility signal is the build DATE: nudge to a newer CLI only if the
// connected backend's build date is >= that CLI's generated-against date. semver
// is compared first (is a newer CLI even out?), then date. If either date is
// empty (dev builds), we degrade to semver and lean permissive.
//
// Guardrails: advisory-only (never auto-updates), at most once/day, stderr only
// (so JSON on stdout stays clean), auto-off in CI / non-TTY, and opt-out via
// --no-update-check or GBCLI_NO_UPDATE_CHECK.
package customcfg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/growthbook/cli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// Build metadata, injected via -ldflags at release (see .goreleaser.yaml). Empty
// in local/dev builds.
//   - Version:  the CLI's own semver, e.g. "1.0.0"
//   - Commit:   short git SHA of the CLI build (human-readable messages only)
//   - SpecDate: the generated-against date — the commit date (YYYY-MM-DD) of the
//     backend spec.yaml the CLI was generated from. The orderable compat signal.
var (
	Version  = "dev"
	Commit   = ""
	SpecDate = ""
)

const updateCheckInterval = 24 * time.Hour

// serverBuild is the backend's GET /v1/version payload (tag `meta`).
type serverBuild struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

// cliRelease describes the latest published CLI release. Sourced from GitHub
// release metadata; see latestCLIRelease (not yet wired — release infra).
type cliRelease struct {
	Version  string
	SpecDate string
}

// MaybeCheckForUpdate is the guarded entry point, called from OnStartup. It never
// blocks meaningfully (bounded network timeout, at most once/day) and never
// returns an error — an advisory must not break the user's command.
func MaybeCheckForUpdate(cmd *cobra.Command) {
	if updateCheckDisabled(cmd) || recentlyChecked() {
		return
	}
	touchCheckStamp() // record the attempt up-front so a hang can't re-fire all day

	latest, err := latestCLIRelease()
	if err != nil || latest == nil {
		return // can't tell whether a newer CLI exists → say nothing
	}

	ctx, cancel := context.WithTimeout(cmd.Context(), 1500*time.Millisecond)
	defer cancel()
	server, _ := fetchServerBuild(ctx, resolveServerURL(cmd), resolveBearer(cmd))

	if msg, ok := shouldNudge(Version, SpecDate, latest, server); ok {
		fmt.Fprintln(cmd.ErrOrStderr(), msg)
	}
}

// shouldNudge decides whether to advise an upgrade, and returns the message.
// Pure (no I/O) so it is unit-testable.
//
//   - localVer / localSpecDate: this CLI's version and generated-against date.
//   - latest:                   the latest published CLI release.
//   - server:                   the connected backend's build (nil if unknown).
//
// Rule: only nudge if a newer CLI exists (semver). Then, if we know both the
// latest CLI's generated-against date and the server's build date, require the
// server to be at least as new — otherwise upgrading would point the CLI at
// endpoints the server lacks. Missing dates → degrade to semver (permissive).
func shouldNudge(localVer, localSpecDate string, latest *cliRelease, server *serverBuild) (string, bool) {
	if latest == nil || compareSemver(latest.Version, localVer) <= 0 {
		return "", false // already on the latest (or newer) CLI
	}
	if latest.SpecDate != "" && server != nil && server.Date != "" {
		if server.Date < latest.SpecDate {
			// Backend predates what the newer CLI targets — do NOT push the
			// upgrade; it would 404 on the new endpoints.
			return "", false
		}
	}
	msg := fmt.Sprintf("growthbook: a newer CLI is available (%s → %s). Upgrade with `npm update -g growthbook` (or see the README).", localVer, latest.Version)
	return msg, true
}

// Versioned404Hint returns an actionable message for a 404 on a versioned
// endpoint — defense-in-depth pointing at the legacy group / a server upgrade.
// (Wiring it into the SDK error path is release infra; exposed for that use.)
func Versioned404Hint(requestPath string) string {
	if strings.Contains(requestPath, "/v2/features") {
		return "Hint: this server may not support the v2 feature endpoints. Try the `features-v1` commands, or upgrade your GrowthBook server."
	}
	return "Hint: this endpoint may not exist on your GrowthBook server version — check the server is up to date."
}

// --- guardrails -------------------------------------------------------------

func updateCheckDisabled(cmd *cobra.Command) bool {
	if v := os.Getenv("GBCLI_NO_UPDATE_CHECK"); v != "" && v != "0" && v != "false" {
		return true
	}
	if flagChanged(cmd, "no-update-check") {
		if b, err := cmd.Flags().GetBool("no-update-check"); err == nil && b {
			return true
		}
	}
	if isCI() {
		return true
	}
	// Advisory is for humans: only when stderr is a terminal.
	return !term.IsTerminal(int(os.Stderr.Fd()))
}

func isCI() bool {
	for _, k := range []string{"CI", "CONTINUOUS_INTEGRATION", "BUILD_NUMBER", "GITHUB_ACTIONS"} {
		if v := os.Getenv(k); v != "" && v != "0" && v != "false" {
			return true
		}
	}
	return false
}

func checkStampPath() string {
	dir := filepath.Dir(config.GetConfigPath())
	if dir == "" || dir == "." {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		dir = filepath.Join(home, ".config", "growthbook")
	}
	return filepath.Join(dir, "update-check")
}

// recentlyChecked reports whether we checked within updateCheckInterval.
func recentlyChecked() bool {
	p := checkStampPath()
	if p == "" {
		return false
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return false
	}
	unix, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		return false
	}
	return time.Since(time.Unix(unix, 0)) < updateCheckInterval
}

func touchCheckStamp() {
	p := checkStampPath()
	if p == "" {
		return
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return
	}
	_ = os.WriteFile(p, []byte(strconv.FormatInt(time.Now().Unix(), 10)), 0o600)
}

// --- backend / release lookups ----------------------------------------------

func fetchServerBuild(ctx context.Context, serverURL, bearer string) (*serverBuild, error) {
	if serverURL == "" {
		return nil, fmt.Errorf("no server URL")
	}
	url := strings.TrimRight(serverURL, "/") + "/v1/version"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+strings.TrimPrefix(bearer, "Bearer "))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<16))
	if err != nil {
		return nil, err
	}
	var b serverBuild
	if err := json.Unmarshal(body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

// latestCLIRelease fetches the latest published CLI release (version +
// generated-against date). DEFERRED: needs the release pipeline to publish the
// generated-against date in GitHub release metadata (no releases exist yet).
// Returns (nil, nil) so the advisory cleanly no-ops until that lands.
func latestCLIRelease() (*cliRelease, error) {
	// TODO(release-infra): GET the latest release from the GitHub API and read
	// its version tag + the generated-against date published in the release
	// metadata, then return it here.
	return nil, nil
}

// resolveServerURL / resolveBearer read the effective connection settings the
// same way the generated client does, after applyProfile has run.
func resolveServerURL(cmd *cobra.Command) string {
	for _, name := range []string{"server-url", "server", "domain"} {
		if fl := cmd.Flags().Lookup(name); fl != nil {
			if v := fl.Value.String(); v != "" {
				return v
			}
		}
	}
	return ""
}

func resolveBearer(cmd *cobra.Command) string {
	if fl := cmd.Flags().Lookup("bearer-auth"); fl != nil && fl.Changed {
		if v := fl.Value.String(); v != "" {
			return v
		}
	}
	return os.Getenv("GBCLI_BEARER_AUTH")
}

// compareSemver returns -1/0/1 for a<b / a==b / a>b. Tolerant of a leading "v"
// and of non-numeric/missing segments (treated as 0); pre-release suffixes on
// the patch segment are ignored.
func compareSemver(a, b string) int {
	pa, pb := parseSemver(a), parseSemver(b)
	for i := 0; i < 3; i++ {
		if pa[i] != pb[i] {
			if pa[i] < pb[i] {
				return -1
			}
			return 1
		}
	}
	return 0
}

func parseSemver(v string) [3]int {
	v = strings.TrimPrefix(strings.TrimSpace(v), "v")
	var out [3]int
	for i, seg := range strings.SplitN(v, ".", 3) {
		if i > 2 {
			break
		}
		// keep the leading digits only, dropping any pre-release/build suffix
		// ("0-rc1" → "0", "0+meta" → "0", "rc" → "" → 0).
		num := seg
		for j, r := range seg {
			if r < '0' || r > '9' {
				num = seg[:j]
				break
			}
		}
		out[i], _ = strconv.Atoi(num) // "" → 0 via the error path
	}
	return out
}
