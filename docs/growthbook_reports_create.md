## growthbook reports create

Create a new report

### Synopsis

Create a new report

```
growthbook reports create [flags]
```

### Examples

```
  growthbook reports create --experiment-id <id>
```

### Options

```
      --activation-metric string                     Activation metric ID
      --attribution-model string                     Metric conversion window attribution model. Defaults to experiment setting. (options: firstExposure, experimentDuration, lookbackOverride)
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --custom-metric-slices string                  Custom metric slice definitions
      --date-ended string                            Analysis end date (ISO 8601)
      --date-started string                          Analysis start date (ISO 8601)
      --description string                           Report description
      --difference-type string                       How lifts are expressed in results. Defaults to experiment setting. (options: relative, absolute, scaled)
      --dimension string                             Dimension to cut results by
  -e, --experiment-id string                         The experiment to create a report for [required]
      --goal-metrics stringArray                     Goal metric IDs (defaults to experiment's goal metrics)
      --guardrail-metrics stringArray                Guardrail metric IDs (defaults to experiment's guardrail metrics)
  -h, --help                                         help for create
  -l, --lookback-override string                     JSON value (variants: date: { value: value }, window: { value: number, valueUnit: string })
      --lookback-override.date string                postReport_lookbackOverride_Date variant as JSON
      --lookback-override.window string              postReport_lookbackOverride_Window variant as JSON
      --lookback-override.window.value float         [required]
      --lookback-override.window.value-unit string   options: minutes, hours, days, weeks [required]
  -m, --metric-overrides string                      Per-metric window, risk, and regression-adjustment overrides
      --query-filter string                          Raw SQL WHERE clause added to the exposure query. Defaults to experiment setting.
  -r, --regression-adjustment-enabled                Enable CUPED regression adjustment
      --secondary-metrics stringArray                Secondary metric IDs (defaults to experiment's secondary metrics)
      --segment string                               Segment ID to filter users by. Defaults to experiment setting.
      --sequential-testing-enabled                   Enable sequential testing
      --sequential-testing-tuning-parameter float    Tuning parameter for sequential testing (frequentist only)
      --share-level private                          Visibility of the created report. Defaults to private. Set to `public` to receive a shareable `shareUrl` in the response. (options: public, organization, private)
      --skip-partial-data                            When true, exclude users who have not completed the full conversion window.
      --stats-engine string                          Stats engine override (options: bayesian, frequentist)
  -t, --title string                                 Report title (defaults to experiment name)
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

* [growthbook reports](growthbook_reports.md)	 - Custom analysis reports built on top of experiment snapshots
