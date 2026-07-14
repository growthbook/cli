## growthbook segments update

Update a single segment

### Synopsis

Update a single segment

```
growthbook segments update [flags]
```

### Examples

```
  growthbook segments update --id <id>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --datasource-id string     ID of the datasource this segment belongs to
      --description string       Description of the segment
      --fact-table-id string     ID of the fact table this segment belongs to. This is required for FACT segments.
      --filters stringArray      Optional array of fact table filter ids that can further define the Fact Table based Segment.
  -h, --help                     help for update
      --id string                The id of the requested resource [required]
      --identifier-type string   Type of identifier (user, anonymous, etc.)
  -m, --managed-by string        Where this Segment must be managed from. If not set (empty string), it can be managed from anywhere. (options: , api)
  -n, --name string              Name of the segment
      --owner string             The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --projects stringArray     List of project IDs for projects that can access this segment
      --query string             SQL query that defines the Segment. This is required for SQL segments.
  -t, --type string              GrowthBook supports two types of Segments, SQL and FACT. SQL segments are defined by a SQL query, and FACT segments are defined by a fact table and filters. (options: SQL, FACT)
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

* [growthbook segments](growthbook_segments.md)	 - Segments used during experiment analysis
