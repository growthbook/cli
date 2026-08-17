## growthbook constant-revisions schedule-publish

Schedule (or cancel) a deferred publish

### Synopsis

Arms a revision to publish automatically at a future time. Pass `scheduledPublishAt` as an RFC3339 timestamp in the future to arm, or `null` to cancel a pending schedule. Requires the `scheduled-revisions` commercial feature and publish permission on the Constant. A draft that still requires approval must request review first (or be armed with `bypassApproval` by a caller who can bypass).

```
growthbook constant-revisions schedule-publish [flags]
```

### Examples

```
  growthbook constant-revisions schedule-publish --key <key> --version-param <value> --scheduled-publish-at <value>
```

### Options

```
      --body string                                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --bypass-approval                             boolean flag
  -h, --help                                        help for schedule-publish
  -i, --ignore-warnings                             Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
  -k, --key string                                  [required]
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

* [growthbook constant-revisions](growthbook_constant-revisions.md)	 - **Beta** — these endpoints are new and may change in backwards-incompatible ways
