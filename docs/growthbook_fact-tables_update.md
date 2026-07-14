## growthbook fact-tables update

Update a single fact table

### Synopsis

Update a single fact table

```
growthbook fact-tables update [flags]
```

### Examples

```
  growthbook fact-tables update --id <id>
```

### Options

```
      --aggregated-fact-table-settings string   Settings for maintaining shared daily aggregated tables (a subset of userIdTypes plus the daily update time and restate lookback window) used to speed up CUPED. Requires the data pipeline (pipeline-mode) feature.
      --archived                                boolean flag
      --body string                             Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --columns string                          Optional array of columns that you want to update. Only allows updating properties of existing columns. Cannot create new columns or delete existing ones. Columns cannot be added or deleted; column structure is determined by SQL parsing. Slice-related properties require an enterprise license.
      --columns-error string                    Error message if there was an issue parsing the SQL schema
      --description string                      Description of the fact table
  -e, --event-name string                       The event name used in SQL template variables
  -h, --help                                    help for update
  -i, --id string                               The id of the requested resource [required]
  -m, --managed-by string                       Set this to "api" to disable editing in the GrowthBook UI (options: , api, admin)
  -n, --name string                             string value
      --owner string                            The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --projects stringArray                    List of associated project ids
  -s, --sql string                              The SQL query for this fact table
  -t, --tags stringArray                        List of associated tags
  -u, --user-id-types stringArray               List of identifier columns in this table. For example, "id" or "anonymous_id"
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
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
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
