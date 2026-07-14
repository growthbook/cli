## growthbook feature-revisions put-feature-revision-rule

Update a rule in a draft revision

### Synopsis

Patches fields on an existing rule (identified by `ruleId`). The rule `type` cannot be changed. Scope can be updated via `allEnvironments` / `environments` patch fields.

**Scheduling:** For `force` and `rollout` rules, update the schedule via `rampSchedule` (multi-step ramp) or `schedule` (simple start/end window) — these manage standalone ramp actions and set `pendingRamp: "create"` on the rule. For `experiment-ref` and `safe-rollout` rules, only `schedule` is supported and updates legacy schedule fields on the rule itself (`rampSchedule` is not available for these rule types).

```
growthbook feature-revisions put-feature-revision-rule [flags]
```

### Examples

```
  growthbook feature-revisions put-feature-revision-rule --id <id> --version-param <value> --rule-id <id> --rule '{}'
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                      help for put-feature-revision-rule
  -i, --id string                 [required]
      --ramp-schedule schedule    Multi-step ramp schedule for force/rollout rules. Not supported for experiment-ref or safe-rollout rules. Mutually exclusive with schedule.
      --revision-comment string   Comment for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --revision-title string     Title for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --rule string               [required]
      --rule-id string            [required]
  -s, --schedule rampSchedule     Simple start/end date window. For force/rollout rules this manages a standalone ramp action; for experiment-ref/safe-rollout rules this updates legacy schedule fields on the rule. Mutually exclusive with rampSchedule.
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

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
