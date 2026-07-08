## growthbook features-v1 toggle

Toggle a feature in one or more environments

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [POST /v2/features/:id/toggle](#operation/toggleFeatureV2) instead.

Enables or disables a feature in one or more environments simultaneously. Accepts a map of environment name → boolean and immediately publishes the change.

Returns 403 if the API key lacks permission or if approval rules are enabled for an affected environment and the org setting "REST API always bypasses approval requirements" is off.

```
growthbook features-v1 toggle [flags]
```

### Examples

```
  growthbook features-v1 toggle --id <id> --environments '{"key":"<value>"}'
```

### Options

```
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --environments string   [required]
  -h, --help                  help for toggle
  -i, --id string             The id of the requested resource [required]
  -r, --reason string         string value
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

* [growthbook features-v1](growthbook_features-v1.md)	 - Control your feature flags programatically
