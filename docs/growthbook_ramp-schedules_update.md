## growthbook ramp-schedules update

Update a single rampSchedule

### Synopsis

Updates the name, steps, endActions, startDate, or cutoffDate of a ramp schedule.

Only allowed when the schedule is in `pending`, `ready`, or `paused` status.

**targetId shorthand**: When providing `steps` or `endActions`, you may omit `targetId`
(or pass `"t1"`) in each action. If the schedule has exactly one active target, the server
will resolve it automatically. For schedules with multiple targets, provide the explicit
target UUID from `targets[].id`.

**Coverage on monitored steps**: See the create endpoint description for details
on how `coverage` is interpreted for monitored steps (total enrollment, not
variation-1 exposure).

```
growthbook ramp-schedules update [flags]
```

### Examples

```
  growthbook ramp-schedules update --id <id>
```

### Options

```
      --body string                       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --cutoff-date string                date/time value
      --end-actions string                list of values
      --experiment-health-action string   options: rollback, hold, warn
  -h, --help                              help for update
  -i, --id string                         [required]
  -l, --lockdown-config string            When mode is 'locked', blocks all feature edits while the ramp is actively running.
  -m, --monitoring-config string          JSON object
  -n, --name string                       string value
      --start-actions string              list of values
      --start-date string                 date/time value
      --steps string                      list of values
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

* [growthbook ramp-schedules](growthbook_ramp-schedules.md)	 - Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring
