## growthbook feature-revisions post-feature-revision-rebase-preview

Preview a rebase without applying it

### Synopsis

Dry-run of the rebase: runs the same three-way merge with the supplied `conflictResolutions` and returns every conflict (resolved and unresolved) plus the merged result once all are resolved — without modifying the draft. Use it to iterate on resolutions before committing them via the rebase endpoint.

```
growthbook feature-revisions post-feature-revision-rebase-preview [flags]
```

### Examples

```
  growthbook feature-revisions post-feature-revision-rebase-preview --id <id> --version-param <value>
```

### Options

```
      --body string                                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --conflict-resolutions defaultValue              Map of conflict key → resolution. Keys come from the returned conflicts: defaultValue, `prerequisites`, `archived`, `holdout`, `environmentsEnabled.<env>`, `metadata.<field>`, `rules.<ruleId>`, and `rules.order`. `overwrite` keeps the draft's version of that item; `discard` keeps live's. The blanket `rules` key applies one strategy to all rule-level conflicts.
      --expected-draft-date-updated draftDateUpdated   Optimistic-concurrency guard for the draft side: the draft's draftDateUpdated timestamp as returned by merge-status or rebase preview. If the draft has been modified since (e.g. by a co-author), the request fails with `409` instead of applying resolutions against changed draft content.
      --expected-live-version 409                      Optimistic-concurrency guard: the live version the resolutions were authored against (as returned by merge-status or rebase preview). If live has since moved, the request fails with 409 instead of applying resolutions to different conflicts.
  -h, --help                                           help for post-feature-revision-rebase-preview
  -i, --id string                                      [required]
  -v, --version-param string                           [required]
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
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
