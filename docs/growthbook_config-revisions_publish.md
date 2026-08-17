## growthbook config-revisions publish

Publish a draft revision

### Synopsis

Publishes the draft and makes its changes live. The caller needs Publish access for the affected environments. When approval is required, the draft must be approved unless the caller has Bypass draft approvals access. If the organization requires rebasing, an out-of-date draft must be rebased first; an authorized caller can instead send `ignoreWarnings: true` to force-publish it. A locked Config cannot be published. A 422 response lists every blocking gate and the available resolution. Publishing a schema change also normalizes descendant Configs according to their inheritance rules.

```
growthbook config-revisions publish [flags]
```

### Examples

```
  growthbook config-revisions publish --key <key> --version-param <value>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.bypass-approval                    Deprecated and ignored. Approval is bypassed automatically when the caller has Bypass draft approvals access for this resource or when the organization enables the REST API approval bypass. Otherwise, the revision must be approved before it can be published.
      --body-param.ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --body-param.skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
  -h, --help                                          help for publish
  -i, --ignore-warnings ignoreWarnings                Deprecated — pass ignoreWarnings in the request body instead.
  -k, --key string                                    [required]
  -s, --skip-schema-validation skipSchemaValidation   Deprecated — pass skipSchemaValidation in the request body instead.
  -v, --version-param string                          [required]
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

* [growthbook config-revisions](growthbook_config-revisions.md)	 - **Beta** — these endpoints are new and may change in backwards-incompatible ways
