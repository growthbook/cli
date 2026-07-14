## growthbook analytics-explorations post-metric-exploration

Create a Metric based visualization

### Synopsis

Create a Metric based visualization

```
growthbook analytics-explorations post-metric-exploration [flags]
```

### Examples

```
  growthbook analytics-explorations post-metric-exploration --datasource <value> --dimensions '[{"dimensionType":"static","column":"<value>","values":["<value 1>","<value 2>","<value 3>"]}]' --chart-type table --date-range '{"predefined":"last30Days"}' --dataset '{"type":"metric","values":[{"name":"<value>","rowFilters":[],"type":"metric","metricId":"<id>","unit":null,"denominatorUnit":"<value>"}]}'
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
  -h, --help                help for post-metric-exploration
  -s, --show-as string      options: total, per_unit
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

* [growthbook analytics-explorations](growthbook_analytics-explorations.md)	 - Operations for analytics-explorations
