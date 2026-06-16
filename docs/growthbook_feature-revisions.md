## growthbook feature-revisions

Draft revisions for feature flags, including rules, scheduling, and approval workflows

### Synopsis

Draft revisions for feature flags, including rules, scheduling, and approval workflows.

**These are v1 endpoints.** New integrations should use the v2 Feature Revisions endpoints.

```
growthbook feature-revisions [flags]
```

### Options

```
  -h, --help   help for feature-revisions
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook feature-revisions delete-feature-revision-rule](growthbook_feature-revisions_delete-feature-revision-rule.md)	 - Delete a rule from a draft revision
* [growthbook feature-revisions delete-feature-revision-rule-ramp-schedule](growthbook_feature-revisions_delete-feature-revision-rule-ramp-schedule.md)	 - Remove ramp schedule from a rule
* [growthbook feature-revisions get](growthbook_feature-revisions_get.md)	 - List revisions for a feature
* [growthbook feature-revisions get](growthbook_feature-revisions_get.md)	 - Get a single feature revision
* [growthbook feature-revisions get-feature-revision-latest](growthbook_feature-revisions_get-feature-revision-latest.md)	 - Get the most recent active draft revision
* [growthbook feature-revisions get-feature-revision-merge-status](growthbook_feature-revisions_get-feature-revision-merge-status.md)	 - Get merge status for a draft revision
* [growthbook feature-revisions list-revisions](growthbook_feature-revisions_list-revisions.md)	 - List feature revisions
* [growthbook feature-revisions post](growthbook_feature-revisions_post.md)	 - Create a draft revision
* [growthbook feature-revisions post-feature-revision-discard](growthbook_feature-revisions_post-feature-revision-discard.md)	 - Discard a draft revision
* [growthbook feature-revisions post-feature-revision-publish](growthbook_feature-revisions_post-feature-revision-publish.md)	 - Publish a draft revision
* [growthbook feature-revisions post-feature-revision-rebase](growthbook_feature-revisions_post-feature-revision-rebase.md)	 - Rebase a draft revision onto the current live version
* [growthbook feature-revisions post-feature-revision-request-review](growthbook_feature-revisions_post-feature-revision-request-review.md)	 - Request review for a draft revision
* [growthbook feature-revisions post-feature-revision-revert](growthbook_feature-revisions_post-feature-revision-revert.md)	 - Revert the feature to a prior revision
* [growthbook feature-revisions post-feature-revision-rule-add](growthbook_feature-revisions_post-feature-revision-rule-add.md)	 - Add a rule to a draft revision
* [growthbook feature-revisions post-feature-revision-rules-reorder](growthbook_feature-revisions_post-feature-revision-rules-reorder.md)	 - Reorder rules in an environment
* [growthbook feature-revisions post-feature-revision-submit-review](growthbook_feature-revisions_post-feature-revision-submit-review.md)	 - Submit a review on a draft revision
* [growthbook feature-revisions post-feature-revision-toggle](growthbook_feature-revisions_post-feature-revision-toggle.md)	 - Toggle an environment on/off in a draft revision
* [growthbook feature-revisions put-feature-revision-archive](growthbook_feature-revisions_put-feature-revision-archive.md)	 - Set archived state in a draft revision
* [growthbook feature-revisions put-feature-revision-default-value](growthbook_feature-revisions_put-feature-revision-default-value.md)	 - Set the default value in a draft revision
* [growthbook feature-revisions put-feature-revision-holdout](growthbook_feature-revisions_put-feature-revision-holdout.md)	 - Set holdout in a draft revision
* [growthbook feature-revisions put-feature-revision-metadata](growthbook_feature-revisions_put-feature-revision-metadata.md)	 - Update revision metadata (comment, title, feature metadata)
* [growthbook feature-revisions put-feature-revision-prerequisites](growthbook_feature-revisions_put-feature-revision-prerequisites.md)	 - Set feature-level prerequisites in a draft revision
* [growthbook feature-revisions put-feature-revision-rule](growthbook_feature-revisions_put-feature-revision-rule.md)	 - Update a rule in a draft revision
* [growthbook feature-revisions put-feature-revision-rule-ramp-schedule](growthbook_feature-revisions_put-feature-revision-rule-ramp-schedule.md)	 - Set ramp schedule for a rule
