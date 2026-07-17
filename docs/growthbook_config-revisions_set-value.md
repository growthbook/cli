## growthbook config-revisions set-value

Update the value of a config draft revision

### Synopsis

Stages a new `value` (this config's own JSON object) on the draft. Pass `version: "new"` to auto-create a draft. A `@config:` inheritance entry in the value is rejected — express lineage via the `parent`/`extends` metadata fields instead. Configs are environment-agnostic: there is no per-environment override (use a Constant for that).

Inheritance is a deep (targeted) patch: this value is merged onto the resolved parent recursively, key by key — restate only the leaves you want to change and the rest are inherited. Arrays and scalars replace wholesale, `null` is a value (it does not delete a key), and a value composed from a constant via `$extends` is applied whole.

Set `inferSchemaIfMissing: true` to derive and stage a field schema from the value when the config has none yet.

```
growthbook config-revisions set-value [flags]
```

### Examples

```
  growthbook config-revisions set-value --key <key> --version-param <value>
```

### Options

```
      --body string                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                            help for set-value
      --ignore-warnings string          Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
      --infer-schema-if-missing         When the config has no schema yet, infer one from the supplied value and stage it on the same draft.
  -k, --key string                      [required]
      --revision-comment string         string value
      --revision-title string           string value
  -s, --skip-schema-validation string   Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
      --value string                    This config's own value as a JSON object — a targeted patch deep-merged onto the resolved parent value.
      --version-param string            [required]
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

* [growthbook config-revisions](growthbook_config-revisions.md)	 - Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)
