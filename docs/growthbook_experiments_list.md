## growthbook experiments list

Get all experiments

### Synopsis

Get all experiments

```
growthbook experiments list [flags]
```

### Examples

```
  growthbook experiments list
```

### Options

```
  -a, --all                             Automatically paginate and fetch all results (streams NDJSON for JSON output)
      --archived true                   Filter by archived status. Set to true to return only archived experiments, `false` to exclude them. If omitted, both archived and non-archived experiments are returned.
  -b, --bandits string                  When true, return only multi-armed bandits; when false, exclude them (options: true, false)
      --datasource-id string            Filter by Data Source
  -e, --experiment-id trackingKey       Filter the returned list by the experiment tracking key (not the internal experiment ID). Note, this was deprecated to help reduce confusion, consider using trackingKey instead, which is functionally identical. You cannot use both params at the same time.
  -h, --help                            help for list
  -i, --implementation-type bandits     Filter by comma-separated implementation types (feature, visualChange, redirect) — the kinds of changes linked to the experiment. To filter standard experiments vs bandits, use bandits instead
  -l, --limit int                       The number of items to return (default 10)
      --max-pages int                   Maximum number of pages to fetch when using --all (0 = no limit)
  -m, --metric-id string                Filter by comma-separated metric ids. Matches experiments that use a metric as a goal, secondary, or guardrail metric
      --offset int                      How many items to skip (use in conjunction with limit for pagination)
      --owner string                    Filter by comma-separated owner ids, names, or emails
  -p, --project-id string               Filter by project id
      --q status:running tag:checkout   Raw experiment search/filter string (same syntax as the app's experiment list filters, e.g. status:running tag:checkout). Negation (`!`) and operators (`~`, `^`, `>`, `<`, `=`) are not supported and return a 400
  -r, --result stringArray              Filter by comma-separated results (won, lost, inconclusive, dnf). Matches the experiment's recorded result — set when an experiment is stopped and retained if it's later restarted, so running experiments can match too
      --sort-by string                  Field to sort the results by (options: dateCreated, dateUpdated, name) (default "dateCreated")
      --sort-order sortBy               Sort direction (used with sortBy) (options: asc, desc) (default "asc")
      --status string                   options: draft, running, stopped
      --tag string                      Filter by comma-separated tags
      --tracking-key string             Filter by experiment tracking key
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
