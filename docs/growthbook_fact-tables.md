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
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook fact-tables delete](growthbook_fact-tables_delete.md)	 - Deletes a single fact table
* [growthbook fact-tables delete-fact-table-filter](growthbook_fact-tables_delete-fact-table-filter.md)	 - Deletes a single fact table filter
* [growthbook fact-tables get](growthbook_fact-tables_get.md)	 - Get a single fact table
* [growthbook fact-tables get-aggregated](growthbook_fact-tables_get-aggregated.md)	 - Get the materialization status of a fact table's shared daily aggregated tables
* [growthbook fact-tables get-fact-table-filter](growthbook_fact-tables_get-fact-table-filter.md)	 - Get a single fact filter
* [growthbook fact-tables list](growthbook_fact-tables_list.md)	 - Get all fact tables
* [growthbook fact-tables list-fact-table-filters](growthbook_fact-tables_list-fact-table-filters.md)	 - Get all filters for a fact table
* [growthbook fact-tables post](growthbook_fact-tables_post.md)	 - Create a single fact table
* [growthbook fact-tables post-bulk-import-facts](growthbook_fact-tables_post-bulk-import-facts.md)	 - Bulk import fact tables, filters, and metrics
* [growthbook fact-tables post-fact-table-filter](growthbook_fact-tables_post-fact-table-filter.md)	 - Create a single fact table filter
* [growthbook fact-tables refresh-aggregated](growthbook_fact-tables_refresh-aggregated.md)	 - Force a refresh or full restate of a fact table's shared daily aggregated tables
* [growthbook fact-tables update](growthbook_fact-tables_update.md)	 - Update a single fact table
* [growthbook fact-tables update-fact-table-filter](growthbook_fact-tables_update-fact-table-filter.md)	 - Update a single fact table filter
