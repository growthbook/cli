## growthbook releases publish-revisions

Atomically publish revisions across multiple entities

### Synopsis

Publishes a set of revisions — at most one per entity — across Feature Flags, Saved Groups, configs, and constants as a single all-or-nothing operation.

Validation, guards, and custom hooks run against the combined end-state of the whole set, so interdependent changes (e.g. a config schema change plus the values that depend on it) publish together even when the in-between states would be invalid.

A blocked publish returns one 422 listing every gate across every item and the flag that clears each. A concurrent change to any target aborts with a 409 and nothing publishes. A failure after the commit starts rolls everything back and emits `revision.publishFailed` for each revision in the set. SDK payloads refresh once per request. Pass `dryRun: true` for the full gate report with zero writes.

Requires the `releases` commercial feature.

```
growthbook releases publish-revisions [flags]
```

### Examples

```
  growthbook releases publish-revisions --revisions '[{"entityType":"constant","revisionId":"<id>"}]'
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --comment string           An optional publish comment recorded on every revision in this release — it appears in each entity's revision history and is passed to any custom validation hooks that run for the publish.
      --dry-run                  Report every gate and outcome without writing anything.
  -h, --help                     help for publish-revisions
  -i, --ignore-warnings          Set to true to acknowledge warnings from experiment guards, schema changes, or archives. It can also force-publish an out-of-date draft when the caller has Bypass draft approvals access for that resource.
  -r, --revisions string         The revisions to publish, at most one per entity. [required]
      --skip-hooks               Set to true to publish despite Custom Hook rejections. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project.
      --skip-schema-validation   Set to true to publish despite schema or invariant failures. Validation still runs and the response still reports each failure. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project.
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
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

* [growthbook releases](growthbook_releases.md)	 - **Beta** — these endpoints are new and may change in backwards-incompatible ways
