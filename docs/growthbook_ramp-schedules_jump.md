## growthbook ramp-schedules jump

Jump to a specific step

### Synopsis

Teleports the schedule to `targetStepIndex` (forward or backward) and leaves
it `paused`. Resets timing anchors so the destination step's interval starts
fresh when the schedule is next resumed or started.

Pass `-1` to return to the pre-start position without applying rollback rule
patches — useful for resetting a non-started schedule. For a full traffic
revert, use `/actions/rollback` instead.

Accepts any non-terminal schedule status.

```
growthbook ramp-schedules jump [flags]
```

### Examples

```
  growthbook ramp-schedules jump --id <id> --target-step-index 938354
```

### Options

```
      --body string             Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                    help for jump
  -i, --id string               [required]
  -t, --target-step-index int   Zero-based index of the step to jump to; -1 = pre-start [required]
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
