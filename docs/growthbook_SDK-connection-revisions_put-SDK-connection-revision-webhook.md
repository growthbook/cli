## growthbook SDK-connection-revisions put-SDK-connection-revision-webhook

Update a webhook in a draft revision

### Synopsis

Stages changes to one of the draft's webhooks. Only the fields sent are changed. Pass `version: "new"` to create the draft.

```
growthbook SDK-connection-revisions put-SDK-connection-revision-webhook [flags]
```

### Examples

```
  growthbook SDK-connection-revisions put-SDK-connection-revision-webhook --id <id> --version-param <value> --webhook-id <id>
```

### Options

```
      --body string                              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --endpoint string                          string value
      --headers string                           JSON-encoded object, e.g. "{\"X-Token\": \"abc\"}"
  -h, --help                                     help for put-SDK-connection-revision-webhook
      --http-method string                       options: GET, POST, PUT, DELETE, PATCH, PURGE
      --id string                                The SDK Connection id [required]
      --ignore-warnings                          Set to true to continue past a warning: one raised by a Custom Hook, or a webhook that references a secret that doesn't exist. On publish, it also force-publishes an out-of-date draft when the caller has Bypass draft approvals access. It never clears a Custom Hook rejection. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
  -n, --name string                              string value
      --payload-format string                    options: standard, standard-no-payload, sdkPayload, edgeConfig, edgeConfigUnescaped, none
      --payload-key string                       string value
      --revision-comment version                 Markdown. What changed and why, for reviewers. Only used when version is "new"; ignored for existing drafts.
      --revision-title Enable for beta testers   Short draft title (e.g. Enable for beta testers). Only used when `version` is "new"; ignored for existing drafts.
  -s, --skip-hooks                               Set to true to continue despite a Custom Hook rejection. The caller must have Bypass draft approvals access for SDK Connections in every Project. Otherwise, this field is ignored. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
  -t, --timeout-seconds string                   Request timeout in seconds, up to the server maximum (120 unless configured). Null uses the default.
  -v, --version-param string                     [required]
  -w, --webhook-id string                        The webhook's id in this revision [required]
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

* [growthbook SDK-connection-revisions](growthbook_SDK-connection-revisions.md)	 - Draft revisions for SDK Connections: staged changes to a connection's settings and webhooks, their review, and their lifecycle (publish, discard, rebase)
