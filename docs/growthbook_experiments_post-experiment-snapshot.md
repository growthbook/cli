## growthbook experiments post-experiment-snapshot

Create Experiment Snapshot

### Synopsis

Create Experiment Snapshot

```
growthbook experiments post-experiment-snapshot [flags]
```

### Examples

```
  growthbook experiments post-experiment-snapshot --id <id>
```

### Options

```
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --dimension string      Dimension to break results down by. For Unit Dimensions, use the dimension id (e.g. "dim_abc123"). For Experiment Dimensions, use "exp:<dimensionName>" (e.g. "exp:country"). Built-in pre-exposure dimensions include "pre:date" and, when configured, "pre:activation". Omit this field to create a standard snapshot.
  -h, --help                  help for post-experiment-snapshot
  -i, --id string             The experiment id of the experiment to update [required]
  -p, --phase int             Zero-based phase index to snapshot, where 0 is the first experiment phase. Defaults to the latest phase.
  -t, --triggered-by string   Set to "schedule" if you want this request to trigger notifications and other events as it if were a scheduled update. Defaults to manual. (options: manual, schedule)
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

* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
