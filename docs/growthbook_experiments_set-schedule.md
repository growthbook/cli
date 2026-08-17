## growthbook experiments set-schedule

Set an experiment's schedule and shipping automation

### Synopsis

Full-replace of the experiment's scheduled start/end and end-of-experiment shipping automation. The body is the complete desired state: any omitted field is cleared (omit `startAt` to remove a scheduled start; send an empty body to clear the whole schedule). Provide either `stopAt` or `stopAfter`, not both; a relative `stopAfter` resolves to a concrete stop when the experiment starts. Setting a scheduled end (`stopAt` or `stopAfter`) requires a `scheduledStopPlan` (any mode, including `notify`). The scheduled end must be in the future: a `stopAt` (or a `stopAfter` that resolves) in the past is rejected — including one whose end date has already passed and been acted on, so changing the plan after a soft end requires committing to a new end date. Auto-ship shipping requires the Decision Framework.

```
growthbook experiments set-schedule [flags]
```

### Examples

```
  growthbook experiments set-schedule --id <id>
```

### Options

```
      --body string                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                         help for set-schedule
  -i, --id string                    The id of the requested resource [required]
      --scheduled-stop-plan notify   What happens at the scheduled end date. notify keeps the experiment running and just notifies (soft). `auto-ship` (requires the Decision Framework) ships the winning variation and stops; multi-winner ties break on `tiebreakerMetricId` (higher lift); with no clear winner, `fallback` either keeps running (`notify`) or ships `fallbackVariationId`. `force-ship` stops and rolls out `fallbackVariationId`. `stop` is a hard deadline that stops with no rollout. For `force-ship` and `stop`, the Decision Framework verdict (won/lost/inconclusive) is recorded as metadata when available.
      --start-at string              ISO datetime when the experiment should start; must be in the future. Omit to clear a scheduled start. Staging still happens via POST /experiments/{id}/start.
      --stop-after string            Deferred relative end, resolved to a concrete stop at the experiment's actual start (or now, if already running).
      --stop-at string               Absolute ISO datetime to stop the experiment.
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
