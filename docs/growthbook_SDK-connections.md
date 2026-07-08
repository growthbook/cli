## growthbook SDK-connections

Client keys and settings for connecting SDKs to a GrowthBook instance

### Synopsis

Client keys and settings for connecting SDKs to a GrowthBook instance

```
growthbook SDK-connections [flags]
```

### Options

```
  -h, --help   help for SDK-connections
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook SDK-connections delete](growthbook_SDK-connections_delete.md)	 - Deletes a single SDK connection
* [growthbook SDK-connections get](growthbook_SDK-connections_get.md)	 - Get a single sdk connection
* [growthbook SDK-connections list](growthbook_SDK-connections_list.md)	 - Get all sdk connections
* [growthbook SDK-connections lookup-SDK-connection-by-key](growthbook_SDK-connections_lookup-SDK-connection-by-key.md)	 - Find a single sdk connection by its key
* [growthbook SDK-connections post](growthbook_SDK-connections_post.md)	 - Create a single sdk connection
* [growthbook SDK-connections put](growthbook_SDK-connections_put.md)	 - Update a single sdk connection
