## growthbook ramp-schedules eject-target

Remove a target rule from a ramp schedule

### Synopsis

Detaches a target rule from this ramp schedule. Identify the target either
by its `targetId` or by the `[ruleId, environment]` pair.

If this is the last target on the schedule, the schedule is deleted entirely
and the response contains `deleted: true` instead of `rampSchedule`.

```
growthbook ramp-schedules eject-target [flags]
```

### Examples

```
  growthbook ramp-schedules eject-target --id <id>
```

### Options

```
      --body string          Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --environment string   Deprecated pre-v2 disambiguator. Optional when used with ruleId; omit on v2 ramps.
  -h, --help                 help for eject-target
  -i, --id string            [required]
  -r, --rule-id string       Rule ID — use as an alternative to targetId
  -t, --target-id string     Target ID (from the targets array)
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
