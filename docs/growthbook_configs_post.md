## growthbook configs post

Create a single config

### Synopsis

Create a single config

```
growthbook configs post [flags]
```

### Examples

```
  growthbook configs post --key <key> --name <value>
```

### Options

```
      --body string                      Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string               string value
      --experiment-guard                 Enable the experiment guard on this config: publishing a change served to a running experiment soft-blocks unless overridden. Omit to inherit the org default.
      --extends stringArray              Additional composition bases (config keys) layered on top of parent, in precedence order (later overrides earlier; all override parent; own keys win last). Set inheritance here, never via a @config: entry in value.
      --extensible                       boolean flag
  -h, --help                             help for post
      --ignore-warnings string           Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
      --invariants string                Cross-field validation rules. Each rule's expression is a mongo condition (mongrule). Stored on the config schema and enforced at publish.
  -k, --key string                       Stable reference handle (lowercase slug, unique per org), referenced as @config:key [required]
  -n, --name string                      The display name of the config [required]
      --owner string                     The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. When omitted, it defaults to the user associated with the request's Personal Access Token (PAT), if one is being used.
      --parent string                    The key of the config to inherit from (the primary lineage spine). Express inheritance via parent/extends, NEVER via a @config: entry in value (which is rejected).
      --project string                   string value
      --schema string                    JSON value (variants: json-schema: { value: object }, typescript: { value: string }, protobuf: { value: string }, python: { value: string }, go: { value: string }, rust: { value: string })
      --schema.go string                 ConfigSchemaSource_Go variant as JSON
      --schema.go.value string           Go source — a struct definition. [required]
      --schema.json-schema string        ConfigSchemaSource_JSONSchema variant as JSON
      --schema.protobuf string           ConfigSchemaSource_Protobuf variant as JSON
      --schema.protobuf.value string     Protobuf (proto3) source — a message definition. [required]
      --schema.python string             ConfigSchemaSource_Python variant as JSON
      --schema.python.value string       Python source — a Pydantic BaseModel class. [required]
      --schema.rust string               ConfigSchemaSource_Rust variant as JSON
      --schema.rust.value string         Rust source — a serde struct definition. [required]
      --schema.typescript string         ConfigSchemaSource_Typescript variant as JSON
      --schema.typescript.value string   TypeScript source — an interface or object type. [required]
      --scoped-overrides string          Ordered, first-match-wins environment/project-scoped variant selection. Each entry points at a flavor config (a child config, by key) whose value is deep-merged onto this config's resolved value when the (environment, project) scope matches — resolved at build time, per layer. This is how you create an environment-scoped override (as opposed to a plain child config): make a child config for the override value, then add it here with its scope. Send the complete list to replace it; an empty array clears all overrides. Entries must reference existing configs, may not reference this config itself, and may not be unreachable (fully subsumed by an earlier entry).
      --skip-schema-validation string    Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
      --source string                    Optional identifier of the consuming codebase/service. When a typed-code schema (typescript/protobuf/python/go/rust) is supplied, its named-type structure is captured under this source so GET /configs/:key/schema?source=<id>&format=<lang> can reproduce those names.
  -v, --value string                     This config's base value as a JSON object. Per-environment/project variants are expressed via scopedOverrides.
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

* [growthbook configs](growthbook_configs.md)	 - Reusable, typed, inheritable JSON objects referenced from feature flag values as `@config:key`
