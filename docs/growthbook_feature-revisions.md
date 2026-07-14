## growthbook feature-revisions

Draft revisions for feature flags, including rules, scheduling, and approval workflows

### Synopsis

Draft revisions for feature flags, including rules, scheduling, and approval workflows.

Revision `rules` is a flat array with per-rule scope fields.

```
growthbook feature-revisions [flags]
```

### Options

```
  -h, --help   help for feature-revisions
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook feature-revisions add-log-comment](growthbook_feature-revisions_add-log-comment.md)	 - Edit the comment text of an owned log entry
* [growthbook feature-revisions add-rule](growthbook_feature-revisions_add-rule.md)	 - Add a rule to a draft revision
* [growthbook feature-revisions archive](growthbook_feature-revisions_archive.md)	 - Set archived state in a draft revision
* [growthbook feature-revisions create](growthbook_feature-revisions_create.md)	 - Create a draft revision
* [growthbook feature-revisions delete-log-entry](growthbook_feature-revisions_delete-log-entry.md)	 - Delete an owned revision Comment entry
* [growthbook feature-revisions delete-rule](growthbook_feature-revisions_delete-rule.md)	 - Delete a rule from a draft revision
* [growthbook feature-revisions delete-rule-ramp-schedule](growthbook_feature-revisions_delete-rule-ramp-schedule.md)	 - Remove ramp schedule from a rule
* [growthbook feature-revisions diff](growthbook_feature-revisions_diff.md)	 - Diff a revision against another revision
* [growthbook feature-revisions discard](growthbook_feature-revisions_discard.md)	 - Discard a draft revision
* [growthbook feature-revisions get](growthbook_feature-revisions_get.md)	 - Get a single feature revision
* [growthbook feature-revisions latest](growthbook_feature-revisions_latest.md)	 - Get the most recent active draft revision
* [growthbook feature-revisions list-for-feature](growthbook_feature-revisions_list-for-feature.md)	 - List revisions for a feature
* [growthbook feature-revisions list-revisions](growthbook_feature-revisions_list-revisions.md)	 - List revisions across all features
* [growthbook feature-revisions log](growthbook_feature-revisions_log.md)	 - List the activity log for a revision
* [growthbook feature-revisions merge-status](growthbook_feature-revisions_merge-status.md)	 - Get merge status for a draft revision
* [growthbook feature-revisions publish](growthbook_feature-revisions_publish.md)	 - Publish a draft revision
* [growthbook feature-revisions rebase](growthbook_feature-revisions_rebase.md)	 - Rebase a draft revision onto the current live version
* [growthbook feature-revisions rebase-preview](growthbook_feature-revisions_rebase-preview.md)	 - Preview a rebase without applying it
* [growthbook feature-revisions recall-review](growthbook_feature-revisions_recall-review.md)	 - Recall a review request (revert to draft)
* [growthbook feature-revisions reopen](growthbook_feature-revisions_reopen.md)	 - Reopen a discarded revision as a draft
* [growthbook feature-revisions reorder-rules](growthbook_feature-revisions_reorder-rules.md)	 - Reorder rules in the revision
* [growthbook feature-revisions request-review](growthbook_feature-revisions_request-review.md)	 - Request review for a draft revision
* [growthbook feature-revisions revert](growthbook_feature-revisions_revert.md)	 - Revert the feature to a prior revision
* [growthbook feature-revisions schedule-publish](growthbook_feature-revisions_schedule-publish.md)	 - Schedule (or cancel) a deferred publish for a draft revision
* [growthbook feature-revisions set-default-value](growthbook_feature-revisions_set-default-value.md)	 - Set the default value in a draft revision
* [growthbook feature-revisions set-holdout](growthbook_feature-revisions_set-holdout.md)	 - Set holdout in a draft revision
* [growthbook feature-revisions set-metadata](growthbook_feature-revisions_set-metadata.md)	 - Update revision metadata
* [growthbook feature-revisions set-prerequisites](growthbook_feature-revisions_set-prerequisites.md)	 - Set feature-level prerequisites in a draft revision
* [growthbook feature-revisions set-rule-ramp-schedule](growthbook_feature-revisions_set-rule-ramp-schedule.md)	 - Set ramp schedule for a rule
* [growthbook feature-revisions submit-review](growthbook_feature-revisions_submit-review.md)	 - Submit a review on a draft revision
* [growthbook feature-revisions toggle](growthbook_feature-revisions_toggle.md)	 - Toggle an environment on/off in a draft revision
* [growthbook feature-revisions undo-review](growthbook_feature-revisions_undo-review.md)	 - Undo a reviewer's own review verdict
* [growthbook feature-revisions update-rule](growthbook_feature-revisions_update-rule.md)	 - Update a rule in a draft revision
