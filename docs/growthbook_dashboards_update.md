## growthbook dashboards update

Update a single dashboard

### Synopsis

Update a single dashboard

```
growthbook dashboards update [flags]
```

### Examples

```
  growthbook dashboards update --id <id>
```

### Options

```
  -b, --blocks string                       list of values
      --body string                         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --edit-level string                   Dashboards that are "published" are editable by organization members with appropriate permissions (options: published, private)
      --enable-auto-updates                 If enabled for a General Dashboard, also requires an updateSchedule
  -h, --help                                help for update
  -i, --id string                           [required]
  -p, --projects stringArray                General Dashboards only, Experiment Dashboards use the experiment's projects
  -s, --share-level string                  General Dashboards only. Dashboards that are "published" are viewable by organization members with appropriate permissions (options: published, private)
  -t, --title string                        The display name of the Dashboard
  -u, --update-schedule string              JSON value (variants: stale: { hours: number }, cron: { cron: string })
      --update-schedule.cron string         updateDashboard_updateSchedule_Cron variant as JSON
      --update-schedule.cron.cron string    [required]
      --update-schedule.stale string        updateDashboard_updateSchedule_Stale variant as JSON
      --update-schedule.stale.hours float   [required]
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

* [growthbook dashboards](growthbook_dashboards.md)	 - Operations for dashboards
