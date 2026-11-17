## growthbook SDK-webhooks post-SDK-connection-webhook

Add a webhook to an SDK Connection

### Synopsis

Called whenever the connection's payload changes. `httpMethod` defaults to POST. More than one webhook per organization requires a paid plan.

```
growthbook SDK-webhooks post-SDK-connection-webhook [flags]
```

### Examples

```
  growthbook SDK-webhooks post-SDK-connection-webhook --id <id> --name <value> --endpoint https://brisk-taxicab.name/
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --endpoint string          [required]
      --headers string           JSON-encoded object, e.g. "{\"X-Token\": \"abc\"}"
  -h, --help                     help for post-SDK-connection-webhook
      --http-method string       options: GET, POST, PUT, DELETE, PATCH, PURGE
  -i, --id string                The SDK Connection id [required]
  -n, --name string              [required]
      --payload-format string    options: standard, standard-no-payload, sdkPayload, edgeConfig, edgeConfigUnescaped, none
      --payload-key string       string value
  -t, --timeout-seconds string   Request timeout in seconds, up to the server maximum (120 unless configured). Null uses the default.
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

* [growthbook SDK-webhooks](growthbook_SDK-webhooks.md)	 - Webhooks that fire when an SDK Connection's payload changes
