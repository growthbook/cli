## growthbook feature-revisions-v2

Draft revisions for feature flags, including rules, scheduling, and approval workflows

### Synopsis

Draft revisions for feature flags, including rules, scheduling, and approval workflows.

Revision `rules` is a flat array with per-rule scope fields.

```
growthbook feature-revisions-v2 [flags]
```

### Options

```
  -h, --help   help for feature-revisions-v2
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook feature-revisions-v2 delete-feature-revision-log-entry-v2](growthbook_feature-revisions-v2_delete-feature-revision-log-entry-v2.md)	 - Delete an owned revision Comment entry
* [growthbook feature-revisions-v2 delete-feature-revision-rule-ramp-schedule-v2](growthbook_feature-revisions-v2_delete-feature-revision-rule-ramp-schedule-v2.md)	 - Remove ramp schedule from a rule
* [growthbook feature-revisions-v2 delete-feature-revision-rule-v2](growthbook_feature-revisions-v2_delete-feature-revision-rule-v2.md)	 - Delete a rule from a draft revision
* [growthbook feature-revisions-v2 get](growthbook_feature-revisions-v2_get.md)	 - List revisions for a feature
* [growthbook feature-revisions-v2 get-feature-revision-diff-v2](growthbook_feature-revisions-v2_get-feature-revision-diff-v2.md)	 - Diff a revision against another revision
* [growthbook feature-revisions-v2 get-feature-revision-latest-v2](growthbook_feature-revisions-v2_get-feature-revision-latest-v2.md)	 - Get the most recent active draft revision
* [growthbook feature-revisions-v2 get-feature-revision-log-v2](growthbook_feature-revisions-v2_get-feature-revision-log-v2.md)	 - List the activity log for a revision
* [growthbook feature-revisions-v2 get-feature-revision-merge-status-v2](growthbook_feature-revisions-v2_get-feature-revision-merge-status-v2.md)	 - Get merge status for a draft revision
* [growthbook feature-revisions-v2 get-feature-revision-v2](growthbook_feature-revisions-v2_get-feature-revision-v2.md)	 - Get a single feature revision
* [growthbook feature-revisions-v2 list-revisions-v2](growthbook_feature-revisions-v2_list-revisions-v2.md)	 - List revisions across all features
* [growthbook feature-revisions-v2 post-feature-revision-discard-v2](growthbook_feature-revisions-v2_post-feature-revision-discard-v2.md)	 - Discard a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-publish-v2](growthbook_feature-revisions-v2_post-feature-revision-publish-v2.md)	 - Publish a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-rebase-preview-v2](growthbook_feature-revisions-v2_post-feature-revision-rebase-preview-v2.md)	 - Preview a rebase without applying it
* [growthbook feature-revisions-v2 post-feature-revision-rebase-v2](growthbook_feature-revisions-v2_post-feature-revision-rebase-v2.md)	 - Rebase a draft revision onto the current live version
* [growthbook feature-revisions-v2 post-feature-revision-recall-review-v2](growthbook_feature-revisions-v2_post-feature-revision-recall-review-v2.md)	 - Recall a review request (revert to draft)
* [growthbook feature-revisions-v2 post-feature-revision-reopen-v2](growthbook_feature-revisions-v2_post-feature-revision-reopen-v2.md)	 - Reopen a discarded revision as a draft
* [growthbook feature-revisions-v2 post-feature-revision-request-review-v2](growthbook_feature-revisions-v2_post-feature-revision-request-review-v2.md)	 - Request review for a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-revert-v2](growthbook_feature-revisions-v2_post-feature-revision-revert-v2.md)	 - Revert the feature to a prior revision
* [growthbook feature-revisions-v2 post-feature-revision-rule-add-v2](growthbook_feature-revisions-v2_post-feature-revision-rule-add-v2.md)	 - Add a rule to a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-rules-reorder-v2](growthbook_feature-revisions-v2_post-feature-revision-rules-reorder-v2.md)	 - Reorder rules in the revision
* [growthbook feature-revisions-v2 post-feature-revision-submit-review-v2](growthbook_feature-revisions-v2_post-feature-revision-submit-review-v2.md)	 - Submit a review on a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-toggle-v2](growthbook_feature-revisions-v2_post-feature-revision-toggle-v2.md)	 - Toggle an environment on/off in a draft revision
* [growthbook feature-revisions-v2 post-feature-revision-undo-review-v2](growthbook_feature-revisions-v2_post-feature-revision-undo-review-v2.md)	 - Undo a reviewer's own review verdict
* [growthbook feature-revisions-v2 post-feature-revision-v2](growthbook_feature-revisions-v2_post-feature-revision-v2.md)	 - Create a draft revision
* [growthbook feature-revisions-v2 put-feature-revision-archive-v2](growthbook_feature-revisions-v2_put-feature-revision-archive-v2.md)	 - Set archived state in a draft revision
* [growthbook feature-revisions-v2 put-feature-revision-default-value-v2](growthbook_feature-revisions-v2_put-feature-revision-default-value-v2.md)	 - Set the default value in a draft revision
* [growthbook feature-revisions-v2 put-feature-revision-holdout-v2](growthbook_feature-revisions-v2_put-feature-revision-holdout-v2.md)	 - Set holdout in a draft revision
* [growthbook feature-revisions-v2 put-feature-revision-log-comment-v2](growthbook_feature-revisions-v2_put-feature-revision-log-comment-v2.md)	 - Edit the comment text of an owned log entry
* [growthbook feature-revisions-v2 put-feature-revision-metadata-v2](growthbook_feature-revisions-v2_put-feature-revision-metadata-v2.md)	 - Update revision metadata
* [growthbook feature-revisions-v2 put-feature-revision-prerequisites-v2](growthbook_feature-revisions-v2_put-feature-revision-prerequisites-v2.md)	 - Set feature-level prerequisites in a draft revision
* [growthbook feature-revisions-v2 put-feature-revision-rule-ramp-schedule-v2](growthbook_feature-revisions-v2_put-feature-revision-rule-ramp-schedule-v2.md)	 - Set ramp schedule for a rule
* [growthbook feature-revisions-v2 put-feature-revision-rule-v2](growthbook_feature-revisions-v2_put-feature-revision-rule-v2.md)	 - Update a rule in a draft revision
