## growthbook reports put-report-metadata

Update report metadata (title, description, visibility)

### Synopsis

Update report metadata (title, description, visibility)

```
growthbook reports put-report-metadata [flags]
```

### Examples

```
  growthbook reports put-report-metadata --id <id>
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string        Report description
  -e, --edit-level organization   Who can edit the report in the GrowthBook UI. organization allows any org member with the `createAnalyses` permission; `private` restricts editing to the report owner. (options: organization, private)
  -h, --help                      help for put-report-metadata
  -i, --id string                 The id of the requested resource [required]
      --share-level public        Visibility of the report. Setting to public enables a shareable `shareUrl`; setting back to `organization` or `private` revokes public access (the share token is preserved, so re-publishing exposes the same URL). (options: public, organization, private)
      --status string             UI lifecycle marker for the report (options: published, private)
  -t, --title string              Report title
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth                   If using Bearer auth, pass the Secret Key as the token:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                                      ```
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
      --password                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       username
```

### SEE ALSO

* [growthbook reports](growthbook_reports.md)	 - Custom analysis reports built on top of experiment snapshots
