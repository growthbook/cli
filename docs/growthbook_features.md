## growthbook features

Control your feature flags programatically

### Synopsis

Control your feature flags programatically.

Rules are returned as a unified top-level array; each rule carries `allEnvironments` / `environments` scope fields instead of being bucketed by environment.

```
growthbook features [flags]
```

### Options

```
  -h, --help   help for features
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
* [growthbook features create](growthbook_features_create.md)	 - Create a single feature
* [growthbook features delete](growthbook_features_delete.md)	 - Deletes a single feature
* [growthbook features get](growthbook_features_get.md)	 - Get a single feature
* [growthbook features get-feature-keys](growthbook_features_get-feature-keys.md)	 - Get list of feature keys
* [growthbook features get-feature-stale](growthbook_features_get-feature-stale.md)	 - Get stale status for one or more features
* [growthbook features list](growthbook_features_list.md)	 - Get all features
* [growthbook features revert](growthbook_features_revert.md)	 - Revert a feature to a specific revision
* [growthbook features toggle](growthbook_features_toggle.md)	 - Toggle a feature in one or more environments
* [growthbook features update](growthbook_features_update.md)	 - Partially update a feature
