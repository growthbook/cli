## growthbook experiment-templates update

Update a single experimentTemplate

### Synopsis

Update a single experimentTemplate

```
growthbook experiment-templates update [flags]
```

### Examples

```
  growthbook experiment-templates update --id <id>
```

### Options

```
  -a, --activation-metric string        string value
      --body string                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --custom-fields string            value
      --custom-metric-slices string     list of values
      --datasource string               string value
      --description string              string value
      --disable-sticky-bucketing        boolean flag
  -e, --exposure-query-id string        string value
  -f, --fallback-attribute string       string value
      --goal-metrics stringArray        list of values
      --guardrail-metrics stringArray   list of values
      --hash-attribute string           string value
  -h, --help                            help for update
      --hypothesis string               string value
  -i, --id string                       [required]
  -p, --project string                  string value
      --secondary-metrics stringArray   list of values
      --segment string                  string value
      --skip-partial-data               boolean flag
      --stats-engine string             options: bayesian, frequentist
      --tags stringArray                list of values
      --targeting string                JSON object
      --template-metadata string        JSON object
      --type string                     options: standard
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

* [growthbook experiment-templates](growthbook_experiment-templates.md)	 - Operations for experiment-templates
