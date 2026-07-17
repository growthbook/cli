## growthbook ramp-schedules approve-step

Approve the pending approval gate

### Synopsis

Clears whichever approval gate is currently pending on the schedule:

- **Start gate** — a schedule created with `requiresStartApproval` sits in `ready` at step -1 with its rule disabled (zero traffic). Approving starts the ramp (or, if a future `startDate` is set, arms it to start on that date).
- **Step gate** — the `holdConditions.requiresApproval` gate on the current step of a `running` schedule.

For a step gate, approval is the **final** gate: it can only be granted once every other hold has cleared. This endpoint rejects the request (`400`) if the step is not yet ready — while the interval timer is still counting down, or (for monitored steps) before fresh analysis is available or while a guardrail/health signal is failing. Poll `/status` and call this once it reports awaiting approval.

**Non-monitored steps**: once the interval has elapsed, approving clears the last hold and advances immediately, chaining through subsequent instant steps.

**Monitored steps**: approving clears the last hold and the agenda advances on its next tick (re-checking analysis first).

Different from `/actions/advance`: `approve-step` works within the normal evaluation flow and refuses to skip ahead of the interval or any other unmet gate. Use `/actions/advance` to bypass all remaining holds.

Requires update + publish (start gate) or review (step gate) permissions for the associated feature.

```
growthbook ramp-schedules approve-step [flags]
```

### Examples

```
  growthbook ramp-schedules approve-step --id <id>
```

### Options

```
  -h, --help        help for approve-step
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
