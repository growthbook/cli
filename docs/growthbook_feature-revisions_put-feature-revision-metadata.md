## growthbook feature-revisions put-feature-revision-metadata

Update revision metadata (comment, title, feature metadata)

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [PUT /v2/features/:id/revisions/:version/metadata](#operation/putFeatureRevisionMetadataV2) instead.

Updates draft-level metadata (`comment`, `title`) and/or feature-level metadata (owner, project, tags, customFields, jsonSchema, etc.). Merge semantics: omitted fields are left unchanged; any provided field replaces the current value (pass an empty string/array/object to clear). Feature metadata changes are staged on the revision and applied to the feature on publish. Changing `project` requires publish permission on both the old and new project.

```
growthbook feature-revisions put-feature-revision-metadata [flags]
```

### Examples

```
  growthbook feature-revisions put-feature-revision-metadata --id <id> --version-param <value>
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --comment string         string value
      --custom-fields string   value
      --description string     string value
  -h, --help                   help for put-feature-revision-metadata
  -i, --id string              [required]
  -j, --json-schema string     JSON object
  -n, --never-stale            boolean flag
      --owner string           The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --project string         string value
      --tags stringArray       list of values
      --title string           string value
  -v, --version-param string   [required]
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

* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
