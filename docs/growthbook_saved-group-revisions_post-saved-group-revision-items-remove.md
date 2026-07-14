## growthbook saved-group-revisions post-saved-group-revision-items-remove

Remove items from a list saved group draft revision

### Synopsis

Removes the provided items from the draft's `values` array. Only valid for `list` saved groups. Pass `version: "new"` to auto-create a draft.

```
growthbook saved-group-revisions post-saved-group-revision-items-remove [flags]
```

### Examples

```
  growthbook saved-group-revisions post-saved-group-revision-items-remove --saved-group-id <id> --version-param <value> --items '["<value 1>"]'
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                      help for post-saved-group-revision-items-remove
  -i, --items stringArray         [required]
      --revision-comment string   string value
      --revision-title string     string value
  -s, --saved-group-id string     [required]
  -v, --version-param string      [required]
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
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook saved-group-revisions](growthbook_saved-group-revisions.md)	 - Draft revisions for saved groups, including pending changes, approvals, and lifecycle (publish, discard, revert)
