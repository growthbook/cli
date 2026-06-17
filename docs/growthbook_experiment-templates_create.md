## growthbook experiment-templates create

Create a single experimentTemplate

### Synopsis

Create a single experimentTemplate

```
growthbook experiment-templates create [flags]
```

### Examples

```
  growthbook experiment-templates create --template-metadata '{"name":"<value>"}' --type standard --datasource <value> --exposure-query-id <id> --stats-engine bayesian
```

### Options

```
  -a, --activation-metric string        string value
      --body string                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --custom-fields string            value
      --custom-metric-slices string     list of values
      --datasource string               [required]
      --description string              string value
      --disable-sticky-bucketing        boolean flag
  -e, --exposure-query-id string        [required]
  -f, --fallback-attribute string       string value
      --goal-metrics stringArray        list of values
      --guardrail-metrics stringArray   list of values
      --hash-attribute string           string value
  -h, --help                            help for create
      --hypothesis string               string value
  -p, --project string                  string value
      --secondary-metrics stringArray   list of values
      --segment string                  string value
      --skip-partial-data               boolean flag
      --stats-engine string             options: bayesian, frequentist [required]
      --tags stringArray                list of values
      --targeting string                [required]
      --template-metadata string        [required]
      --type string                     options: standard [required]
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

* [growthbook experiment-templates](growthbook_experiment-templates.md)	 - Operations for experiment-templates
