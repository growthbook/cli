## growthbook ramp-schedules refresh-monitoring

Trigger a manual monitoring update

### Synopsis

Queues a new analysis snapshot for the schedule's monitoring experiment.
The snapshot runs asynchronously — poll `GET /ramp-schedules/:id/status`
until `snapshotAt` advances to confirm results are ready.

Only available when the schedule is within its monitored step window:
- Not in a terminal state (`completed` or `rolled-back`).
- Has at least one step with `monitored: true`.
- `currentStepIndex` is within `[firstMonitoredStepIndex, lastMonitoredStepIndex]`.

Violating any condition returns **409 Conflict** with a descriptive message.

Requires the `runQueries` permission on the configured datasource (enforced via `canRunExperimentQueries`).

```
growthbook ramp-schedules refresh-monitoring [flags]
```

### Examples

```
  growthbook ramp-schedules refresh-monitoring --id <id>
```

### Options

```
  -h, --help        help for refresh-monitoring
  -i, --id string   [required]
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
