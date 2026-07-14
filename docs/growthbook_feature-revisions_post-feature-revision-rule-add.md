## growthbook feature-revisions post-feature-revision-rule-add

Add a rule to a draft revision

### Synopsis

Appends a new rule to the revision's rule list. Supply `allEnvironments: true` on the rule to target all environments, or `environments: [...]` to scope to specific ones.

**Scheduling:** For `force` and `rollout` rules, attach a schedule via `rampSchedule` (multi-step ramp) or `schedule` (simple start/end window) — these create standalone ramp actions and set `pendingRamp: "create"` on the rule. For `experiment-ref` and `safe-rollout` rules, only `schedule` is supported and is stored as legacy schedule fields on the rule itself (`rampSchedule` is not available for these rule types).

```
growthbook feature-revisions post-feature-revision-rule-add [flags]
```

### Examples

```
  growthbook feature-revisions post-feature-revision-rule-add --id <id> --version-param <value> --rule '{"value":"<value>"}'
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                      help for post-feature-revision-rule-add
  -i, --id string                 [required]
      --ramp-schedule string      Multi-step ramp schedule for force/rollout rules. Not supported for experiment-ref or safe-rollout rules. Mutually exclusive with schedule.
      --revision-comment string   Comment for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --revision-title string     Title for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --rule string               JSON value (one of: { description: string, enabled: boolean, condition: string, savedGroups: object[], ... })
  -s, --schedule string           Simple start/end date window. For force/rollout rules this creates a standalone ramp action; for experiment-ref/safe-rollout rules this sets legacy schedule fields on the rule. Mutually exclusive with rampSchedule.
  -v, --version-param string      [required]
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

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
