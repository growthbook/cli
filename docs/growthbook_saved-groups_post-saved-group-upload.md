## growthbook saved-groups post-saved-group-upload

Upload a new version of a remote saved group's IDs

### Synopsis

Records the next version of the group's IDs. The start and end of the file are checked for one ID per line, then the whole file in the background; loaders trim and deduplicate the IDs. Poll `GET /saved-groups/{id}/uploads/{version}` until it's valid, then set the group's `currentUploadVersion` with `PUT /saved-groups-revisions/{id}/{version}/upload-version` to use it. Loaders pick up the group's version from `currentUpload` and download it from `GET /saved-groups/{id}/uploads/{version}/file`, which redirects to a short-lived signed storage URL.

```
growthbook saved-groups post-saved-group-upload [flags]
```

### Examples

```
  growthbook saved-groups post-saved-group-upload --id <id> --file-key <value>
```

### Options

```
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --file-key fileKey   The fileKey from the upload URL endpoint [required]
      --file-name string   The file's original name, used when downloading it
  -h, --help               help for post-saved-group-upload
  -i, --id string          The id of the requested resource [required]
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
