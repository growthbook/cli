## growthbook experiments

Experiments (A/B Tests)

### Synopsis

Experiments (A/B Tests)

```
growthbook experiments [flags]
```

### Options

```
  -h, --help   help for experiments
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook experiments complete-start-checklist](growthbook_experiments_complete-start-checklist.md)	 - Mark manual pre-launch checklist items complete
* [growthbook experiments create](growthbook_experiments_create.md)	 - Create a single experiment
* [growthbook experiments create-snapshot](growthbook_experiments_create-snapshot.md)	 - Create Experiment Snapshot
* [growthbook experiments delete-variation-screenshot](growthbook_experiments_delete-variation-screenshot.md)	 - Delete a variation screenshot
* [growthbook experiments get](growthbook_experiments_get.md)	 - Get a single experiment
* [growthbook experiments get-start-checklist](growthbook_experiments_get-start-checklist.md)	 - Get an experiment pre-launch checklist status
* [growthbook experiments list](growthbook_experiments_list.md)	 - Get all experiments
* [growthbook experiments list-names](growthbook_experiments_list-names.md)	 - Get a list of experiments with names and ids
* [growthbook experiments list-results](growthbook_experiments_list-results.md)	 - Get latest results for many experiments
* [growthbook experiments modify-temporary-rollout](growthbook_experiments_modify-temporary-rollout.md)	 - Modify temporary rollout status for a stopped experiment
* [growthbook experiments results](growthbook_experiments_results.md)	 - Get results for an experiment
* [growthbook experiments start](growthbook_experiments_start.md)	 - Start/Stage an experiment
* [growthbook experiments stop](growthbook_experiments_stop.md)	 - Stop an experiment
* [growthbook experiments update](growthbook_experiments_update.md)	 - Update a single experiment
* [growthbook experiments upload-variation-screenshot](growthbook_experiments_upload-variation-screenshot.md)	 - Upload a variation screenshot
