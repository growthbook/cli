## growthbook feature-revisions-v2 post-feature-revision-publish-v2

Publish a draft revision

### Synopsis

Immediately publishes a draft revision, making it the live version of the feature. Any pending ramp actions (`pendingRamp` on rules) are executed atomically — ramp schedules are created or detached as queued.

```
growthbook feature-revisions-v2 post-feature-revision-publish-v2 [flags]
```

### Examples

```
  growthbook feature-revisions-v2 post-feature-revision-publish-v2 --id <id> --version-param 717368
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --comment string         string value
  -h, --help                   help for post-feature-revision-publish-v2
  -i, --id string              [required]
  -v, --version-param string   [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
