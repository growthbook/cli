## growthbook experiments update

Update a single experiment

### Synopsis

Update a single experiment

```
growthbook experiments update [flags]
```

### Examples

```
  growthbook experiments update --id <id>
```

### Options

```
      --activation-metric string                     Users must convert on this metric before being included
      --analysis string                              Analysis summary or conclusions for the experiment. Maps to resultSummary.conclusions in the GET response.
      --archived                                     boolean flag
      --assignment-query-id string                   string value
      --attribution-model string                     Setting attribution model to "experimentDuration" is the same as selecting "Ignore Conversion Windows" for the Conversion Window Override. Setting it to "lookbackOverride" requires a lookbackOverride object to be provided. (options: firstExposure, experimentDuration, lookbackOverride)
      --auto-refresh                                 boolean flag
      --bandit-burn-in-unit string                   options: days, hours
      --bandit-burn-in-value float                   number value
      --bandit-conversion-window-unit string         options: days, hours
      --bandit-conversion-window-value float         number value
      --bandit-schedule-unit string                  options: days, hours
      --bandit-schedule-value float                  number value
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --bucket-version float                         number value
      --bypass-duplicate-key-check                   If true, allow updating the tracking key even if another experiment with the same tracking key already exist. This is ignored if the organization requires unique tracking keys as a rule.
      --custom-fields string                         value
      --custom-metric-slices string                  Custom slices that apply to ALL applicable metrics in the experiment
      --datasource-id string                         Can only be set if existing experiment does not have a datasource
      --decision-framework-settings string           Controls the decision framework and metric overrides for the experiment. Replaces the entire stored object on update (does not patch individual fields).
      --default-dashboard-id string                  ID of the default dashboard for this experiment.
      --description string                           Description of the experiment
      --disable-sticky-bucketing                     boolean flag
  -e, --exclude-from-payload                         If true, the experiment is excluded from the SDK payload. Maps to resultSummary.excludeFromPayload in the GET response.
  -f, --fallback-attribute string                    string value
  -g, --guardrail-metrics stringArray                list of values
      --hash-attribute string                        string value
      --hash-version float                           number value
  -h, --help                                         help for update
      --hypothesis string                            Hypothesis of the experiment
      --id string                                    The id of the requested resource [required]
      --in-progress-conversions string               options: loose, strict
  -l, --lookback-override string                     Controls the lookback override for the experiment. For type "window", value must be a non-negative number and valueUnit is required.
      --metric-overrides string                      Per-metric analysis overrides for this experiment. Replaces the entire stored array (does not patch individual entries).
      --metrics stringArray                          list of values
      --min-bucket-version float                     number value
  -n, --name string                                  Name of the experiment
      --owner string                                 The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --phases string                                list of values
      --post-stratification-enabled string           When null, the organization default is used.
      --precomputed-unit-dimension-ids stringArray   list of values
      --project string                               Project ID which the experiment belongs to
      --query-filter string                          WHERE clause to add to the default experiment query
      --regression-adjustment-enabled                Controls whether regression adjustment (CUPED) is enabled for experiment analyses
      --released-variation-id string                 The ID of the released variation. Maps to resultSummary.releasedVariationId in the GET response.
      --results string                               The result status of the experiment. Maps to resultSummary.status in the GET response. (options: dnf, won, lost, inconclusive)
      --secondary-metrics stringArray                list of values
      --segment-id string                            Only users in this segment will be included
      --sequential-testing-enabled                   Only applicable to frequentist analyses
      --sequential-testing-tuning-parameter float    number value
      --share-level string                           options: public, organization
      --stats-engine string                          options: bayesian, frequentist
      --status string                                options: draft, running, stopped
      --status-update-schedule string                JSON object
      --tags stringArray                             list of values
      --tracking-key string                          string value
      --type string                                  options: standard, multi-armed-bandit
  -v, --variations string                            list of values
  -w, --winner float                                 The index of the winning variation (0-indexed). Maps to resultSummary.winner (variation ID) in the GET response.
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

* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
