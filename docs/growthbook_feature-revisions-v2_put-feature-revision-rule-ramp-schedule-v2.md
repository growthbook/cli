## growthbook feature-revisions-v2 put-feature-revision-rule-ramp-schedule-v2

Set ramp schedule for a rule

### Synopsis

Queues a revision-controlled ramp action for this rule. If the rule already has a live ramp schedule, this stores an `update` action applied on publish; otherwise it stores a `create` action. No live schedule config changes are applied immediately by this endpoint.

```
growthbook feature-revisions-v2 put-feature-revision-rule-ramp-schedule-v2 [flags]
```

### Examples

```
  growthbook feature-revisions-v2 put-feature-revision-rule-ramp-schedule-v2 --id <id> --version-param <value> --rule-id <id>
```

### Options

```
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --cutoff-date string         ISO 8601 date-time, e.g. "2025-07-01T00:00:00Z". The ramp ends at this time.
      --end-actions string         list of values
      --environment string         string value
  -h, --help                       help for put-feature-revision-rule-ramp-schedule-v2
  -i, --id string                  [required]
  -l, --lockdown-config string     JSON object
  -m, --monitoring-config string   JSON object
  -n, --name string                string value
      --revision-comment string    Comment for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --revision-title string      Title for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --rule-id string             [required]
      --start-actions string       list of values
      --start-date string          ISO 8601 date-time, e.g. "2025-06-01T00:00:00Z". Absent or null means start immediately on publish.
      --steps string               list of values
  -t, --template-id string         string value
  -v, --version-param string       [required]
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

* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
