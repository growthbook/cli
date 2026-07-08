//go:build live

// Phase 4 — complex resources. Experiments have no delete endpoint, so the
// lifecycle ends in archive-via-update instead of the crud helper's delete.
package live

import (
	"fmt"
	"testing"
)

func TestExperiments(t *testing.T) {
	ds := requireDatasource(t)
	body := fmt.Sprintf(`{"name":%q,"trackingKey":%q,"variations":[{"name":"Control","key":"0"},{"name":"Treatment","key":"1"}],"datasourceId":%q,"assignmentQueryId":%q}`,
		uniqueName("exp"), uniqueName("exp-tk"), ds.id, ds.assignmentQ)

	create := runGB(t, "experiments", "create", "--body", body, "-o", "json").ok(t)
	id := str(create.get("experiment", "id"))
	if id == "" {
		t.Fatalf("experiments create: no experiment.id in response:\n%s", create.Stdout)
	}

	if list := runGB(t, "experiments", "list", "--limit", "100", "-o", "json").ok(t); !list.listContainsID("experiments", id) {
		t.Errorf("experiments list does not include created id %s", id)
	}

	got := runGB(t, "experiments", "get", "--id", id, "-o", "json").ok(t)
	if str(got.get("experiment", "id")) != id {
		t.Errorf("experiments get returned wrong id: %s", got.Stdout)
	}

	runGB(t, "experiments", "update", "--id", id, "--body", `{"description":"updated via e2e"}`, "-o", "json").ok(t)

	// No delete endpoint — archive to clean up.
	archived := runGB(t, "experiments", "update", "--id", id, "--body", `{"archived":true}`, "-o", "json").ok(t)
	if archived.get("experiment", "archived") != true {
		t.Errorf("experiment was not archived: %s", archived.Stdout)
	}
}
