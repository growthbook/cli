## growthbook SDK-connection-revisions

Draft revisions for SDK Connections: staged changes to a connection's settings and webhooks, their review, and their lifecycle (publish, discard, rebase)

### Synopsis

Draft revisions for SDK Connections: staged changes to a connection's settings and webhooks, their review, and their lifecycle (publish, discard, rebase).

Pass `version: "new"` on the edit endpoints to create the draft. Use these when the organization requires approvals for a connection, or to group several changes into one publish.

```
growthbook SDK-connection-revisions [flags]
```

### Options

```
  -h, --help   help for SDK-connection-revisions
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
* [growthbook SDK-connection-revisions delete-SDK-connection-revision-webhook](growthbook_SDK-connection-revisions_delete-SDK-connection-revision-webhook.md)	 - Remove a webhook in a draft revision
* [growthbook SDK-connection-revisions get](growthbook_SDK-connection-revisions_get.md)	 - Get a single SDK Connection revision
* [growthbook SDK-connection-revisions get](growthbook_SDK-connection-revisions_get.md)	 - List revisions for an SDK Connection
* [growthbook SDK-connection-revisions get-SDK-connection-revision-latest](growthbook_SDK-connection-revisions_get-SDK-connection-revision-latest.md)	 - Get the most recent active draft revision
* [growthbook SDK-connection-revisions get-SDK-connection-revision-merge-status](growthbook_SDK-connection-revisions_get-SDK-connection-revision-merge-status.md)	 - Get merge status for a draft revision
* [growthbook SDK-connection-revisions list](growthbook_SDK-connection-revisions_list.md)	 - List SDK Connection revisions across the organization
* [growthbook SDK-connection-revisions post](growthbook_SDK-connection-revisions_post.md)	 - Create a draft revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-discard](growthbook_SDK-connection-revisions_post-SDK-connection-revision-discard.md)	 - Discard a draft revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-publish](growthbook_SDK-connection-revisions_post-SDK-connection-revision-publish.md)	 - Publish a draft revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-rebase](growthbook_SDK-connection-revisions_post-SDK-connection-revision-rebase.md)	 - Rebase a draft revision onto the live SDK Connection
* [growthbook SDK-connection-revisions post-SDK-connection-revision-recall-review](growthbook_SDK-connection-revisions_post-SDK-connection-revision-recall-review.md)	 - Recall a review request
* [growthbook SDK-connection-revisions post-SDK-connection-revision-reopen](growthbook_SDK-connection-revisions_post-SDK-connection-revision-reopen.md)	 - Reopen a discarded revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-request-review](growthbook_SDK-connection-revisions_post-SDK-connection-revision-request-review.md)	 - Request review for a draft revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-revert](growthbook_SDK-connection-revisions_post-SDK-connection-revision-revert.md)	 - Revert the SDK Connection to a prior revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-schedule-publish](growthbook_SDK-connection-revisions_post-SDK-connection-revision-schedule-publish.md)	 - Schedule (or cancel) a deferred publish
* [growthbook SDK-connection-revisions post-SDK-connection-revision-submit-review](growthbook_SDK-connection-revisions_post-SDK-connection-revision-submit-review.md)	 - Submit a review on a draft revision
* [growthbook SDK-connection-revisions post-SDK-connection-revision-undo-review](growthbook_SDK-connection-revisions_post-SDK-connection-revision-undo-review.md)	 - Retract your own review verdict
* [growthbook SDK-connection-revisions post-SDK-connection-revision-webhook](growthbook_SDK-connection-revisions_post-SDK-connection-revision-webhook.md)	 - Add a webhook in a draft revision
* [growthbook SDK-connection-revisions put-SDK-connection-revision-archive](growthbook_SDK-connection-revisions_put-SDK-connection-revision-archive.md)	 - Stage an archive or restore in a draft revision
* [growthbook SDK-connection-revisions put-SDK-connection-revision-settings](growthbook_SDK-connection-revisions_put-SDK-connection-revision-settings.md)	 - Update SDK Connection settings in a draft revision
* [growthbook SDK-connection-revisions put-SDK-connection-revision-webhook](growthbook_SDK-connection-revisions_put-SDK-connection-revision-webhook.md)	 - Update a webhook in a draft revision
