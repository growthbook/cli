## growthbook ramp-schedules api-advance

Advance to the next step, overriding any holds

### Synopsis

Moves the schedule to the next step, bypassing **all** hold conditions —
interval, min sample size, and monitoring signal holds. Accepts `running`
or `paused` status; if paused, the schedule is implicitly resumed (timing
anchors recalculated) before the step moves.

**Approval gate**: if the current step has an unsatisfied
`holdConditions.requiresApproval` gate, this endpoint returns **409** by
default. Either call `/actions/approve-step` first (recommended), or pass
`force: true` to override the approval gate. `force: true` requires
`canBypassApprovalChecks` permission and is logged in the audit trail.

**Two common uses:**
- **Post-interval monitoring hold** (`decision: "hold"`, interval elapsed): the



  step timer has completed but a signal or guardrail is flagging concern. Use
  this after reviewing the `/status` health summary and deciding to accept the
  risk and proceed.
- **Hard override**: skip a step regardless of where it is in its interval or



  hold conditions (CI gate, external deployment pipeline).

When to use other actions instead:
- **`/actions/resume`** — restores a paused schedule without moving the step.
- **`/actions/approve-step`** — clears only the approval gate; other conditions



  still resolve naturally.
- **`/actions/rollback`** — preferred response when `decision: "rollback"` or



  signals include `guardrail-failing`.

```
growthbook ramp-schedules api-advance [flags]
```

### Examples

```
  growthbook ramp-schedules api-advance --id <id>
```

### Options

```
      --body string                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --force canBypassApprovalChecks   Bypass a pending approval gate on the current step. Requires admin-level (canBypassApprovalChecks) permission. When omitted or `false`, a 409 is returned if the step has an unsatisfied `holdConditions.requiresApproval` gate.
  -h, --help                            help for api-advance
  -i, --id string                       [required]
  -r, --reason string                   Reason for advancing
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook ramp-schedules](growthbook_ramp-schedules.md)	 - Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring
