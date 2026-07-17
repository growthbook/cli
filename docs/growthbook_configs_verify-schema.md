## growthbook configs verify-schema

Verify a config's schema against a source (drift check)

### Synopsis

Verify a config's schema against a source (drift check)

```
growthbook configs verify-schema [flags]
```

### Examples

```
  growthbook configs verify-schema --key <key> --schema '{"type":"typescript","value":"<value>"}'
```

### Options

```
      --body string                      Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                             help for verify-schema
  -k, --key string                       The key of the config [required]
  -s, --schema string                    JSON value (variants: json-schema: { value: object }, typescript: { value: string }, protobuf: { value: string }, python: { value: string }, go: { value: string }, rust: { value: string })
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
