## growthbook data-sources preview-warehouse-column-values

Preview distinct column values

### Synopsis

Preview distinct column values

```
growthbook data-sources preview-warehouse-column-values [flags]
```

### Examples

```
  growthbook data-sources preview-warehouse-column-values --id <id> --database-name <value> --table-schema <value> --table-name <value> --columns '["<value 1>","<value 2>","<value 3>"]'
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --columns stringArray    [required]
      --database-name string   Database or project name. [required]
  -h, --help                   help for preview-warehouse-column-values
  -i, --id string              The id of the data source [required]
  -l, --limit int              Max rows to return. Defaults to 20.
      --table-name string      Table name (unqualified). [required]
      --table-schema string    Schema or dataset name. [required]
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

* [growthbook data-sources](growthbook_data-sources.md)	 - How GrowthBook connects and queries your data, including cached database schema metadata (information schemas) for tables and columns
