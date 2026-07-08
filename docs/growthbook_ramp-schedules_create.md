## growthbook ramp-schedules create

Create a ramp schedule

### Synopsis

Create a ramp schedule

```
growthbook ramp-schedules create [flags]
```

### Examples

```
  growthbook ramp-schedules create
```

### Options

```
      --body string                       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --cutoff-date string                date/time value
      --end-actions string                list of values
      --environment string                string value
      --experiment-health-action string   options: rollback, hold, warn
  -f, --feature-id string                 string value
  -h, --help                              help for create
  -l, --lockdown-config string            JSON object
  -m, --monitoring-config string          JSON object
  -n, --name string                       string value
  -r, --rule-id string                    string value
      --start-actions string              list of values
      --start-date string                 date/time value
      --steps string                      list of values
  -t, --template-id string                string value
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
      --password string               HTTP Basic auth: your Secret Key as the username, with an empty password. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: your Secret Key as the username, with an empty password. username
```

### SEE ALSO

* [growthbook ramp-schedules](growthbook_ramp-schedules.md)	 - Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring
