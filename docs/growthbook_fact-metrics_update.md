## growthbook fact-metrics update

Update a single fact metric

### Synopsis

Update a single fact metric

```
growthbook fact-metrics update [flags]
```

### Examples

```
  growthbook fact-metrics update --id <id>
```

### Options

```
  -a, --archived                                     Set to true to archive the metric. Archived metrics are hidden by default in the UI and excluded from new experiments.
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --capping-settings string                      Controls how outliers are handled
      --denominator string                           Only when metricType is 'ratio'
      --description string                           string value
      --display-as-percentage                        If true and the metric is a ratio or dailyParticipation metric, variation means will be displayed as a percentage. Defaults to true for dailyParticipation metrics and false for ratio metrics.
  -f, --funnel-settings string                       Funnel metric settings (required when metricType is "funnel")
  -h, --help                                         help for update
      --id string                                    The id of the requested resource [required]
      --inverse                                      Set to true for things like Bounce Rate, where you want the metric to decrease
      --managed-by string                            Set this to "api" to disable editing in the GrowthBook UI (options: , api, admin)
      --max-percent-change float                     Maximum percent change to consider uplift significant, as a proportion (e.g. put 0.5 for 50%)
      --metric-auto-slices stringArray               Array of slice column names that will be automatically included in metric analysis. This is an enterprise feature.
      --metric-type string                           options: proportion, retention, mean, quantile, ratio, dailyParticipation, funnel
      --min-percent-change float                     Minimum percent change to consider uplift significant, as a proportion (e.g. put 0.005 for 0.5%)
      --min-sample-size float                        number value
      --name string                                  string value
      --numerator string                             JSON object
      --owner string                                 The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --prior-settings string                        Controls the bayesian prior for the metric. If omitted, organization defaults will be used.
      --projects stringArray                         list of values
      --quantile-settings string                     Controls the settings for quantile metrics (mandatory if metricType is "quantile")
      --regression-adjustment-settings string        Controls the regression adjustment (CUPED) settings for the metric
      --replaces stringArray                         Ids of older metrics (legacy or fact) that this metric supersedes, for example the legacy metric it was migrated from. Cannot include this metric's own id. Informational only - GrowthBook uses it to link the old and new definitions in the UI and to keep showing results from a snapshot that was created before an experiment switched to this metric. This field can only be set through the API.
      --risk-threshold-danger float                  No longer used. Threshold for Risk to be considered too high, as a proportion (e.g. put 0.0125 for 1.25%). <br/> Must be a non-negative number.
      --risk-threshold-success riskThresholdDanger   No longer used. Threshold for Risk to be considered low enough, as a proportion (e.g. put 0.0025 for 0.25%). <br/> Must be a non-negative number and must not be higher than riskThresholdDanger.
      --tags stringArray                             list of values
      --target-mde float                             The percentage change that you want to reliably detect before ending an experiment, as a proportion (e.g. put 0.1 for 10%). This is used to estimate the "Days Left" for running experiments.
  -w, --window-settings string                       Controls the conversion window for the metric
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

* [growthbook fact-metrics](growthbook_fact-metrics.md)	 - Fact Metrics are metrics built on top of Fact Table definitions
