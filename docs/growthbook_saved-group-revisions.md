## growthbook saved-group-revisions

Draft revisions for saved groups, including pending changes, approvals, and lifecycle (publish, discard, revert)

### Synopsis

Draft revisions for saved groups, including pending changes, approvals, and lifecycle (publish, discard, revert).

Most callers can interact with these endpoints via shorthand actions (`/items/add`, `/items/remove`, single-field PUTs) instead of authoring JSON Patch ops directly. Pass `version: "new"` on edit endpoints to auto-create a draft.

```
growthbook saved-group-revisions [flags]
```

### Options

```
  -h, --help   help for saved-group-revisions
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
* [growthbook saved-group-revisions add-items](growthbook_saved-group-revisions_add-items.md)	 - Append items to a list saved group draft revision
* [growthbook saved-group-revisions archive](growthbook_saved-group-revisions_archive.md)	 - Stage an archive/unarchive in a draft revision
* [growthbook saved-group-revisions create](growthbook_saved-group-revisions_create.md)	 - Create a draft revision
* [growthbook saved-group-revisions discard](growthbook_saved-group-revisions_discard.md)	 - Discard a draft revision
* [growthbook saved-group-revisions get](growthbook_saved-group-revisions_get.md)	 - Get a single saved group revision
* [growthbook saved-group-revisions latest](growthbook_saved-group-revisions_latest.md)	 - Get the most recent active draft revision
* [growthbook saved-group-revisions list](growthbook_saved-group-revisions_list.md)	 - List saved-group revisions across the organization
* [growthbook saved-group-revisions list-for-saved-group](growthbook_saved-group-revisions_list-for-saved-group.md)	 - List revisions for a saved group
* [growthbook saved-group-revisions merge-status](growthbook_saved-group-revisions_merge-status.md)	 - Get merge status for a draft revision
* [growthbook saved-group-revisions publish](growthbook_saved-group-revisions_publish.md)	 - Publish a draft revision
* [growthbook saved-group-revisions rebase](growthbook_saved-group-revisions_rebase.md)	 - Rebase a draft revision onto the current live saved group
* [growthbook saved-group-revisions remove-items](growthbook_saved-group-revisions_remove-items.md)	 - Remove items from a list saved group draft revision
* [growthbook saved-group-revisions request-review](growthbook_saved-group-revisions_request-review.md)	 - Request review for a draft revision
* [growthbook saved-group-revisions revert](growthbook_saved-group-revisions_revert.md)	 - Revert the saved group to a prior revision
* [growthbook saved-group-revisions set-condition](growthbook_saved-group-revisions_set-condition.md)	 - Update the condition of a condition saved group draft revision
* [growthbook saved-group-revisions set-metadata](growthbook_saved-group-revisions_set-metadata.md)	 - Update saved group metadata in a draft revision
* [growthbook saved-group-revisions set-values](growthbook_saved-group-revisions_set-values.md)	 - Replace the values list in a list saved group draft revision
* [growthbook saved-group-revisions submit-review](growthbook_saved-group-revisions_submit-review.md)	 - Submit a review on a draft revision
