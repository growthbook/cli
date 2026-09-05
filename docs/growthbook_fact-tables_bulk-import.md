## growthbook fact-tables bulk-import

Bulk import fact tables, filters, and metrics

### Synopsis

Creates or updates Fact Tables, Fact Table filters, and Fact Metrics. Resources upsert by `id`. Pass `dryRun: true` to validate with zero writes. Not transactional: a live mid-loop failure returns HTTP 400 (403 for a permission failure) with write counts and `errors`.

```
growthbook fact-tables bulk-import [flags]
```

### Examples

```
  growthbook fact-tables bulk-import
```

### Options

```
      --body string                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --default-managed-by managedBy   Fallback managedBy for Fact Tables and Fact Metrics that omit the field. Defaults to `"api"`. Filters inherit `"api"` only when the parent Fact Table is api-managed. (options: , api, admin)
      --dry-run                        Validate with zero writes.
      --fact-metrics string            list of values
      --fact-table-filters string      list of values
      --fact-tables string             list of values
  -h, --help                           help for bulk-import
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
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

* [growthbook fact-tables](growthbook_fact-tables.md)	 - Fact Tables describe the shape of your data warehouse tables
