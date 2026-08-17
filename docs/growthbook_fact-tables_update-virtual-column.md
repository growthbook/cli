## growthbook fact-tables update-virtual-column

Update a virtual (computed) column on a fact table

### Synopsis

Update a virtual (computed) column on a fact table

```
growthbook fact-tables update-virtual-column [flags]
```

### Examples

```
  growthbook fact-tables update-virtual-column --fact-table-id <id> --id <id>
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --datatype string        The data type of the computed column (options: number, string, date, boolean, json, binary, other)
      --description string     string value
  -f, --fact-table-id string   Specify a specific fact table [required]
  -h, --help                   help for update-virtual-column
  -i, --id string              The id of the requested resource [required]
      --name string            Display name for the column
      --number-format string   options: , currency, time:seconds, memory:bytes, memory:kilobytes
  -s, --sql string             The SQL expression that computes the column value
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

* [growthbook fact-tables](growthbook_fact-tables.md)	 - Fact Tables describe the shape of your data warehouse tables
