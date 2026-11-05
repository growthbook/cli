## growthbook ramp-schedules rollback-matching

Roll back matching ramp schedules

### Synopsis

Rolls back every `running` or `paused` ramp schedule that matches the query
filters. Built for an alerting tool's webhook, such as a Datadog monitor.

Each filter takes comma-separated values and matches any of them; different
filters must all match. At least one filter is required.

Returns each matched schedule and whether the rollback was applied or skipped,
for example because the API key can't publish its Feature Flag. Unknown body
fields are ignored, so a tool can send its own payload.

```
growthbook ramp-schedules rollback-matching [flags]
```

### Examples

```
  growthbook ramp-schedules rollback-matching
```

### Options

```
      --body string                                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.additional-properties string     value
      --body-param.reason Alert:                    Recorded on each ramp, prefixed with Alert: . Use the alert title and link.
  -c, --custom-fields customFields[<key>]=<value>   Custom field values, as customFields[<key>]=<value>. Comma-separate several values for one key.
  -e, --environment stringArray                     Comma-separated environments. Matches ramps whose rules run in any of them.
  -f, --feature stringArray                         Comma-separated Feature Flag ids.
  -h, --help                                        help for rollback-matching
  -p, --project stringArray                         Comma-separated Project ids. Matches ramps on Feature Flags in any of them.
  -r, --reason reason                               Used when the body has no reason, for tools whose payload can't be changed.
  -t, --tag stringArray                             Comma-separated tags. Matches ramps on Feature Flags with any of them.
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

* [growthbook ramp-schedules](growthbook_ramp-schedules.md)	 - Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring
