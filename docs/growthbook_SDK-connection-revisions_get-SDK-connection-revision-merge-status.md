## growthbook SDK-connection-revisions get-SDK-connection-revision-merge-status

Get merge status for a draft revision

### Synopsis

Dry-runs a merge of the draft into the live SDK Connection and returns any conflicts. A draft whose base is out of date can be brought up to date with `rebase`.

```
growthbook SDK-connection-revisions get-SDK-connection-revision-merge-status [flags]
```

### Examples

```
  growthbook SDK-connection-revisions get-SDK-connection-revision-merge-status --id <id> --version-param <value>
```

### Options

```
  -h, --help                   help for get-SDK-connection-revision-merge-status
  -i, --id string              The SDK Connection id [required]
  -v, --version-param string   [required]
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

* [growthbook SDK-connection-revisions](growthbook_SDK-connection-revisions.md)	 - Draft revisions for SDK Connections: staged changes to a connection's settings and webhooks, their review, and their lifecycle (publish, discard, rebase)
