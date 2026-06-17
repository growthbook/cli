## growthbook feature-revisions-v2 post-feature-revision-rebase-v2

Rebase a draft revision onto the current live version

### Synopsis

Updates the draft's base revision to match the currently-live revision, applying the draft's changes on top. Supply `conflictResolutions` to resolve conflicting items individually — including per-rule (`rules.<ruleId>`) and rule-order (`rules.order`) conflicts. Supply `expectedLiveVersion` and/or `expectedDraftDateUpdated` (both returned by merge-status and rebase preview) to fail fast with `409` if either side changes between conflict review and submission. Unresolved conflicts also respond with `409`.

```
growthbook feature-revisions-v2 post-feature-revision-rebase-v2 [flags]
```

### Examples

```
  growthbook feature-revisions-v2 post-feature-revision-rebase-v2 --id <id> --version-param 985805
```

### Options

```
      --body string                                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --conflict-resolutions defaultValue              Map of conflict key → resolution. Keys come from the returned conflicts: defaultValue, `prerequisites`, `archived`, `holdout`, `environmentsEnabled.<env>`, `metadata.<field>`, `rules.<ruleId>`, and `rules.order`. `overwrite` keeps the draft's version of that item; `discard` keeps live's. The blanket `rules` key applies one strategy to all rule-level conflicts.
      --expected-draft-date-updated draftDateUpdated   Optimistic-concurrency guard for the draft side: the draft's draftDateUpdated timestamp as returned by merge-status or rebase preview. If the draft has been modified since (e.g. by a co-author), the request fails with `409` instead of applying resolutions against changed draft content.
      --expected-live-version 409                      Optimistic-concurrency guard: the live version the resolutions were authored against (as returned by merge-status or rebase preview). If live has since moved, the request fails with 409 instead of applying resolutions to different conflicts.
  -h, --help                                           help for post-feature-revision-rebase-v2
  -i, --id string                                      [required]
  -v, --version-param string                           [required]
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

* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
