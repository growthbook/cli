## growthbook fact-tables list-fact-table-filters

Get all filters for a fact table

### Synopsis

Get all filters for a fact table

```
growthbook fact-tables list-fact-table-filters [flags]
```

### Examples

```
  growthbook fact-tables list-fact-table-filters --fact-table-id <id>
```

### Options

```
  -f, --fact-table-id string   Specify a specific fact table [required]
  -h, --help                   help for list-fact-table-filters
  -l, --limit int              The number of items to return (default 10)
      --offset int             How many items to skip (use in conjunction with limit for pagination)
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

* [growthbook fact-tables](growthbook_fact-tables.md)	 - Fact Tables describe the shape of your data warehouse tables
