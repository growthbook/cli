## growthbook experiments start

Start/Stage an experiment

### Synopsis

Starts an experiment or stages it for a future start if a `statusUpdateSchedule` is set on the experiment. Both check the launch checklist that `GET /experiments/{id}/start-checklist` returns, as gates. An incomplete item refuses the start with `checklist_incomplete` unless `skipChecklist` is set; a linked Feature Flag blocker refuses it either way. The refusal lists each blocking gate in `gates`, with the flag or permission that gets past it and a route that resolves it. A linked Feature Flag draft that still needs approval, or a flag locked by a ramp, is bypassed automatically when the caller has Bypass draft approvals access for that flag or the organization enables the REST API approval bypass. A start reports what it got past in `bypassedGates` and the drafts it published in `publishedFeatureDrafts`; if a draft can't publish, for example because it needs a rebase, the start fails with `pending_draft_publish_failed` before any is published. A staged start carries the caller's Bypass draft approvals access to when it runs, for the drafts as they were when it was staged: if one has changed since, or the caller has lost that access, it needs approval like any other. The organization's REST API approval bypass doesn't carry to a staged start.

```
growthbook experiments start [flags]
```

### Examples

```
  growthbook experiments start --id <id>
```

### Options

```
      --body string       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --bypass-approval   Deprecated and ignored. Approval is bypassed automatically when the caller has Bypass draft approvals access for this resource or when the organization enables the REST API approval bypass. Otherwise, the revision must be approved before it can be published.
  -h, --help              help for start
      --id string         The id of the requested resource [required]
      --ignore-warnings   Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
  -s, --skip-checklist    If true, skips validating the experiment satisfies all pre-launch checklist items. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
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
