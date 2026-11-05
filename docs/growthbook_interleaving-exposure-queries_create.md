## growthbook interleaving-exposure-queries create

Create a single interleavingExposureQuery

### Synopsis

Create a single interleavingExposureQuery

```
growthbook interleaving-exposure-queries create [flags]
```

### Examples

```
  growthbook interleaving-exposure-queries create --datasource-id <id> --name <value> --user-id-types '[]' --query <value>
```

### Options

```
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --datasource-id string        [required]
      --description string          string value
  -h, --help                        help for create
  -n, --name string                 [required]
      --owner string                The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. When omitted, it defaults to the requesting member: a personal access token's owner, or the member an organization API key assumes. An organization API key acting as itself leaves it empty.
      --query string                [required]
  -u, --user-id-types stringArray   [required]
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

* [growthbook interleaving-exposure-queries](growthbook_interleaving-exposure-queries.md)	 - Operations for interleaving-exposure-queries
