## growthbook config-revisions set-schema

Update or import the schema of a config draft revision

### Synopsis

Stages this config's field schema on the draft. Provide exactly ONE source:
- `schema`: a schema document — `{ type: "json-schema", value }` (a JSON Schema object) or `{ type: "typescript", value }` (TypeScript source). **JSON Schema is the recommended ("happy path") format** — it is the canonical pivot, preserves nested objects/arrays, and resolves local `$ref`/`$defs` (so generator output with referenced types works). `typescript` is a best-effort convenience parser. All conversions are lossy-by-design and degrade exotic constructs to permissive types WITH warnings (returned in `warnings`).
- `infer: true`: derive the schema from the draft's value.

Fields whose key a published ancestor already owns follow "base wins": an identical re-declaration is stripped with a `redundant-declaration` warning; one with a differing definition is rejected. Pass `version: "new"` to auto-create a draft.

```
growthbook config-revisions set-schema [flags]
```

### Examples

```
  growthbook config-revisions set-schema --key <key> --version-param <value>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.additional-properties              Whether the resulting object schema permits extra keys (family extensibility).
      --body-param.ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --body-param.infer                              Derive the schema from the draft's value instead.
      --body-param.revision-comment string            string value
      --body-param.revision-title string              string value
      --body-param.schema string                      JSON value (variants: json-schema: { value: object }, typescript: { value: string }, protobuf: { value: string }, python: { value: string }, go: { value: string }, rust: { value: string })
      --body-param.schema.go string                   ConfigSchemaSource_Go variant as JSON
      --body-param.schema.go.value struct             Go source — a struct definition. [required]
      --body-param.schema.json-schema string          ConfigSchemaSource_JSONSchema variant as JSON
      --body-param.schema.protobuf string             ConfigSchemaSource_Protobuf variant as JSON
      --body-param.schema.protobuf.value message      Protobuf (proto3) source — a message definition. [required]
      --body-param.schema.python string               ConfigSchemaSource_Python variant as JSON
      --body-param.schema.python.value BaseModel      Python source — a Pydantic BaseModel class. [required]
      --body-param.schema.rust string                 ConfigSchemaSource_Rust variant as JSON
      --body-param.schema.rust.value struct           Rust source — a serde struct definition. [required]
      --body-param.schema.typescript string           ConfigSchemaSource_Typescript variant as JSON
      --body-param.schema.typescript.value string     TypeScript source — an interface or object type. [required]
      --body-param.skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
  -h, --help                                          help for set-schema
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
