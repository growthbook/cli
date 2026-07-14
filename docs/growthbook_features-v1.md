## growthbook features-v1

Control your feature flags programatically

### Synopsis

Control your feature flags programatically.

**These are v1 endpoints.** New integrations should use the v2 Feature Flags endpoints, which expose a unified per-rule environment scope instead of per-environment rule arrays.

```
growthbook features-v1 [flags]
```

### Options

```
  -h, --help   help for features-v1
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
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook features-v1 create](growthbook_features-v1_create.md)	 - Create a single feature
* [growthbook features-v1 delete](growthbook_features-v1_delete.md)	 - Deletes a single feature
* [growthbook features-v1 get](growthbook_features-v1_get.md)	 - Get a single feature
* [growthbook features-v1 get-feature-keys](growthbook_features-v1_get-feature-keys.md)	 - Get list of feature keys
* [growthbook features-v1 get-feature-stale](growthbook_features-v1_get-feature-stale.md)	 - Get stale status for one or more features
* [growthbook features-v1 list](growthbook_features-v1_list.md)	 - Get all features
* [growthbook features-v1 revert](growthbook_features-v1_revert.md)	 - Revert a feature to a specific revision
* [growthbook features-v1 toggle](growthbook_features-v1_toggle.md)	 - Toggle a feature in one or more environments
* [growthbook features-v1 update](growthbook_features-v1_update.md)	 - Partially update a feature
