## growthbook reports put-report-settings

Update report analysis settings

### Synopsis

Updates the analysis settings for an existing report. Changes are staged and do not take effect until you call `POST /reports/:id/refresh`.

```
growthbook reports put-report-settings [flags]
```

### Examples

```
  growthbook reports put-report-settings --id <id>
```

### Options

```
      --activation-metric string                     Activation metric ID
      --attribution-model string                     Metric conversion window attribution model (options: firstExposure, experimentDuration, lookbackOverride)
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --coverage float                               Traffic coverage (0–1) for the latest phase. Used when computing scaled impact.
      --custom-metric-slices string                  Custom metric slice definitions
      --date-ended null                              Analysis end date (ISO 8601). Pass null to clear the end date and analyze through today.
      --date-started string                          Analysis start date (ISO 8601)
      --difference-type string                       How lifts are expressed in results (options: relative, absolute, scaled)
      --dimension string                             Dimension to cut results by
      --goal-metrics stringArray                     Goal metric IDs
      --guardrail-metrics stringArray                Guardrail metric IDs
  -h, --help                                         help for put-report-settings
  -i, --id string                                    The id of the requested resource [required]
  -l, --lookback-override string                     JSON value (variants: date: { value: value }, window: { value: number, valueUnit: string })
      --lookback-override.date string                putReportSettings_lookbackOverride_Date variant as JSON
      --lookback-override.window string              putReportSettings_lookbackOverride_Window variant as JSON
      --lookback-override.window.value float         [required]
      --lookback-override.window.value-unit string   options: minutes, hours, days, weeks [required]
  -m, --metric-overrides string                      Per-metric window, risk, and regression-adjustment overrides
      --query-filter string                          Raw SQL WHERE clause added to the exposure query
  -r, --regression-adjustment-enabled                Enable CUPED regression adjustment
      --secondary-metrics stringArray                Secondary metric IDs
      --segment string                               Segment ID to filter users by
      --sequential-testing-enabled                   Enable sequential testing
      --sequential-testing-tuning-parameter float    Tuning parameter for sequential testing (frequentist only)
      --skip-partial-data                            When true, exclude users who have not completed the full conversion window
      --stats-engine string                          Stats engine override (options: bayesian, frequentist)
  -v, --variations string                            Override variation names, keys, or traffic weights used in this report. Weights are merged into the latest phase. Changes take effect on the next refresh.
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

* [growthbook reports](growthbook_reports.md)	 - Custom analysis reports built on top of experiment snapshots
