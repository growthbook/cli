## growthbook SDK-webhooks put

Update an SDK webhook

### Synopsis

The webhook fires once after an update. Webhooks managed by the Vercel integration can't be changed.

```
growthbook SDK-webhooks put [flags]
```

### Examples

```
  growthbook SDK-webhooks put --id <id>
```

### Options

```
      --body string                              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --endpoint string                          string value
      --headers string                           JSON-encoded object, e.g. "{\"X-Token\": \"abc\"}"
  -h, --help                                     help for put
      --http-method string                       options: GET, POST, PUT, DELETE, PATCH, PURGE
  -i, --id string                                [required]
  -n, --name string                              string value
      --payload-format vercelNativeIntegration   vercelNativeIntegration is accepted only on webhooks that already use it (options: standard, standard-no-payload, sdkPayload, edgeConfig, edgeConfigUnescaped, vercelNativeIntegration, none)
      --payload-key string                       string value
  -s, --sdk-connections stringArray              Replaces the SDK Connections this webhook fires for
  -t, --timeout-seconds string                   Request timeout in seconds, up to the server maximum (120 unless configured). Null uses the default.
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
