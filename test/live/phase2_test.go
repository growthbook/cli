//go:build live

// Phase 2 — resources with light dependencies.
package live

import "testing"

func TestSavedGroups(t *testing.T) {
	name := uniqueName("sg")
	crud{
		group:       "saved-groups",
		itemKey:     "savedGroup",
		listKey:     "savedGroups",
		hasGet:      true,
		paginated:   true,
		create:      []string{"--name", name, "--type", "list", "--attribute-key", "id", "--values", `["a","b","c"]`},
		update:      []string{"--name", name + "-updated"},
		archiveVerb: "archive", // REST delete is refused until the group is archived
	}.run(t)
}

// SDK connections require an existing environment, so create a throwaway one and
// clean it up. The list envelope is "connections" while create/get use
// "sdkConnection".
func TestSDKConnections(t *testing.T) {
	env := uniqueName("sdkenv")
	runGB(t, "environments", "create", "--id", env, "--description", "e2e sdk env", "-o", "json").ok(t)
	t.Cleanup(func() { runGB(t, "environments", "delete", "--id", env, "-o", "json") })

	crud{
		group:     "SDK-connections",
		itemKey:   "sdkConnection",
		listKey:   "connections",
		hasGet:    true,
		paginated: true,
		create:    []string{"--name", uniqueName("sdk"), "--language", "javascript", "--environment", env},
		update:    []string{"--name", uniqueName("sdk") + "-updated"},
	}.run(t)
}

func TestFeatures(t *testing.T) {
	crud{
		group:            "features",
		itemKey:          "feature",
		listKey:          "features",
		hasGet:           true,
		paginated:        true,
		create:           []string{"--id", uniqueName("feat"), "--value-type", "boolean", "--default-value", "false", "--owner", "e2e"},
		update:           []string{"--description", "updated via e2e"},
		archiveViaUpdate: true, // a live feature can't be REST-deleted until archived
	}.run(t)
}
