## growthbook metrics list-metric-experiments

Get results for all experiments that use a metric

### Synopsis

Returns, for each experiment that uses the given metric (directly or via a metric group), the per-variation results for that metric from the latest snapshot. Supports the same filtering as the experiment list views via a raw search string or structured query params. Note: at most the 1000 most recent experiments using the metric are considered; filters and pagination are applied within that set, so results may be incomplete for metrics used by more than 1000 experiments.

```
growthbook metrics list-metric-experiments [flags]
```

### Examples

```
  growthbook metrics list-metric-experiments --id <id>
```

### Options

```
  -b, --bandits string                  When true, return only multi-armed bandits; when false, exclude them (options: true, false)
  -e, --end-date string                 Only include experiments that have a phase which ended on or before this date
  -h, --help                            help for list-metric-experiments
  -i, --id string                       The id of the requested resource [required]
  -l, --limit int                       The number of items to return (default 10)
      --offset int                      How many items to skip (use in conjunction with limit for pagination)
      --owner string                    Filter by comma-separated owner ids, names, or emails
  -p, --project-id string               Filter by comma-separated project ids or names
      --q status:running tag:checkout   Raw experiment search/filter string (same syntax as the app's experiment list filters, e.g. status:running tag:checkout). Negation (`!`) and operators (`~`, `^`, `>`, `<`, `=`) are not supported and return a 400
  -r, --result string                   Filter by comma-separated results (won, lost, inconclusive, dnf)
      --start-date string               Only include experiments that have a phase which ended on or after this date
      --status string                   Filter by comma-separated statuses (draft, running, stopped)
      --tag string                      Filter by comma-separated tags
      --type string                     Filter by comma-separated experiment types (feature, visualChange, redirect)
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

* [growthbook metrics](growthbook_metrics.md)	 - Metrics used as goals and guardrails for experiments
