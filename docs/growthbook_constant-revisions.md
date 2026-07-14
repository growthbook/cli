## growthbook constant-revisions

Draft revisions for constants, including pending changes, approvals, and lifecycle (publish, discard, revert)

### Synopsis

Draft revisions for constants, including pending changes, approvals, and lifecycle (publish, discard, revert). Pass `version: "new"` on edit endpoints to auto-create a draft.

```
growthbook constant-revisions [flags]
```

### Options

```
  -h, --help   help for constant-revisions
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
* [growthbook constant-revisions archive](growthbook_constant-revisions_archive.md)	 - Stage an archive/unarchive in a draft revision
* [growthbook constant-revisions create](growthbook_constant-revisions_create.md)	 - Create a draft revision
* [growthbook constant-revisions discard](growthbook_constant-revisions_discard.md)	 - Discard a draft revision
* [growthbook constant-revisions get](growthbook_constant-revisions_get.md)	 - Get a single constant revision
* [growthbook constant-revisions latest](growthbook_constant-revisions_latest.md)	 - Get the most recent active draft revision
* [growthbook constant-revisions list](growthbook_constant-revisions_list.md)	 - List constant revisions across the organization
* [growthbook constant-revisions list-for-constant](growthbook_constant-revisions_list-for-constant.md)	 - List revisions for a constant
* [growthbook constant-revisions merge-status](growthbook_constant-revisions_merge-status.md)	 - Get merge status for a draft revision
* [growthbook constant-revisions publish](growthbook_constant-revisions_publish.md)	 - Publish a draft revision
* [growthbook constant-revisions rebase](growthbook_constant-revisions_rebase.md)	 - Rebase a draft revision onto the current live constant
* [growthbook constant-revisions request-review](growthbook_constant-revisions_request-review.md)	 - Request review for a draft revision
* [growthbook constant-revisions revert](growthbook_constant-revisions_revert.md)	 - Revert the constant to a prior revision
* [growthbook constant-revisions set-metadata](growthbook_constant-revisions_set-metadata.md)	 - Update constant metadata in a draft revision
* [growthbook constant-revisions set-value](growthbook_constant-revisions_set-value.md)	 - Update the value of a constant draft revision
* [growthbook constant-revisions submit-review](growthbook_constant-revisions_submit-review.md)	 - Submit a review on a draft revision
