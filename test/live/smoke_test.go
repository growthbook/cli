//go:build live

// Read-only smoke coverage: run `list` (and a couple of top-level reads) for
// every resource group whose list needs no parent id, asserting a clean exit and
// valid JSON. This broadens command coverage cheaply without mutating state.
// Groups excluded here are gated (organizations = multi-org only, visual-
// changesets = not in this API version, archetypes = premium) and show up as
// uncovered in the report.
package live

import "testing"

var smokeListGroups = []string{
	"attributes", "constant-revisions", "constants", "contextual-bandit-queries",
	"contextual-bandits", "custom-fields", "dashboards", "data-sources", "dimensions",
	"environments", "experiments", "experiment-templates", "fact-metrics", "fact-tables",
	"features", "members", "metric-groups", "metrics", "namespaces", "ramp-schedules",
	"ramp-schedule-templates", "reports", "saved-group-revisions", "saved-groups",
	"SDK-connections", "segments", "teams",
}

func TestSmokeLists(t *testing.T) {
	for _, g := range smokeListGroups {
		g := g
		t.Run(g, func(t *testing.T) {
			res := runGB(t, g, "list", "-o", "json")
			if res.premiumGated() {
				t.Skipf("%s list requires a premium plan", g)
			}
			res.ok(t)
			if res.JSON == nil {
				t.Errorf("%s list did not return valid JSON:\n%s", g, res.Stdout)
			}
		})
	}
}

func TestSmokeTopLevelReads(t *testing.T) {
	t.Run("whoami", func(t *testing.T) {
		runGB(t, "whoami", "-o", "json").ok(t)
	})
	t.Run("meta get-version", func(t *testing.T) {
		runGB(t, "meta", "get-version", "-o", "json").ok(t)
	})
}
