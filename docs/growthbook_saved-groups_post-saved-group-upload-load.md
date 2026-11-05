## growthbook saved-groups post-saved-group-upload-load

Report loading a remote saved group upload into a data store

### Synopsis

Called by remote saved group loaders. Each report replaces the previous one from the same `storeId`. Needs permission to publish the saved group.

```
growthbook saved-groups post-saved-group-upload-load [flags]
```

### Examples

```
  growthbook saved-groups post-saved-group-upload-load --id <id> --version-param <value> --store-id <id> --status '{"type":"loaded","rowCount":650353}'
```

### Options

```
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                          help for post-saved-group-upload-load
  -i, --id string                     The id of the remote saved group [required]
      --status string                 JSON value (variants: loading: object, loaded: { rowCount: integer }, invalid: { invalidLineCount: integer, errors: string[] }, failed: { error: string })
      --status.failed string          status_Failed variant as JSON
      --status.failed.error string    A problem in the data store, not the file [required]
      --status.invalid string         status_Invalid variant as JSON
      --status.loaded string          status_Loaded variant as JSON
      --status.loaded.row-count int   Unique IDs loaded [required]
      --status.loading string         status_Loading variant as JSON
      --store-id redis-us-east        The data store's ID, such as redis-us-east, set in the loader's config. Loaders sharing a store (such as one Redis) use the same ID. Shown in the UI. [required]
  -v, --version-param string          [required]
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

* [growthbook saved-groups](growthbook_saved-groups.md)	 - Defined sets of attribute values which can be used with feature rules for targeting features at particular users
