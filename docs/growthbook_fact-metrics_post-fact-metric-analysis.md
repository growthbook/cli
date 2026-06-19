## growthbook fact-metrics post-fact-metric-analysis

Create a fact metric analysis

### Synopsis

Create a fact metric analysis

```
growthbook fact-metrics post-fact-metric-analysis [flags]
```

### Examples

```
  growthbook fact-metrics post-fact-metric-analysis --id <id>
```

### Options

```
      --additional-denominator-filters stringArray   We support passing in adhoc filters for an analysis that don't live on the metric itself. These are in addition to the metric's filters. To use this, you can pass in an array of Fact Table Filter Ids.
      --additional-numerator-filters stringArray     We support passing in adhoc filters for an analysis that don't live on the metric itself. These are in addition to the metric's filters. To use this, you can pass in an array of Fact Table Filter Ids.
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                                         help for post-fact-metric-analysis
  -i, --id string                                    The fact metric id to analyze [required]
  -l, --lookback-days float                          Number of days to look back for the analysis. Defaults to 30.
      --population-id string                         The ID of the population (e.g., segment ID) when populationType is not 'factTable'. Defaults to null.
      --population-type string                       The type of population to analyze. Defaults to 'factTable', meaning the analysis will return the metric value for all units found in the fact table. (options: factTable, segment)
      --use-cache                                    Whether to use a cached query if one exists. Defaults to true.
      --user-id-type string                          The identifier type to use for the analysis. If not provided, defaults to the first available identifier type in the fact table.
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

* [growthbook fact-metrics](growthbook_fact-metrics.md)	 - Fact Metrics are metrics built on top of Fact Table definitions
