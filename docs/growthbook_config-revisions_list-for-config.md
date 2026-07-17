## growthbook config-revisions list-for-config

List revisions for a config

### Synopsis

Returns a paginated list of revisions for this config, sorted newest-first. Optionally filtered by status, author, or the calling user's involvement.

```
growthbook config-revisions list-for-config [flags]
```

### Examples

```
  growthbook config-revisions list-for-config --key <key>
```

### Options

```
  -a, --author string     string value
  -h, --help              help for list-for-config
  -k, --key string        [required]
  -l, --limit int         The number of items to return (default 10)
  -m, --mine author       If true, return only revisions authored by the calling user. Requires a user-scoped API key. Mutually exclusive with author.
      --offset int        How many items to skip (use in conjunction with limit for pagination)
      --skip-pagination   If true, return all matching items and ignore limit/offset.
                          Self-hosted only. Has no effect unless API_ALLOW_SKIP_PAGINATION is set to true or 1.
      --status open       Filter by revision status. Accepts a comma-separated list, or the literal open for non-merged/non-discarded revisions.
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

* [growthbook config-revisions](growthbook_config-revisions.md)	 - Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)
