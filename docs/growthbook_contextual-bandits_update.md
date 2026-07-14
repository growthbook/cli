## growthbook contextual-bandits update

Update a single contextualBandit

### Synopsis

Update a single contextualBandit

```
growthbook contextual-bandits update [flags]
```

### Examples

```
  growthbook contextual-bandits update --id <id>
```

### Options

```
  -a, --archived                            boolean flag
      --body string                         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --burn-in-unit string                 options: days, hours
      --burn-in-value float                 number value
      --condition string                    string value
      --contextual-attributes stringArray   list of values
      --contextual-bandit-query-id string   string value
      --conversion-window-unit string       options: hours, days
      --conversion-window-value string      number value
      --coverage float                      number value
      --datasource string                   string value
      --decision-metric string              string value
      --description string                  string value
      --hash-attribute string               string value
  -h, --help                                help for update
  -i, --id string                           [required]
      --max-leaves int                      integer value
      --min-users-per-leaf int              integer value
  -n, --name string                         string value
      --owner string                        The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --prerequisites string                list of values
      --project string                      string value
      --saved-groups string                 list of values
      --schedule-unit string                options: days, hours
      --schedule-value float                number value
      --seed string                         string value
      --status string                       options: draft, running, stopped
      --tags stringArray                    list of values
      --tracking-key string                 string value
      --variation-weights string            list of values
      --variations string                   list of values
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

* [growthbook contextual-bandits](growthbook_contextual-bandits.md)	 - Operations for contextual-bandits
