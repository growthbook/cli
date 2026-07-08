//go:build live

// Phase 1 — standalone resources with no dependencies. Mirrors the phase
// ordering of the back-end's test-rest-api.ts so lifecycles run in an order the
// server accepts.
package live

import "testing"

func TestProjects(t *testing.T) {
	name := uniqueName("proj")
	crud{
		group:     "projects",
		itemKey:   "project",
		listKey:   "projects",
		hasGet:    true,
		paginated: true,
		create:    []string{"--name", name, "--description", "e2e", "--public-id", uniqueName("pid")},
		update:    []string{"--name", name + "-updated"},
	}.run(t)
}

func TestEnvironments(t *testing.T) {
	crud{
		group:   "environments",
		itemKey: "environment",
		listKey: "environments",
		create:  []string{"--id", uniqueName("env"), "--description", "e2e"},
		update:  []string{"--description", "e2e updated"},
	}.run(t)
}

// Attributes are keyed by --property (not --id) and have no get command, so they
// don't fit the crud helper.
func TestAttributes(t *testing.T) {
	prop := uniqueName("attr")
	create := runGB(t, "attributes", "create", "--property", prop, "--datatype", "string", "--description", "e2e", "-o", "json").ok(t)
	if str(create.get("attribute", "property")) != prop {
		t.Fatalf("attributes create: unexpected response:\n%s", create.Stdout)
	}

	if list := runGB(t, "attributes", "list", "-o", "json").ok(t); !attributesContain(list, prop) {
		t.Errorf("attributes list does not include created property %s", prop)
	}

	runGB(t, "attributes", "update", "--property", prop, "--datatype", "string", "--description", "e2e updated", "-o", "json").ok(t)
	runGB(t, "attributes", "delete", "--property", prop, "-o", "json").ok(t)
}

func attributesContain(r gbResult, prop string) bool {
	arr, ok := r.get("attributes").([]any)
	if !ok {
		return false
	}
	for _, el := range arr {
		if m, ok := el.(map[string]any); ok && str(m["property"]) == prop {
			return true
		}
	}
	return false
}

func TestArchetypes(t *testing.T) {
	name := uniqueName("arch")
	crud{
		group:   "archetypes",
		itemKey: "archetype",
		listKey: "archetypes",
		hasGet:  true,
		create:  []string{"--name", name, "--is-public=true", "--attributes", `{"id":"123"}`},
		update:  []string{"--attributes", `{"id":"456"}`},
	}.run(t)
}
