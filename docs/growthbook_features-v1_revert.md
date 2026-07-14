## growthbook features-v1 revert

Revert a feature to a specific revision

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [POST /v2/features/:id/revert](#operation/revertFeatureV2) instead.

Creates a new revision whose rules and values match a previously-published revision, then immediately publishes it. This leaves a clear audit trail of the revert action in the revision history.

Returns 403 if the API key lacks permission, or if approval rules are enabled for an affected environment and neither the "REST API always bypasses approval requirements" nor the "Allow reverts without approval" org setting is enabled.

Returns 422 with a list of `warnings` if the restored values no longer validate against the feature's current value type or JSON schema. Re-submit with `?ignoreWarnings=true` to revert anyway.

```
growthbook features-v1 revert [flags]
```

### Examples

```
  growthbook features-v1 revert --id <id> --revision 1155.7
```

### Options

```
      --body string      Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --comment string   string value
  -h, --help             help for revert
  -i, --id string        The id of the requested resource [required]
  -r, --revision float   [required]
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

* [growthbook features-v1](growthbook_features-v1.md)	 - Control your feature flags programatically
