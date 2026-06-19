## growthbook metrics put

Update a metric

### Synopsis

Update a metric

```
growthbook metrics put [flags]
```

### Examples

```
  growthbook metrics put --id <id>
```

### Options

```
  -a, --archived               boolean flag
  -b, --behavior string        JSON object
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string     Description of the metric
  -h, --help                   help for put
  -i, --id string              The id of the requested resource [required]
      --managed-by string      Where this metric must be managed from. If not set (empty string), it can be managed from anywhere. If set to "api", it can be managed via the API only. Please note that we have deprecated support for setting the managedBy property to "admin". Your existing Legacy Metrics with this value will continue to work, but we suggest migrating to Fact Metrics instead. (options: , api, admin)
      --mixpanel sql           Only use for MixPanel (non-SQL) Data Sources. Only one of sql, `sqlBuilder` or `mixpanel` allowed.
  -n, --name string            Name of the metric
      --owner string           The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --projects stringArray   List of project IDs for projects that can access this metric
      --sql sql                Preferred way to define SQL. Only one of sql, `sqlBuilder` or `mixpanel` allowed.
      --sql-builder sql        An alternative way to specify a SQL metric, rather than a full query. Using sql is preferred to `sqlBuilder`. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed
      --tags stringArray       List of tags
      --type string            Type of metric. See [Metrics documentation](/app/metrics/legacy) (options: binomial, count, duration, revenue)
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth                   If using Bearer auth, pass the Secret Key as the token:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                                      ```
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
      --password                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       username
```

### SEE ALSO

* [growthbook metrics](growthbook_metrics.md)	 - Metrics used as goals and guardrails for experiments
