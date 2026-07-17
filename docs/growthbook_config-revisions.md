## growthbook config-revisions

Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)

### Synopsis

Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert). Publishing a schema change cascades the "base wins" normalization to descendant configs; a publish that removes or retypes fields descendants still use soft-blocks with a 422 unless `?ignoreWarnings=true`. Pass `version: "new"` on edit endpoints to auto-create a draft.

```
growthbook config-revisions [flags]
```

### Options

```
  -h, --help   help for config-revisions
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
      --dry-run                       Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray            Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers               Include HTTP response headers in the output
  -q, --jq string                     Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive                Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-update-check               Disable the once-a-day check for a newer CLI version
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, toon. (default "pretty")
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook config-revisions archive](growthbook_config-revisions_archive.md)	 - Stage an archive/unarchive in a draft revision
* [growthbook config-revisions create](growthbook_config-revisions_create.md)	 - Create a draft revision
* [growthbook config-revisions delete-projection](growthbook_config-revisions_delete-projection.md)	 - Remove a config's per-source render projection on a draft
* [growthbook config-revisions discard](growthbook_config-revisions_discard.md)	 - Discard a draft revision
* [growthbook config-revisions get](growthbook_config-revisions_get.md)	 - Get a single config revision
* [growthbook config-revisions latest](growthbook_config-revisions_latest.md)	 - Get the most recent active draft revision
* [growthbook config-revisions list](growthbook_config-revisions_list.md)	 - List config revisions across the organization
* [growthbook config-revisions list-for-config](growthbook_config-revisions_list-for-config.md)	 - List revisions for a config
* [growthbook config-revisions merge-status](growthbook_config-revisions_merge-status.md)	 - Get merge status for a draft revision
* [growthbook config-revisions publish](growthbook_config-revisions_publish.md)	 - Publish a draft revision
* [growthbook config-revisions rebase](growthbook_config-revisions_rebase.md)	 - Rebase a draft revision onto the current live config
* [growthbook config-revisions recall-review](growthbook_config-revisions_recall-review.md)	 - Recall a review request
* [growthbook config-revisions reopen](growthbook_config-revisions_reopen.md)	 - Reopen a discarded revision
* [growthbook config-revisions request-review](growthbook_config-revisions_request-review.md)	 - Request review for a draft revision
* [growthbook config-revisions revert](growthbook_config-revisions_revert.md)	 - Revert the config to a prior revision
* [growthbook config-revisions schedule-publish](growthbook_config-revisions_schedule-publish.md)	 - Schedule (or cancel) a deferred publish
* [growthbook config-revisions set-metadata](growthbook_config-revisions_set-metadata.md)	 - Update config metadata in a draft revision
* [growthbook config-revisions set-projection](growthbook_config-revisions_set-projection.md)	 - Set (or update) a config's per-source render projection on a draft
* [growthbook config-revisions set-schema](growthbook_config-revisions_set-schema.md)	 - Update or import the schema of a config draft revision
* [growthbook config-revisions set-value](growthbook_config-revisions_set-value.md)	 - Update the value of a config draft revision
* [growthbook config-revisions submit-review](growthbook_config-revisions_submit-review.md)	 - Submit a review on a draft revision
