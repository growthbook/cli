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
	if updateCheckDisabled(cmd) {
		return
	}
	state := readCheckState()
	if time.Since(time.Unix(state.CheckedAt, 0)) < updateCheckInterval {
		return
	}
	// Record the attempt up-front (keeping the last-nudged version) so a hang
	// can't re-fire all day.
	state.CheckedAt = time.Now().Unix()
	writeCheckState(state)

	// One shared budget across both network calls so a slow endpoint can't make
	// startup drag (this runs on a human-interactive command, at most once/day).
	ctx, cancel := context.WithTimeout(cmd.Context(), 2*time.Second)
	defer cancel()

	latest, err := latestCLIRelease(ctx)
	if err != nil || latest == nil {
		return // can't tell whether a newer CLI exists → say nothing
	}
	server, _ := fetchServerBuild(ctx, resolveServerURL(cmd), resolveBearer(cmd))

	msg, ok := shouldNudge(Version, SpecDate, latest, server)
	if !ok || latest.Version == state.NudgedVersion {
		return // nothing to advise, or already nudged about this exact version
	}
	fmt.Fprintln(cmd.ErrOrStderr(), msg)
	state.NudgedVersion = latest.Version
	writeCheckState(state)
}

// shouldNudge decides whether to advise an upgrade, and returns the message.
// Pure (no I/O) so it is unit-testable.
//
//   - localVer / localSpecDate: this CLI's version and generated-against date.
//   - latest:                   the latest published CLI release.
//   - server:                   the connected backend's build (nil if unknown).
//
// Rule: only nudge if a newer CLI exists by at least a minor version (patch-only
// bumps are usually codegen-only and carry no new capability, so they stay
// quiet). Then, if we know both the latest CLI's generated-against date and the
// server's build date, require the server to be at least as new — otherwise
// upgrading would point the CLI at endpoints the server lacks. Missing dates →
// degrade to semver (permissive).
func shouldNudge(localVer, localSpecDate string, latest *cliRelease, server *serverBuild) (string, bool) {
	if latest == nil || !minorOrMajorAhead(localVer, latest.Version) {
		return "", false // already current, or only a patch ahead
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

// checkState persists between runs: when we last hit the network (rate-limits the
// check to once/day) and the last version we nudged about (per-version backoff —
// we advise an upgrade at most once per distinct newer version, so an ignored
// nudge doesn't nag daily). Stored as JSON at checkStampPath.
type checkState struct {
	CheckedAt     int64  `json:"checkedAt"`
	NudgedVersion string `json:"nudgedVersion,omitempty"`
}

func readCheckState() checkState {
	var s checkState
	p := checkStampPath()
	if p == "" {
		return s
	}
	if data, err := os.ReadFile(p); err == nil {
		_ = json.Unmarshal(data, &s)
	}
	return s
}

func writeCheckState(s checkState) {
	p := checkStampPath()
	if p == "" {
		return
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return
	}
	if data, err := json.Marshal(s); err == nil {
		_ = os.WriteFile(p, data, 0o600)
	}
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

// latestCLIRelease fetches the latest published CLI release from the GitHub API:
// the version tag, plus the generated-against date the release pipeline embeds in
// the release body (see .goreleaser.yaml's `gb-spec-date:` footer marker).
//
// GitHub's /releases/latest excludes prereleases, so while only `@next`
// prereleases exist this returns (nil, nil) and the advisory cleanly no-ops —
// stable users are never nudged toward a prerelease.
func latestCLIRelease(ctx context.Context) (*cliRelease, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.github.com/repos/growthbook/cli/releases/latest", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, nil // no stable release yet (e.g. 404 when only prereleases exist)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<16))
	if err != nil {
		return nil, err
	}
	var rel struct {
		TagName string `json:"tag_name"`
		Body    string `json:"body"`
	}
	if err := json.Unmarshal(body, &rel); err != nil {
		return nil, err
	}
	if rel.TagName == "" {
		return nil, nil
	}
	return &cliRelease{
		Version:  strings.TrimPrefix(rel.TagName, "v"),
		SpecDate: parseSpecDate(rel.Body),
	}, nil
}

// parseSpecDate extracts the generated-against date the release pipeline embeds
// in the release body as `gb-spec-date: YYYY-MM-DD`. Returns "" if absent (the
// compat gate then degrades to permissive semver).
func parseSpecDate(body string) string {
	const marker = "gb-spec-date:"
	i := strings.Index(body, marker)
	if i < 0 {
		return ""
	}
	rest := strings.TrimSpace(body[i+len(marker):])
	end := strings.IndexFunc(rest, func(r rune) bool { return r != '-' && (r < '0' || r > '9') })
	if end >= 0 {
		rest = rest[:end]
	}
	return rest
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

// minorOrMajorAhead reports whether b is ahead of a by at least a minor version.
// Patch-only bumps (often codegen-only, no new capability) don't warrant a nudge.
func minorOrMajorAhead(a, b string) bool {
	pa, pb := parseSemver(a), parseSemver(b)
	if pa[0] != pb[0] {
		return pb[0] > pa[0]
	}
	return pb[1] > pa[1]
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
