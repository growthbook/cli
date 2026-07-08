## growthbook feature-revisions-v2 list-revisions-v2

List revisions across all features

### Synopsis

Returns a paginated list of feature revisions across all features in the organization. Use the `featureId` query parameter to filter to a single feature. Revision `rules` is a flat array with per-rule scope.

```
growthbook feature-revisions-v2 list-revisions-v2 [flags]
```

### Examples

```
  growthbook feature-revisions-v2 list-revisions-v2
```

### Options

```
      --archived false           Whether to include revisions for archived features. Defaults to false (non-archived features only). Pass `true` to include revisions for archived features alongside non-archived ones.
      --author string            string value
  -f, --feature-id string        string value
  -h, --help                     help for list-revisions-v2
  -l, --limit int                The number of items to return (default 10)
  -m, --mine string              If true, return only revisions authored by or contributed to by the calling user.
      --offset int               How many items to skip (use in conjunction with limit for pagination)
      --skip-pagination string   If true, return all matching items and ignore limit/offset.
                                 Self-hosted only. Has no effect unless API_ALLOW_SKIP_PAGINATION is set to true or 1.
      --status string            JSON value (one of: string | array of string)
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

* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
