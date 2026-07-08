## growthbook feature-revisions-v2 post-feature-revision-request-review-v2

Request review for a draft revision

### Synopsis

Moves the draft into the `pending-review` state and notifies reviewers.

Set `autoPublishOnApproval` to `true` to publish the revision automatically the moment it is approved (GitHub auto-merge model). This requires the org to have auto-publish-on-approval enabled for the feature and the caller to have publish permission; the auto-publish then executes with the caller's authority.

Set `scheduledPublishAt` to a future ISO date-time to defer the auto-publish until that date (it still also requires approval when review is required). Use `scheduledPublishLockEdits` to freeze edits to this draft while the schedule is pending, and `scheduledPublishLockOthers` to block publishing other drafts of this feature in the meantime.

```
growthbook feature-revisions-v2 post-feature-revision-request-review-v2 [flags]
```

### Examples

```
  growthbook feature-revisions-v2 post-feature-revision-request-review-v2 --id <id> --version-param 659874
```

### Options

```
  -a, --auto-publish-on-approval        boolean flag
      --body string                     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --comment string                  string value
  -h, --help                            help for post-feature-revision-request-review-v2
  -i, --id string                       [required]
      --scheduled-publish-at string     date/time value
      --scheduled-publish-lock-edits    boolean flag
      --scheduled-publish-lock-others   boolean flag
  -v, --version-param string            [required]
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
      --password string               HTTP Basic auth: your Secret Key as the username, with an empty password. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: your Secret Key as the username, with an empty password. username
```

### SEE ALSO

* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
