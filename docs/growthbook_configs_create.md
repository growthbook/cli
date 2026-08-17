## growthbook configs create

Create a single config

### Synopsis

Create a single config

```
growthbook configs create [flags]
```

### Examples

```
  growthbook configs create --body-param.key <key> --body-param.name <value>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.description string                 string value
      --body-param.experiment-guard                   Enable the experiment guard on this config: publishing a change served to a running experiment soft-blocks unless overridden. Omit to inherit the org default.
      --body-param.extends key                        Additional composition bases (config keys) layered on top of `parent`, in precedence order (later overrides earlier; all override `parent`; own keys win last). Set inheritance here, never via a `@config:` entry in `value`.
      --body-param.extensible                         boolean flag
      --body-param.ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --body-param.invariants string                  Cross-field validation rules. Each rule's expression is a mongo condition (mongrule). Stored on the config schema and enforced at publish.
      --body-param.key @config:key                    Stable reference handle (lowercase slug, unique per org), referenced as @config:key [required]
      --body-param.name string                        The display name of the config [required]
      --body-param.owner string                       The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. When omitted, it defaults to the user associated with the request's Personal Access Token (PAT), if one is being used.
      --body-param.parent key                         The key of the config to inherit from (the primary lineage spine). Express inheritance via `parent`/`extends`, NEVER via a `@config:` entry in `value` (which is rejected).
      --body-param.project string                     string value
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
      --body-param.scoped-overrides key               Ordered, first-match-wins environment/project-scoped variant selection. Each entry points at a flavor config (a child config, by key) whose value is deep-merged onto this config's resolved value when the (environment, project) scope matches — resolved at build time, per layer. This is how you create an environment-scoped override (as opposed to a plain child config): make a child config for the override value, then add it here with its scope. Send the complete list to replace it; an empty array clears all overrides. Entries must reference existing configs, may not reference this config itself, and may not be unreachable (fully subsumed by an earlier entry).
      --body-param.skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.source typescript                  Optional identifier of the consuming codebase/service. When a typed-code schema (typescript/`protobuf`/`python`/`go`/`rust`) is supplied, its named-type structure is captured under this source so `GET /configs/:key/schema?source=<id>&format=<lang>` can reproduce those names.
      --body-param.value scopedOverrides              This config's base value as a JSON object. Per-environment/project variants are expressed via scopedOverrides.
  -h, --help                                          help for create
  -i, --ignore-warnings ignoreWarnings                Deprecated — pass ignoreWarnings in the request body instead.
  -s, --skip-schema-validation skipSchemaValidation   Deprecated — pass skipSchemaValidation in the request body instead.
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
