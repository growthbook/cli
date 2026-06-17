## growthbook features revert

Revert a feature to a specific revision

### Synopsis

Creates a new revision whose rules and values match a previously-published revision, then immediately publishes it, leaving a clear audit trail of the revert in the revision history.

Returns 403 if the API key lacks permission, or if approval rules are enabled for an affected environment and neither the "REST API always bypasses approval requirements" nor the "Allow reverts without approval" org setting is enabled.

Returns 422 with a list of `warnings` if the restored values no longer validate against the feature's current value type or JSON schema (e.g. reverting to a config the current schema can no longer read). Re-submit with `?ignoreWarnings=true` to revert anyway.

```
growthbook features revert [flags]
```

### Examples

```
  growthbook features revert --id <id> --revision 255.61
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

* [growthbook features](growthbook_features.md)	 - Control your feature flags programatically
