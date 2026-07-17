## growthbook configs get-schema

Export a config's schema

### Synopsis

Export a config's schema

```
growthbook configs get-schema [flags]
```

### Examples

```
  growthbook configs get-schema --key <key>
```

### Options

```
  -e, --effective string      When true, includes fields inherited across the lineage (the family's accumulated schema). When false (default), returns only this config's own fields.
  -f, --format-param string   Output format. json-schema (default) returns a JSON Schema document; typescript, protobuf, python (Pydantic), go, and rust (serde) render the schema as source in that language. (options: json-schema, typescript, protobuf, python, go, rust)
  -h, --help                  help for get-schema
  -k, --key string            The key of the config [required]
  -s, --source string         Render using a previously-captured source projection (its named types). Only affects the typed-code formats (typescript/protobuf/python/go/rust); ignored if the source has no projection.
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
