## growthbook configs archive

Archive a single config

### Synopsis

Archives a config. A child config (including an environment/project override) is archived outright when its live value is an empty patch or nothing serves it. When archiving would strip a value that live features or other configs still consume, the request returns a 422 listing the blocking gates — re-submit with `"ignoreWarnings": true` in the request body to acknowledge and proceed. A locked config, or one whose org requires approval, returns its own gate (unlock or route the change through a draft revision).

```
growthbook configs archive [flags]
```

### Examples

```
  growthbook configs archive --key <key>
```

### Options

```
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                               help for archive
  -i, --ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
  -k, --key string                         The key of the config [required]
      --skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
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

* [growthbook configs](growthbook_configs.md)	 - **Beta** — these endpoints are new and may change in backwards-incompatible ways
