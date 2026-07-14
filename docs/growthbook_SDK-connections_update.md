## growthbook SDK-connections update

Update a single sdk connection

### Synopsis

Update a single sdk connection

```
growthbook SDK-connections update [flags]
```

### Examples

```
  growthbook SDK-connections update --id <id>
```

### Options

```
  -a, --allowed-custom-fields-in-metadata stringArray   list of values
      --body string                                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --encrypt-payload                                 boolean flag
      --environment string                              string value
      --hash-secure-attributes                          boolean flag
  -h, --help                                            help for update
      --id string                                       The id of the requested resource [required]
      --include-custom-fields-in-metadata               boolean flag
      --include-draft-experiment-refs                   When true, experiment-ref rules linked to draft experiments are included in feature definitions. Off by default.
      --include-draft-experiments                       boolean flag
      --include-experiment-names                        boolean flag
      --include-project-id-in-metadata                  boolean flag
      --include-redirect-experiments                    boolean flag
      --include-rule-ids                                boolean flag
      --include-tags-in-metadata                        boolean flag
      --include-visual-experiments                      boolean flag
  -l, --language string                                 string value
  -n, --name string                                     string value
      --projects stringArray                            list of values
      --proxy-enabled                                   boolean flag
      --proxy-host string                               string value
  -r, --remote-eval-enabled                             boolean flag
      --saved-group-references-enabled                  boolean flag
      --sdk-version string                              string value
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

* [growthbook SDK-connections](growthbook_SDK-connections.md)	 - Client keys and settings for connecting SDKs to a GrowthBook instance
