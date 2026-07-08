## growthbook dimensions create

Create a single dimension

### Synopsis

Create a single dimension

```
growthbook dimensions create [flags]
```

### Examples

```
  growthbook dimensions create --name <value> --datasource-id <id> --identifier-type <value> --query <value>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --datasource-id string     ID of the datasource this dimension belongs to [required]
      --description string       Description of the dimension
  -h, --help                     help for create
  -i, --identifier-type string   Type of identifier (user, anonymous, etc.) [required]
  -m, --managed-by string        Where this dimension must be managed from. If not set (empty string), it can be managed from anywhere. (options: , api)
  -n, --name string              Name of the dimension [required]
      --owner string             The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --query string             SQL query or equivalent for the dimension [required]
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
      --password string               HTTP Basic auth: your Secret Key as the username, with an empty password. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: your Secret Key as the username, with an empty password. username
```

### SEE ALSO

* [growthbook dimensions](growthbook_dimensions.md)	 - Dimensions used during experiment analysis
