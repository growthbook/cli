## growthbook contextual-bandits

Operations for contextual-bandits

### Synopsis

Operations for contextual-bandits

```
growthbook contextual-bandits [flags]
```

### Options

```
  -h, --help   help for contextual-bandits
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
* [growthbook contextual-bandits create](growthbook_contextual-bandits_create.md)	 - Create a single contextualBandit
* [growthbook contextual-bandits delete-contextual-bandit-linked-feature](growthbook_contextual-bandits_delete-contextual-bandit-linked-feature.md)	 - Unlink a feature from a Contextual Bandit
* [growthbook contextual-bandits get](growthbook_contextual-bandits_get.md)	 - Get a single contextualBandit
* [growthbook contextual-bandits get-contextual-bandit-current-weights](growthbook_contextual-bandits_get-contextual-bandit-current-weights.md)	 - Get current Contextual Bandit leaf weights and latest event
* [growthbook contextual-bandits get-contextual-bandit-event](growthbook_contextual-bandits_get-contextual-bandit-event.md)	 - Get a single Contextual Bandit weight-update event
* [growthbook contextual-bandits get-contextual-bandit-linked-features](growthbook_contextual-bandits_get-contextual-bandit-linked-features.md)	 - Get features linked to a Contextual Bandit
* [growthbook contextual-bandits get-contextual-bandit-results](growthbook_contextual-bandits_get-contextual-bandit-results.md)	 - Get latest Contextual Bandit results
* [growthbook contextual-bandits get-contextual-bandit-snapshot](growthbook_contextual-bandits_get-contextual-bandit-snapshot.md)	 - Get a single Contextual Bandit snapshot
* [growthbook contextual-bandits list](growthbook_contextual-bandits_list.md)	 - Get all contextualBandits
* [growthbook contextual-bandits list-contextual-bandit-events](growthbook_contextual-bandits_list-contextual-bandit-events.md)	 - List Contextual Bandit weight-update events
* [growthbook contextual-bandits list-contextual-bandit-snapshots](growthbook_contextual-bandits_list-contextual-bandit-snapshots.md)	 - List Contextual Bandit snapshots
* [growthbook contextual-bandits refresh](growthbook_contextual-bandits_refresh.md)	 - Trigger a Contextual Bandit snapshot refresh
* [growthbook contextual-bandits start](growthbook_contextual-bandits_start.md)	 - Start a Contextual Bandit
* [growthbook contextual-bandits stop](growthbook_contextual-bandits_stop.md)	 - Stop a Contextual Bandit
* [growthbook contextual-bandits update](growthbook_contextual-bandits_update.md)	 - Update a single contextualBandit
