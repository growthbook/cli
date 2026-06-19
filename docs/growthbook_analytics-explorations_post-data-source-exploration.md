## growthbook analytics-explorations post-data-source-exploration

Create a Data Source based visualization

### Synopsis

Create a Data Source based visualization

```
growthbook analytics-explorations post-data-source-exploration [flags]
```

### Examples

```
  growthbook analytics-explorations post-data-source-exploration --datasource <value> --dimensions '[{"dimensionType":"dynamic","column":"<value>","maxValues":4026.37}]' --chart-type table --date-range '{"predefined":"last7Days"}' --dataset '{"type":"data_source","table":"<value>","path":"/net","timestampColumn":"<value>","columnTypes":{"key":"string"},"values":[]}'
```

### Options

```
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --cache preferred     Controls cache behavior for this exploration: preferred (default) returns a cached result if one exists, otherwise runs a new query; `never` always runs a new query, ignoring any cached results; `required` only returns a cached result, if none exists returns exploration: null with a message (options: preferred, required, never)
      --chart-type string   options: line, area, timeseries-table, table, bar, stackedBar, horizontalBar, stackedHorizontalBar, bigNumber [required]
      --dataset string      [required]
      --datasource string   ID of the datasource to query [required]
      --date-range string   [required]
      --dimensions string   [required]
  -h, --help                help for post-data-source-exploration
  -s, --show-as string      options: total, per_unit
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth                   If using Bearer auth, pass the Secret Key as the token:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                                      ```
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
      --password                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       username
```

### SEE ALSO

* [growthbook analytics-explorations](growthbook_analytics-explorations.md)	 - Operations for analytics-explorations
