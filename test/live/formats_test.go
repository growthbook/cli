//go:build live

// Output-format matrix. Formatting is CLI-specific behavior the raw API can't
// validate, so run one representative read through every format and confirm each
// renders without error and carries the data. json additionally must parse.
package live

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestOutputFormats(t *testing.T) {
	// A known project id to look for in every rendering.
	list := runGB(t, "projects", "list", "-o", "json").ok(t)
	arr, ok := list.get("projects").([]any)
	if !ok || len(arr) == 0 {
		t.Fatal("expected at least one project (the default) to format")
	}
	id := str(arr[0].(map[string]any)["id"])

	// table is deliberately disabled (#19 upstream won't-fix; see the
	// unplanned-table-output patch) — it must fail fast with a clear error.
	t.Run("table", func(t *testing.T) {
		res := runGB(t, "projects", "list", "-o", "table")
		if res.ExitCode == 0 {
			t.Fatalf("-o table should be rejected while disabled, got success:\n%s", res.Stdout)
		}
		if !strings.Contains(res.Stderr, "not currently supported") {
			t.Errorf("expected the deliberate-disable error, got:\n%s", res.Stderr)
		}
	})

	for _, format := range []string{"pretty", "json", "yaml", "toon"} {
		format := format
		t.Run(format, func(t *testing.T) {
			res := runGB(t, "projects", "list", "-o", format).ok(t)
			if strings.TrimSpace(res.Stdout) == "" {
				t.Fatalf("%s output is empty", format)
			}
			if format == "json" {
				var v any
				if err := json.Unmarshal([]byte(res.Stdout), &v); err != nil {
					t.Fatalf("json output does not parse: %v\n%s", err, res.Stdout)
				}
			}
			// toon abbreviates ids; the other formats carry it verbatim.
			if format != "toon" && !strings.Contains(res.Stdout, id) {
				t.Errorf("%s output missing project id %s:\n%s", format, id, res.Stdout)
			}
		})
	}
}
