## growthbook feature-revisions-v1 post-feature-revision-rule-add

Add a rule to a draft revision

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [POST /v2/features/:id/revisions/:version/rules](#operation/postFeatureRevisionRuleAddV2) instead, which accepts rules with unified `allEnvironments`/`environments` scope fields instead of a per-environment `environment` parameter.

Appends a new rule to the end of the rule list for the given environment. A `rule.type` of `force`, `rollout`, `experiment-ref`, or `safe-rollout` determines the accepted shape. Use `rampSchedule` for ramp configuration or `schedule` for a simple start/end window; if both are provided, `rampSchedule` wins.

```
growthbook feature-revisions-v1 post-feature-revision-rule-add [flags]
```

### Examples

```
  growthbook feature-revisions-v1 post-feature-revision-rule-add --id <id> --version-param <value> --environment <value> --rule '{"value":"<value>"}'
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --environment string        [required]
  -h, --help                      help for post-feature-revision-rule-add
  -i, --id string                 [required]
      --ramp-schedule string      JSON object
      --revision-comment string   string value
      --revision-title string     string value
      --rule string               JSON value (one of: { description: string, enabled: boolean, condition: string, savedGroups: object[], ... })
  -s, --schedule string           JSON object
  -v, --version-param string      [required]
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

* [growthbook feature-revisions-v1](growthbook_feature-revisions-v1.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
