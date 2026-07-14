## growthbook ramp-schedules update-ramp-schedule-steps

Update ramp schedule steps

### Synopsis

Fully replaces the steps array for a ramp schedule. Only allowed when the schedule is in a non-running, non-terminal state (`ready`, `pending`, or `paused`). Pause a running schedule first; restart a terminal schedule first.

**Step actions** (coverage/targeting patches) are not accepted here — they change the SDK payload and must go through a feature revision draft. Existing step actions are preserved for each position. Use `PUT /v2/features/:id/revisions/:version/rules/:ruleId/ramp-schedule` to modify coverage/targeting.

```
growthbook ramp-schedules update-ramp-schedule-steps [flags]
```

### Examples

```
  growthbook ramp-schedules update-ramp-schedule-steps --id <id> --steps '[]'
```

### Options

```
      --body string    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help           help for update-ramp-schedule-steps
  -i, --id string      [required]
  -s, --steps string   Full replacement of the steps array. Step-level coverage patches (actions) are intentionally excluded — those require a revision publish because they change the SDK payload. Use the revision flow to modify coverage/targeting; use this endpoint to update monitoring flags and hold conditions. [required]
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
