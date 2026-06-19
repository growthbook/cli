## growthbook features-v2

Control your feature flags programatically

### Synopsis

Control your feature flags programatically.

Rules are returned as a unified top-level array; each rule carries `allEnvironments` / `environments` scope fields instead of being bucketed by environment.

```
growthbook features-v2 [flags]
```

### Options

```
  -h, --help   help for features-v2
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook features-v2 delete-feature-v2](growthbook_features-v2_delete-feature-v2.md)	 - Deletes a single feature
* [growthbook features-v2 get-feature-keys-v2](growthbook_features-v2_get-feature-keys-v2.md)	 - Get list of feature keys
* [growthbook features-v2 get-feature-stale-v2](growthbook_features-v2_get-feature-stale-v2.md)	 - Get stale status for one or more features
* [growthbook features-v2 get-feature-v2](growthbook_features-v2_get-feature-v2.md)	 - Get a single feature
* [growthbook features-v2 list](growthbook_features-v2_list.md)	 - Get all features
* [growthbook features-v2 post-feature-v2](growthbook_features-v2_post-feature-v2.md)	 - Create a single feature
* [growthbook features-v2 revert-feature-v2](growthbook_features-v2_revert-feature-v2.md)	 - Revert a feature to a specific revision
* [growthbook features-v2 toggle-feature-v2](growthbook_features-v2_toggle-feature-v2.md)	 - Toggle a feature in one or more environments
* [growthbook features-v2 update-feature-v2](growthbook_features-v2_update-feature-v2.md)	 - Partially update a feature
