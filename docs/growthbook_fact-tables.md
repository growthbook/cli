## growthbook fact-tables

Fact Tables describe the shape of your data warehouse tables

### Synopsis

Fact Tables describe the shape of your data warehouse tables

```
growthbook fact-tables [flags]
```

### Options

```
  -h, --help   help for fact-tables
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook fact-tables create](growthbook_fact-tables_create.md)	 - Create a single fact table
* [growthbook fact-tables delete](growthbook_fact-tables_delete.md)	 - Deletes a single fact table
* [growthbook fact-tables delete-fact-table-filter](growthbook_fact-tables_delete-fact-table-filter.md)	 - Deletes a single fact table filter
* [growthbook fact-tables get](growthbook_fact-tables_get.md)	 - Get a single fact table
* [growthbook fact-tables get-aggregated](growthbook_fact-tables_get-aggregated.md)	 - Get the materialization status of a fact table's shared daily aggregated tables
* [growthbook fact-tables get-aggregated-table-run](growthbook_fact-tables_get-aggregated-table-run.md)	 - Get a single aggregated table run
* [growthbook fact-tables get-fact-table-filter](growthbook_fact-tables_get-fact-table-filter.md)	 - Get a single fact filter
* [growthbook fact-tables list](growthbook_fact-tables_list.md)	 - Get all fact tables
* [growthbook fact-tables list-aggregated-table-runs](growthbook_fact-tables_list-aggregated-table-runs.md)	 - List aggregated table runs
* [growthbook fact-tables list-fact-table-filters](growthbook_fact-tables_list-fact-table-filters.md)	 - Get all filters for a fact table
* [growthbook fact-tables post-bulk-import-facts](growthbook_fact-tables_post-bulk-import-facts.md)	 - Bulk import fact tables, filters, and metrics
* [growthbook fact-tables post-fact-table-filter](growthbook_fact-tables_post-fact-table-filter.md)	 - Create a single fact table filter
* [growthbook fact-tables refresh-aggregated](growthbook_fact-tables_refresh-aggregated.md)	 - Force a refresh or full restate of a fact table's shared daily aggregated tables
* [growthbook fact-tables update](growthbook_fact-tables_update.md)	 - Update a single fact table
* [growthbook fact-tables update-fact-table-filter](growthbook_fact-tables_update-fact-table-filter.md)	 - Update a single fact table filter
