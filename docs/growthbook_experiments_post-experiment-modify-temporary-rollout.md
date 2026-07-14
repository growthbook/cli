## growthbook experiments post-experiment-modify-temporary-rollout

Modify temporary rollout status for a stopped experiment

### Synopsis

Modify temporary rollout status for a stopped experiment

```
growthbook experiments post-experiment-modify-temporary-rollout [flags]
```

### Examples

```
  growthbook experiments post-experiment-modify-temporary-rollout --id <id> --enable-temporary-rollout true
```

### Options

```
      --body string                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --enable-temporary-rollout       If true, keep the stopped experiment in SDK payload and force traffic to the winner variation. If false, end temporary rollout and remove from SDK payload. [required]
  -h, --help                           help for post-experiment-modify-temporary-rollout
  -i, --id string                      The id of the requested resource [required]
  -r, --released-variation-id string   Variation ID (e.g. var_abc123) to release to 100% of traffic eligible for this experiment. Required if enableTemporaryRollout is true.
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

* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
