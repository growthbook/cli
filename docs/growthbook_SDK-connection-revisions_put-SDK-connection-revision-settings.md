## growthbook SDK-connection-revisions put-SDK-connection-revision-settings

Update SDK Connection settings in a draft revision

### Synopsis

Stages settings changes on the draft. Takes the same fields as `PUT /sdk-connections/{id}`, and only the fields sent are changed. Pass `version: "new"` to create the draft. Nothing changes on the live SDK Connection until the revision is published.

```
growthbook SDK-connection-revisions put-SDK-connection-revision-settings [flags]
```

### Examples

```
  growthbook SDK-connection-revisions put-SDK-connection-revision-settings --id <id> --version-param <value>
```

### Options

```
  -a, --allowed-custom-fields-in-metadata stringArray     list of values
      --body string                                       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --encrypt-payload                                   boolean flag
      --environment string                                string value
      --hash-secure-attributes                            boolean flag
  -h, --help                                              help for put-SDK-connection-revision-settings
      --id string                                         The SDK Connection id [required]
      --ignore-warnings                                   Set to true to continue past a warning: one raised by a Custom Hook, or a webhook that references a secret that doesn't exist. On publish, it also force-publishes an out-of-date draft when the caller has Bypass draft approvals access. It never clears a Custom Hook rejection. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
      --include-custom-fields-in-metadata                 boolean flag
      --include-draft-experiment-refs                     When true, experiment-ref rules linked to draft experiments are included in feature definitions. Off by default.
      --include-draft-experiments                         boolean flag
      --include-experiment-names                          boolean flag
      --include-experiment-schedule-in-metadata           boolean flag
      --include-project-id-in-metadata                    boolean flag
      --include-redirect-experiments                      boolean flag
      --include-referenced-prerequisites                  Carry prerequisite Feature Flags into this payload even when they target other Projects. Defaults to true for new connections.
      --include-rule-ids                                  boolean flag
      --include-tags-in-metadata                          boolean flag
      --include-visual-experiments                        boolean flag
  -l, --language string                                   string value
  -n, --name Web app                                      Short display name, usually app and platform (e.g. Web app).
      --projects stringArray                              list of values
      --proxy-enabled                                     boolean flag
      --proxy-host string                                 string value
      --remote-eval-enabled                               boolean flag
      --revision-comment version                          Markdown. What changed and why, for reviewers. Only used when version is "new"; ignored for existing drafts.
      --revision-title Enable for beta testers            Short draft title (e.g. Enable for beta testers). Only used when `version` is "new"; ignored for existing drafts.
      --saved-group-format string                         options: inline, referencesV1, referencesV2
      --saved-group-references-enabled savedGroupFormat   Deprecated. Use savedGroupFormat.
      --sdk-version string                                string value
      --skip-hooks                                        Set to true to continue despite a Custom Hook rejection. The caller must have Bypass draft approvals access for SDK Connections in every Project. Otherwise, this field is ignored. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
  -t, --traces-enabled                                    Send SDK evaluation traces to GrowthBook. Only available to organizations with tracing enabled.
  -v, --version-param string                              [required]
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
