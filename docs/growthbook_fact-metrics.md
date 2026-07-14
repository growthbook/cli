## growthbook fact-metrics

Fact Metrics are metrics built on top of Fact Table definitions

### Synopsis

Fact Metrics are metrics built on top of Fact Table definitions

```
growthbook fact-metrics [flags]
```

### Options

```
  -h, --help   help for fact-metrics
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
* [growthbook fact-metrics create](growthbook_fact-metrics_create.md)	 - Create a single fact metric
* [growthbook fact-metrics delete](growthbook_fact-metrics_delete.md)	 - Deletes a single fact metric
* [growthbook fact-metrics get](growthbook_fact-metrics_get.md)	 - Get a single fact metric
* [growthbook fact-metrics list](growthbook_fact-metrics_list.md)	 - Get all fact metrics
* [growthbook fact-metrics post-fact-metric-analysis](growthbook_fact-metrics_post-fact-metric-analysis.md)	 - Create a fact metric analysis
* [growthbook fact-metrics update](growthbook_fact-metrics_update.md)	 - Update a single fact metric
