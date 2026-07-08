## growthbook feature-revisions get-feature-revision-diff

Diff a revision against another revision

### Synopsis

Returns a schema-keyed JSON diff between this revision and a baseline. The same shapes the in-app review surface produces under `Copy as → Minimal JSON` / `Full JSON`: `minimal` lists only what changed (with id-keyed arrays bucketed into added/removed/modified items and reorder detection), while `full` returns the complete before/after content of the revision. Lifecycle fields (version, status, comment, date, createdBy, publishedBy) are excluded from the diff body and echoed via `from` / `to` instead. Defaults to diffing against the revision's own `baseVersion`; pass `?base=live` to diff against the current live revision, or `?base=<version>` for an arbitrary historical one.

```
growthbook feature-revisions get-feature-revision-diff [flags]
```

### Examples

```
  growthbook feature-revisions get-feature-revision-diff --id <id> --version-param <value>
```

### Options

```
  -b, --base baseVersion       Compare against: baseVersion (default — the revision's own `baseVersion`, matches the in-app review view), `live` (the currently-live revision), or an integer version (an arbitrary historical revision).
  -f, --format-param minimal   minimal (default) returns only what changed, with id-keyed arrays bucketed into added/removed/modified items. `full` returns the complete before/after content of the revision. (options: minimal, full)
  -h, --help                   help for get-feature-revision-diff
  -i, --id string              [required]
  -v, --version-param string   [required]
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

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
