## growthbook configs update

Partially update a single config

### Synopsis

Applies the change immediately and records it as a published revision, so it appears in history and fires revision webhooks. When the organization requires approvals, open a draft instead or pass `bypassApproval` with the bypass permission.

```
growthbook configs update [flags]
```

### Examples

```
  growthbook configs update --key <key>
```

### Options

```
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --bypass-approval                    Set to true to write directly to the live Config without approval. The caller must have Bypass draft approvals access in the Config's Project, unless the organization enables the REST API approval bypass. This field has no effect when approval is not required.
      --description string                 string value
      --experiment-guard                   Enable or disable the experiment guard for this Config. Disabling it requires Bypass draft approvals access.
      --extends parent                     Replace the composition bases (mixins) layered on top of parent, in precedence order (later overrides earlier; all override `parent`; own keys win last). Send the complete set; an empty array clears all mixins. Set inheritance here, never via a `@config:` entry in `value`.
      --extensible                         boolean flag
  -h, --help                               help for update
      --ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --invariants string                  Replace the config's cross-field validation rules. Each rule's expression is a mongo condition (mongrule). Send the complete set; an empty array clears all rules. Omit to leave them unchanged.
  -k, --key string                         The key of the config [required]
  -n, --name string                        string value
      --owner string                       The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --parent key                         Change the lineage parent (the key of the config to inherit from). Set to an empty string to detach from the parent and make this a root config.
      --project string                     string value
      --schema string                      JSON value (variants: json-schema: { value: object }, typescript: { value: string }, protobuf: { value: string }, python: { value: string }, go: { value: string }, rust: { value: string })
      --schema.go string                   ConfigSchemaSource_Go variant as JSON
      --schema.go.value struct             Go source — a struct definition. [required]
      --schema.json-schema string          ConfigSchemaSource_JSONSchema variant as JSON
      --schema.protobuf string             ConfigSchemaSource_Protobuf variant as JSON
      --schema.protobuf.value message      Protobuf (proto3) source — a message definition. [required]
      --schema.python string               ConfigSchemaSource_Python variant as JSON
      --schema.python.value BaseModel      Python source — a Pydantic BaseModel class. [required]
      --schema.rust string                 ConfigSchemaSource_Rust variant as JSON
      --schema.rust.value struct           Rust source — a serde struct definition. [required]
      --schema.typescript string           ConfigSchemaSource_Typescript variant as JSON
      --schema.typescript.value string     TypeScript source — an interface or object type. [required]
      --scoped-overrides key               Replace the ordered, first-match-wins environment/project-scoped variant selection. Each entry points at a flavor config (a child config, by key) whose value is deep-merged onto this config's resolved value when the (environment, project) scope matches. Send the complete list; an empty array clears all overrides; omit to leave unchanged. Entries must reference existing configs, may not reference this config itself, and may not be unreachable.
      --skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --source typescript                  Optional identifier of the consuming codebase/service. When a typescript or `protobuf` schema is supplied, its named-type structure is captured under this source for reproduction on export.
  -v, --value scopedOverrides              This config's base value as a JSON object. Per-environment/project variants are expressed via scopedOverrides.
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
