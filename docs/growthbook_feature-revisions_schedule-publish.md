## growthbook feature-revisions schedule-publish

Schedule (or cancel) a deferred publish for a draft revision

### Synopsis

Schedules the draft to publish on or after `scheduledPublishAt`. When approval is required, publishing waits until the draft is also approved. Send `scheduledPublishAt: null` to cancel the schedule.

Set `lockEdits` to prevent content changes while the schedule is pending; rebasing remains allowed. Set `lockOthers` to prevent other drafts of this Feature Flag from being published until this schedule runs or is canceled. The caller needs Publish access, and that access is checked again when the schedule runs. A caller with Bypass draft approvals access can schedule an unapproved draft by sending `bypassApproval: true`. That schedule must be canceled and recreated before it can be changed.

```
growthbook feature-revisions schedule-publish [flags]
```

### Examples

```
  growthbook feature-revisions schedule-publish --id <id> --version-param <value> --scheduled-publish-at 2025-10-13T11:16:06.111Z
```

### Options

```
      --body string                                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --bypass-approval                             boolean flag
  -h, --help                                        help for schedule-publish
      --id string                                   [required]
      --ignore-warnings                             Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --lock-edits                                  boolean flag
      --lock-others                                 boolean flag
      --scheduled-publish-at 2026-01-31T09:00:00Z   When to publish, as an RFC3339 timestamp (e.g. 2026-01-31T09:00:00Z or `2026-01-31T02:00:00-07:00`), or `null` to cancel a pending schedule. [required]
      --skip-hooks skipSchemaValidation             Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --skip-schema-validation skipHooks            Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
  -v, --version-param string                        [required]
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
