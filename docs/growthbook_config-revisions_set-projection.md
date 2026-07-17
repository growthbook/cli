## growthbook config-revisions set-projection

Set (or update) a config's per-source render projection on a draft

### Synopsis

Stages a per-source render projection on the draft, AND the schema it implies. Provide a named `schema` source (`{ type: "typescript" | "protobuf" | "python" | "go" | "rust" | "json-schema", value }`) for the consuming codebase identified by `source`: GrowthBook derives the config's canonical schema from it (so the change projects into the Config) and captures that source's named-type structure under `renderProjections[source]`. Both are staged on the draft and published through the normal flow. Pass `version: "new"` to auto-create a draft. Lossy conversions degrade with `warnings`.

```
growthbook config-revisions set-projection [flags]
```

### Examples

```
  growthbook config-revisions set-projection --key <key> --version-param <value> --source <value> --schema '{"type":"rust","value":"<value>"}'
```

### Options

```
  -a, --additional-properties                         Whether the resulting object schema permits extra keys (family extensibility).
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                                          help for set-projection
  -i, --ignore-warnings blockPublishOnSchemaError     Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
  -k, --key string                                    [required]
      --revision-comment string                       string value
      --revision-title string                         string value
      --schema string                                 JSON value (variants: json-schema: { value: object }, typescript: { value: string }, protobuf: { value: string }, python: { value: string }, go: { value: string }, rust: { value: string })
      --schema.go string                              ConfigSchemaSource_Go variant as JSON
      --schema.go.value struct                        Go source — a struct definition. [required]
      --schema.json-schema string                     ConfigSchemaSource_JSONSchema variant as JSON
      --schema.protobuf string                        ConfigSchemaSource_Protobuf variant as JSON
      --schema.protobuf.value message                 Protobuf (proto3) source — a message definition. [required]
      --schema.python string                          ConfigSchemaSource_Python variant as JSON
      --schema.python.value BaseModel                 Python source — a Pydantic BaseModel class. [required]
      --schema.rust string                            ConfigSchemaSource_Rust variant as JSON
      --schema.rust.value struct                      Rust source — a serde struct definition. [required]
      --schema.typescript string                      ConfigSchemaSource_Typescript variant as JSON
      --schema.typescript.value string                TypeScript source — an interface or object type. [required]
      --skip-schema-validation bypassApprovalChecks   Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
      --source string                                 Identifier of the consuming codebase/service this projection belongs to. [required]
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

* [growthbook config-revisions](growthbook_config-revisions.md)	 - Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)
