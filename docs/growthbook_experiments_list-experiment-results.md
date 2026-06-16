## growthbook experiments list-experiment-results

Get latest results for many experiments

### Synopsis

Returns the latest non-dimension snapshot for each experiment matching the filters. Use this to scan results across a portfolio in one call.

Pagination semantics:
- `total` is the count of experiments matching the filters.
- `count` is the length of the returned `experimentResults` array.
- Experiments without a completed snapshot are omitted from `experimentResults`, so `count` may be less than the page slice and a page may legitimately return `count: 0` while `hasMore: true`.
- `hasMore` and `nextOffset` advance over experiments matching the filters, not over returned results.

Use the per-experiment `GET /experiments/{id}/results` endpoint to inspect specific phases or dimensions.

```
growthbook experiments list-experiment-results [flags]
```

### Examples

```
  growthbook experiments list-experiment-results
```

### Options

```
      --datasource-id string   Filter by Data Source
  -h, --help                   help for list-experiment-results
  -l, --limit int              The number of items to return (default 10)
      --offset int             How many items to skip (use in conjunction with limit for pagination)
  -p, --project-id string      Filter by project id
  -s, --status string          options: draft, running, stopped
  -t, --tracking-key string    Filter by experiment tracking key
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
