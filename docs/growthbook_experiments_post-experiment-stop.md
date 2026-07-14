## growthbook experiments post-experiment-stop

Stop an experiment

### Synopsis

Stop an experiment

```
growthbook experiments post-experiment-stop [flags]
```

### Examples

```
  growthbook experiments post-experiment-stop --id <id> --results dnf
```

### Options

```
  -a, --analysis string                                Optional markdown summary displayed on the experiment results page.
      --body string                                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --date-ended string                              Optional ISO datetime for ending the latest phase. Defaults to the current date and time.
  -e, --enable-temporary-rollout releasedVariationId   If true, include this stopped experiment in SDK payload and force the release variation (releasedVariationId) to all traffic.
  -h, --help                                           help for post-experiment-stop
  -i, --id string                                      The id of the requested resource [required]
      --reason string                                  Optional reason for ending the phase stored on the latest phase metadata.
      --released-variation-id string                   Required if enableTemporaryRollout is true. Variation ID (e.g. var_abc123) to release to 100% of traffic eligible for this experiment.
      --results string                                 The experiment conclusion status. (options: dnf, won, lost, inconclusive) [required]
  -w, --winner-variation-id string                     Variation ID (e.g. var_abc123) of the winning variation. Used only as metadata. Required if results is 'won' and there are multiple test variations. Otherwise, defaults to the test variation when results is 'won' and to the baseline variation for other results.
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
