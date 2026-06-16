## growthbook features

Control your feature flags programatically

### Synopsis

Control your feature flags programatically.

**These are v1 endpoints.** New integrations should use the v2 Feature Flags endpoints, which expose a unified per-rule environment scope instead of per-environment rule arrays.

```
growthbook features [flags]
```

### Options

```
  -h, --help   help for features
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
* [growthbook features delete](growthbook_features_delete.md)	 - Deletes a single feature
* [growthbook features get](growthbook_features_get.md)	 - Get a single feature
* [growthbook features get-feature-keys](growthbook_features_get-feature-keys.md)	 - Get list of feature keys
* [growthbook features get-feature-stale](growthbook_features_get-feature-stale.md)	 - Get stale status for one or more features
* [growthbook features list](growthbook_features_list.md)	 - Get all features
* [growthbook features post](growthbook_features_post.md)	 - Create a single feature
* [growthbook features revert](growthbook_features_revert.md)	 - Revert a feature to a specific revision
* [growthbook features toggle](growthbook_features_toggle.md)	 - Toggle a feature in one or more environments
* [growthbook features update](growthbook_features_update.md)	 - Partially update a feature
