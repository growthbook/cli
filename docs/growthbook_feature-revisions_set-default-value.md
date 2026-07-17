## growthbook feature-revisions set-default-value

Set the default value in a draft revision

### Synopsis

Set the default value in a draft revision

```
growthbook feature-revisions set-default-value [flags]
```

### Examples

```
  growthbook feature-revisions set-default-value --id <id> --version-param <value> --default-value <value>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --default-value baseConfig                      New default value. In Config mode (feature has baseConfig), the default must be exactly a config with no overrides: send `"{}"` to use `baseConfig`, or set `defaultValueConfig` to point at a descendant. [required]
      --default-value-config baseConfig               Key of a config within the feature's baseConfig family that the default value resolves to (the base itself or a descendant). The default is exactly that config with no overrides; pass `null` to use `baseConfig`. Do not embed `@config:` in `defaultValue` — use this field.
  -h, --help                                          help for set-default-value
      --id string                                     [required]
      --ignore-warnings blockPublishOnSchemaError     Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
      --revision-comment string                       Comment for a newly created draft. Only used when version is "new"; ignored for existing revisions.
      --revision-title string                         Title for a newly created draft. Only used when version is "new"; ignored for existing revisions.
  -s, --skip-schema-validation bypassApprovalChecks   Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
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

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
