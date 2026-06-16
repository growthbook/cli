## growthbook saved-group-revisions put-saved-group-revision-metadata

Update saved group metadata in a draft revision

### Synopsis

Stages metadata changes (name, owner, description, projects) on the draft. Pass `version: "new"` to auto-create a draft. The change is only applied to the live saved group when the revision is merged.

```
growthbook saved-group-revisions put-saved-group-revision-metadata [flags]
```

### Examples

```
  growthbook saved-group-revisions put-saved-group-revision-metadata --saved-group-id <id> --version-param <value>
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string        string value
  -h, --help                      help for put-saved-group-revision-metadata
  -n, --name string               string value
      --owner string              The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --projects stringArray      list of values
      --revision-comment string   string value
      --revision-title string     string value
  -s, --saved-group-id string     [required]
  -v, --version-param string      [required]
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

* [growthbook saved-group-revisions](growthbook_saved-group-revisions.md)	 - Draft revisions for saved groups, including pending changes, approvals, and lifecycle (publish, discard, revert)
