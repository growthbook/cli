## growthbook interleavings create

Create a single interleaving

### Synopsis

Create a single interleaving

```
growthbook interleavings create [flags]
```

### Examples

```
  growthbook interleavings create --name <value> --key <key> --variations '{"control":{"key":"<key>","name":"<value>","config":"<value>"},"treatment":{"key":"<key>","name":"<value>","config":"<value>"} }'
```

### Options

```
      --body string                             Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --condition string                        string value
      --coverage float                          number value
      --datasource-id string                    string value
      --description string                      string value
  -e, --environment-settings string             value
      --hash-attribute string                   string value
  -h, --help                                    help for create
  -i, --interleaving-exposure-query-id string   string value
  -j, --json-schema date                        Validation schema to stage on the draft. The server sets date, so don't send it.
  -k, --key interleave()                        The key your code passes to interleave() as `key`. Can't be changed later. Several experiments may share a key so you can iterate without changing code, but only one of them can be running in a given environment. [required]
  -m, --metrics string                          list of values
  -n, --name string                             [required]
      --owner string                            The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --prerequisites string                    list of values
      --project string                          string value
  -s, --saved-groups string                     list of values
      --tags stringArray                        list of values
      --tracking-key experiment_id              The experiment_id recorded in exposures. Leave this unset: it is generated from `key`, with a `_v1` style suffix when another experiment already uses it. Tracking keys must be unique across the organization, so supplying one that is taken is rejected.
  -u, --user-id-type string                     string value
  -v, --variations string                       Exactly two rankers, by role. [required]
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

* [growthbook interleavings](growthbook_interleavings.md)	 - Operations for interleavings
