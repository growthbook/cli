## growthbook experiments create-snapshot

Create Experiment Snapshot

### Synopsis

Create Experiment Snapshot

```
growthbook experiments create-snapshot [flags]
```

### Examples

```
  growthbook experiments create-snapshot --id <id>
```

### Options

```
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --dimension string      Dimension to break results down by. For Unit Dimensions, use the dimension id (e.g. "dim_abc123"). For Experiment Dimensions, use "exp:<dimensionName>" (e.g. "exp:country"). Built-in pre-exposure dimensions include "pre:date" and, when configured, "pre:activation". Omit this field to create a standard snapshot.
  -h, --help                  help for create-snapshot
  -i, --id string             The experiment id of the experiment to update [required]
  -p, --phase int             Zero-based phase index to snapshot, where 0 is the first experiment phase. Defaults to the latest phase.
  -t, --triggered-by string   Set to "schedule" if you want this request to trigger notifications and other events as it if were a scheduled update. Defaults to manual. (options: manual, schedule)
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

* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
