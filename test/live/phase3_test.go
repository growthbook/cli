//go:build live

// Phase 3 — datasource-dependent resources. These need a data source, which the
// API can't create (data-sources is read-only). scripts/e2e-setup.sh provisions
// the built-in demo datasource; if it's absent these tests skip and show up as
// uncovered in the coverage report rather than failing.
package live

import (
	"fmt"
	"testing"
)

// datasource holds the ids the Phase 3/4 lifecycles need.
type datasource struct {
	id             string
	identifierType string
	assignmentQ    string
}

// requireDatasource returns the first data source, skipping the test if none
// exists (so a run without the demo datasource degrades gracefully).
func requireDatasource(t *testing.T) datasource {
	t.Helper()
	list := runGB(t, "data-sources", "list", "-o", "json").ok(t)
	arr, ok := list.get("dataSources").([]any)
	if !ok || len(arr) == 0 {
		t.Skip("no data source available — run scripts/e2e-setup.sh to create the demo datasource")
	}
	m := arr[0].(map[string]any)
	ds := datasource{id: str(m["id"])}
	if its, ok := m["identifierTypes"].([]any); ok && len(its) > 0 {
		ds.identifierType = str(its[0].(map[string]any)["id"])
	}
	if aqs, ok := m["assignmentQueries"].([]any); ok && len(aqs) > 0 {
		ds.assignmentQ = str(aqs[0].(map[string]any)["id"])
	}
	return ds
}

func TestMetrics(t *testing.T) {
	ds := requireDatasource(t)
	body := fmt.Sprintf(`{"datasourceId":%q,"name":%q,"type":"binomial","sql":{"identifierTypes":[%q],"conversionSQL":"SELECT %s as %s, '2026-01-01' as timestamp"}}`,
		ds.id, uniqueName("metric"), ds.identifierType, ds.identifierType, ds.identifierType)
	crud{
		group:     "metrics",
		itemKey:   "metric",
		listKey:   "metrics",
		hasGet:    true,
		paginated: true,
		create:    []string{"--body", body},
		update:    []string{"--body", `{"name":"` + uniqueName("metric") + `-updated"}`},
	}.run(t)
}

func TestDimensions(t *testing.T) {
	ds := requireDatasource(t)
	body := fmt.Sprintf(`{"datasourceId":%q,"name":%q,"identifierType":%q,"query":"SELECT %s as %s, 'test' as value"}`,
		ds.id, uniqueName("dim"), ds.identifierType, ds.identifierType, ds.identifierType)
	crud{
		group:     "dimensions",
		itemKey:   "dimension",
		listKey:   "dimensions",
		hasGet:    true,
		paginated: true,
		create:    []string{"--body", body},
		update:    []string{"--body", `{"name":"` + uniqueName("dim") + `-updated"}`},
	}.run(t)
}

func TestSegments(t *testing.T) {
	ds := requireDatasource(t)
	body := fmt.Sprintf(`{"datasourceId":%q,"name":%q,"identifierType":%q,"type":"SQL","query":"SELECT %s as %s"}`,
		ds.id, uniqueName("seg"), ds.identifierType, ds.identifierType, ds.identifierType)
	crud{
		group:     "segments",
		itemKey:   "segment",
		listKey:   "segments",
		hasGet:    true,
		paginated: true,
		create:    []string{"--body", body},
		update:    []string{"--body", `{"name":"` + uniqueName("seg") + `-updated","owner":""}`},
	}.run(t)
}

func TestFactTables(t *testing.T) {
	ds := requireDatasource(t)
	body := fmt.Sprintf(`{"name":%q,"datasource":%q,"userIdTypes":[%q],"sql":"SELECT '%s' as %s, 10 as amount, '2026-01-01 00:00:00' as timestamp"}`,
		uniqueName("ft"), ds.id, ds.identifierType, ds.identifierType, ds.identifierType)
	crud{
		group:     "fact-tables",
		itemKey:   "factTable",
		listKey:   "factTables",
		hasGet:    true,
		paginated: true,
		create:    []string{"--body", body},
		update:    []string{"--body", `{"name":"` + uniqueName("ft") + `-updated"}`},
	}.run(t)
}
