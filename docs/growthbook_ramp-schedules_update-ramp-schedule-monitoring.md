## growthbook ramp-schedules update-ramp-schedule-monitoring

Update ramp monitoring configuration

### Synopsis

Replaces the monitoring configuration. Metric IDs, snapshot cadence, and health-action thresholds (`srmAction`, `noTrafficAction`, etc.) can be updated at any time.

`datasourceId` and `exposureQueryId` are locked once monitoring starts — stop and recreate the schedule to change the data source.

Changes to guardrail or signal metric IDs take effect on the next analysis run.

```
growthbook ramp-schedules update-ramp-schedule-monitoring [flags]
```

### Examples

```
  growthbook ramp-schedules update-ramp-schedule-monitoring --id <id> --datasource-id <id> --exposure-query-id <id> --guardrail-metric-ids '["<value 1>","<value 2>","<value 3>"]'
```

### Options

```
  -a, --auto-update                                     boolean flag
      --body string                                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --datasource-id string                            [required]
  -e, --exposure-query-id string                        [required]
  -g, --guardrail-metric-ids stringArray                [required]
  -h, --help                                            help for update-ramp-schedule-monitoring
  -i, --id string                                       [required]
      --monitoring-mode string                          options: auto, manual
      --multiple-exposure-action string                 options: rollback, hold, warn
      --no-traffic-action string                        options: rollback, hold, warn
      --no-traffic-grace-period-hours noTrafficAction   How long to wait for traffic before applying noTrafficAction. Defaults to 24 hours when null or not set.
      --signal-metric-ids stringArray                   list of values
      --srm-action string                               options: rollback, hold, warn
  -u, --update-schedule-minutes string                  number value
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
