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
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password string               HTTP Basic auth: your Secret Key as the username, with an empty password. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: your Secret Key as the username, with an empty password. username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook saved-group-revisions create](growthbook_saved-group-revisions_create.md)	 - Create a draft revision
* [growthbook saved-group-revisions get](growthbook_saved-group-revisions_get.md)	 - Get a single saved group revision
* [growthbook saved-group-revisions get-saved-group-revision-latest](growthbook_saved-group-revisions_get-saved-group-revision-latest.md)	 - Get the most recent active draft revision
* [growthbook saved-group-revisions get-saved-group-revision-merge-status](growthbook_saved-group-revisions_get-saved-group-revision-merge-status.md)	 - Get merge status for a draft revision
* [growthbook saved-group-revisions list](growthbook_saved-group-revisions_list.md)	 - List saved-group revisions across the organization
* [growthbook saved-group-revisions list-for-saved-group](growthbook_saved-group-revisions_list-for-saved-group.md)	 - List revisions for a saved group
* [growthbook saved-group-revisions post-saved-group-revision-discard](growthbook_saved-group-revisions_post-saved-group-revision-discard.md)	 - Discard a draft revision
* [growthbook saved-group-revisions post-saved-group-revision-items-add](growthbook_saved-group-revisions_post-saved-group-revision-items-add.md)	 - Append items to a list saved group draft revision
* [growthbook saved-group-revisions post-saved-group-revision-items-remove](growthbook_saved-group-revisions_post-saved-group-revision-items-remove.md)	 - Remove items from a list saved group draft revision
* [growthbook saved-group-revisions post-saved-group-revision-publish](growthbook_saved-group-revisions_post-saved-group-revision-publish.md)	 - Publish a draft revision
* [growthbook saved-group-revisions post-saved-group-revision-rebase](growthbook_saved-group-revisions_post-saved-group-revision-rebase.md)	 - Rebase a draft revision onto the current live saved group
* [growthbook saved-group-revisions post-saved-group-revision-request-review](growthbook_saved-group-revisions_post-saved-group-revision-request-review.md)	 - Request review for a draft revision
* [growthbook saved-group-revisions post-saved-group-revision-revert](growthbook_saved-group-revisions_post-saved-group-revision-revert.md)	 - Revert the saved group to a prior revision
* [growthbook saved-group-revisions post-saved-group-revision-submit-review](growthbook_saved-group-revisions_post-saved-group-revision-submit-review.md)	 - Submit a review on a draft revision
* [growthbook saved-group-revisions put-saved-group-revision-archive](growthbook_saved-group-revisions_put-saved-group-revision-archive.md)	 - Stage an archive/unarchive in a draft revision
* [growthbook saved-group-revisions put-saved-group-revision-condition](growthbook_saved-group-revisions_put-saved-group-revision-condition.md)	 - Update the condition of a condition saved group draft revision
* [growthbook saved-group-revisions put-saved-group-revision-metadata](growthbook_saved-group-revisions_put-saved-group-revision-metadata.md)	 - Update saved group metadata in a draft revision
* [growthbook saved-group-revisions put-saved-group-revision-values](growthbook_saved-group-revisions_put-saved-group-revision-values.md)	 - Replace the values list in a list saved group draft revision
